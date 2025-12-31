package mytips

import (
	"fmt"
	"log"
	"mytipster/internal/fixtures/service"
	"mytipster/lib"
	"mytipster/lib/common"
	fixture_module "mytipster/models/fixture"
	m "mytipster/models/mytips"
	odds_models "mytipster/models/odds"
	prediction_models "mytipster/models/prediction"
	"sync"
	"time"
)

// Process single fixture with retry
func ProcessBuildPredictionsJson(fixtureID string, bet []odds_models.Bet) (*m.MyTipsAnalytics, error) {
	var pred *prediction_models.PredictionResponse
	var fx *fixture_module.Response
	var err error

	// ดึง prediction พร้อม retry
	err = lib.RetryWithBackoff(func() error {
		pred, err = service.QueryPrediction(fixtureID)
		if err != nil {
			return err
		}
		if pred == nil {
			return fmt.Errorf("prediction is nil")
		}
		return nil
	}, 3, 500*time.Millisecond)

	if err != nil {
		return nil, fmt.Errorf("prediction error for %s: %w", fixtureID, err)
	}
	// ดึง fixture พร้อม retry
	err = lib.RetryWithBackoff(func() error {
		time.Sleep(500 * time.Millisecond)
		fx, err = service.QueryFixtureId(fixtureID)
		if err != nil {
			return err
		}
		if fx == nil {
			return fmt.Errorf("fixture is nil")
		}
		return nil
	}, 3, 500*time.Millisecond)

	if err != nil {
		return nil, fmt.Errorf("fixture error for %s: %w", fixtureID, err)
	}

	// สร้าง result
	home, away := 0, 0
	if fx.Goals.Home != nil {
		home = *fx.Goals.Home
	}
	if fx.Goals.Away != nil {
		away = *fx.Goals.Away
	}
	item := &m.MyTipsAnalytics{
		FixtureID:           fx.Fixture.ID,
		Date:                common.TimestampUTCDate(fx.Fixture.Timestamp),
		TimeStamp:           common.Timestamp(fx.Fixture.Timestamp),
		Country:             fx.League.Country,
		League:              fx.League.Name,
		Home:                pred.Teams.Home.Name,
		Away:                pred.Teams.Away.Name,
		HomeScore:           pred.Teams.Home.Last5.Form,
		AwayScore:           pred.Teams.Away.Last5.Form,
		FormLeagueHomeCount: len(pred.Teams.Home.League.Form),
		FormLeagueAwayCount: len(pred.Teams.Away.League.Form),
		HomeFormScore14:     lib.FormScore(14, pred.Teams.Home.League.Form),
		AwayFormScore14:     lib.FormScore(14, pred.Teams.Away.League.Form),
		HomeFormScore12:     lib.FormScore(12, pred.Teams.Home.League.Form),
		AwayFormScore12:     lib.FormScore(12, pred.Teams.Away.League.Form),
		HomeFormScore10:     lib.FormScore(10, pred.Teams.Home.League.Form),
		AwayFormScore10:     lib.FormScore(10, pred.Teams.Away.League.Form),
		HomeFormScore7:      lib.FormScore(7, pred.Teams.Home.League.Form),
		AwayFormScore7:      lib.FormScore(7, pred.Teams.Away.League.Form),
		HomeFormScore5:      lib.FormScore(5, pred.Teams.Home.League.Form),
		AwayFormScore5:      lib.FormScore(5, pred.Teams.Away.League.Form),
		MatchFinish:         fx.Fixture.Status.Long,
		MatchResult:         fmt.Sprintf("%d-%d", home, away),
		Handicap:            bet[0],
		BetPick: m.BetPick{
			Odds:   "",
			Picked: "",
			Stake:  "",
		},
	}

	return item, nil
}

func Predictions(ids []string, oddsMap map[string][]odds_models.Bet) (*m.RootMyTipsAnalytics, error) {

	// ใช้ concurrent processing
	const maxConcurrent = 1 // จำกัด concurrent requests
	sem := make(chan struct{}, maxConcurrent)

	var mu sync.Mutex
	results := make([]m.MyTipsAnalytics, 0, len(ids))
	var wg sync.WaitGroup

	successCount := 0
	errorCount := 0

	log.Printf("🚀 เริ่มประมวลผล %d fixtures...\n", len(ids))

	for i, fixtureID := range ids {
		time.Sleep(800 * time.Millisecond)
		// ตรวจสอบว่ามี odds หรือไม่
		bets, ok := oddsMap[fixtureID]
		if !ok || len(bets) == 0 {
			log.Printf("⏭️  [%d/%d] Skip %s: no odds data\n", i+1, len(ids), fixtureID)
			continue
		}

		wg.Add(1)
		go func(id string, idx int, betData []odds_models.Bet) {
			defer wg.Done()

			// Rate limiting
			sem <- struct{}{}
			defer func() { <-sem }()

			log.Printf("⏳ [%d/%d] Processing fixture %s...\n", idx+1, len(ids), id)

			// Process with retry
			item, err := ProcessBuildPredictionsJson(id, betData)

			mu.Lock()
			defer mu.Unlock()

			if err != nil {
				errorCount++
				log.Printf("❌ [%d/%d] Failed %s: %v\n", idx+1, len(ids), id, err)
				return
			}

			successCount++
			results = append(results, *item)
			log.Printf("✅ [%d/%d] Success %s (%s vs %s)\n",
				idx+1, len(ids), id, item.Home, item.Away)
		}(fixtureID, i, bets)

		// Rate limiting between goroutine starts
		time.Sleep(800 * time.Millisecond)
	}

	// รอให้ทุก goroutine เสร็จ
	wg.Wait()

	log.Printf("\n📊 สรุปผลลัพธ์:\n")
	log.Printf("   ✅ สำเร็จ: %d\n", successCount)
	log.Printf("   ❌ ล้มเหลว: %d\n", errorCount)
	log.Printf("   📦 รวมทั้งหมด: %d\n\n", len(results))

	return &m.RootMyTipsAnalytics{
		Count: len(results),
		Items: results,
	}, nil
}

func MatchResult() {

}
