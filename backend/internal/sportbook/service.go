package sportbook

import (
	"context"
	"fmt"
	"log"
	"mytipster/internal/database"
	"mytipster/internal/fixtures"
	"mytipster/lib"
	"mytipster/lib/common"
	"mytipster/models"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/uptrace/bun"
)

type SportBook struct {
	db *bun.DB
}





func NewSportBook(db *bun.DB) *SportBook {
	return &SportBook{
		db: db,
	}
}

func GetMarketOdds(date string) ([]models.SportsBook, error) {

	ids, err := fixtures.GetIdsWithFilterCountry(date, []string{})
	if len(ids) == 0 {
		log.Println("⚠️ No fixtures after country filter")
		return []models.SportsBook{}, nil
	}
	if err != nil {
		return nil, err
	}

	var (
		markets      []models.SportsBook
		wg           sync.WaitGroup
		mu           sync.Mutex
		successCount int64
		failCount    int64
		skippedCount int64
	)

	const maxConcurrent = 4
	sem := make(chan struct{}, maxConcurrent)

	for _, id := range ids {
		wg.Add(1)

		go func(fxId int) {
			defer wg.Done()

			// ---- acquire semaphore ----
			sem <- struct{}{}
			defer func() { <-sem }()

			time.Sleep(200 * time.Millisecond)

			// ---- fixture ----
			fxResp, err := fixtures.QueryFixtureId(strconv.Itoa(fxId))
			if err != nil {
				log.Printf("❌ Fixture %d API error: %v", fxId, err)
				atomic.AddInt64(&failCount, 1)
				return
			}

			if fxResp == nil {
				log.Printf("⚠️ Fixture %d empty response", fxId)
				atomic.AddInt64(&failCount, 1)
				return
			}

			// ---- throttle (odds) ----
			time.Sleep(200 * time.Millisecond)
			// ---- odds ----
			oddsResp, err := fixtures.QueryFixtureOdds(strconv.Itoa(fxResp.FixtureInfo.ID))
			if err != nil || oddsResp == nil {
				log.Printf("❌ Odds %d failed: %v", fxId, err)
				atomic.AddInt64(&failCount, 1)
				return
			}

			filtered := lib.FilterBookmaker(oddsResp, []string{"Betano"})
			if len(filtered) == 0 {
				log.Printf("⏭️ Fixture %d skipped (no Betano odds)", fxId)
				atomic.AddInt64(&skippedCount, 1)
				return
			}

			now := time.Now().UTC()

			market := models.SportsBook{
				FxId:      fxResp.FixtureInfo.ID,
				MatchDate: common.StringToTime(fxResp.FixtureInfo.Date),
				Country:   fxResp.League.Country,
				League:    fxResp.League.Name,
				Home:      fxResp.Teams.Home.Name,
				Away:      fxResp.Teams.Away.Name,

				Bet:       []models.Bet{},
				CreatedAt: now,
				UpdatedAt: now,
			}

			for _, bm := range filtered {
				market.Bet = append(market.Bet, bm.Bets...)
			}

			mu.Lock()
			markets = append(markets, market)
			mu.Unlock()

			atomic.AddInt64(&successCount, 1)

		}(id)
	}
	wg.Wait()

	log.Printf(
		"📊 Market Odds Summary | Total: %d | Success: %d | Failed: %d | Skipped: %d",
		len(ids),
		successCount,
		failCount,
		skippedCount,
	)
	fileName := "sportbook/bookmaker.json"
	if err := lib.WriteJSONWithCustomDate(date, fileName, markets); err != nil {
		log.Printf("❌ Write JSON failed: %v", err)
		return markets, err
	}

	log.Printf("✅ JSON written successfully (%d records)", len(markets))
	// insert db

	return markets, nil
}

func (s *SportBook) InsertBookMaker(date string, ctx context.Context) error {
	path := fmt.Sprintf("bin/%s/sportbook/bookmaker.json", date)
	if len(path) == 0 {
		log.Println("⚠️ No records to insert")
		return nil
	}
	lists, err := lib.ReadJson[[]models.SportsBook](path)
	if err != nil {
		return fmt.Errorf("read json failed: %w", err)
	}
	_, err = s.db.NewInsert().Model(&lists).Exec(ctx)

	if err != nil {
		return fmt.Errorf("insert bookmaker failed: %w", err)
	}
	log.Printf("✅ Inserted %d records\n", len(lists))
	return nil

}

// all in one process
func (s *SportBook) SyncMarketOdds(date string, ctx context.Context) error {
	markets, err := GetMarketOdds(date)
	if err != nil {
		return err
	}

	if len(markets) == 0 {
		log.Println("⚠️ No markets to insert")
		return nil
	}
	// crate table partition sportsbook_2026_01_02
	matchDate := markets[0].MatchDate
	if err := database.EnsurePartition(ctx, s.db, matchDate); err != nil {
		return err
	}
	// insert table
	if err := s.InsertBookMaker(date, ctx); err != nil {
		return err
	}
	return nil
}
