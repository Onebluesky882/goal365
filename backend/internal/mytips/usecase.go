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
	"os"
	"path/filepath"
	"strconv"
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
			Picked: "",
			Team:   "",
			Odds:   "",
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
	var failedFixturesMu sync.Mutex
	failedFixtures := make([]int, 0)
	for i, fixtureID := range ids {
		time.Sleep(1000 * time.Millisecond)
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
				fixture, err := strconv.Atoi(id)
				errorCount++
				log.Printf("❌ [%d/%d] Failed %s: %v\n", idx+1, len(ids), id, err)
				failedFixturesMu.Lock()
				failedFixtures = append(failedFixtures, fixture)
				failedFixturesMu.Unlock()
				return
			}

			successCount++
			results = append(results, *item)
			log.Printf("✅ [%d/%d] Success %s (%s vs %s)\n",
				idx+1, len(ids), id, item.Home, item.Away)
		}(fixtureID, i, bets)

		// Rate limiting between goroutine starts
		time.Sleep(1000 * time.Millisecond)
	}

	// รอให้ทุก goroutine เสร็จ
	wg.Wait()
	// บันทึก fixtures ที่ล้มเหลวลงไฟล์ (manual)
	if len(failedFixtures) > 0 {
		if err := WriteFailedPredictions(failedFixtures); err != nil {
			log.Println("❌ Failed to write failed fixtures:", err)
		} else {
			log.Printf("📝 Wrote %d failed fixtures to file for manual retry\n", len(failedFixtures))
		}
	}
	log.Printf("\n📊 สรุปผลลัพธ์:\n")
	log.Printf("   ✅ สำเร็จ: %d\n", successCount)
	log.Printf("   ❌ ล้มเหลว: %d\n", errorCount)
	log.Printf("   📦 รวมทั้งหมด: %d\n\n", len(results))

	return &m.RootMyTipsAnalytics{
		Count: len(results),
		Items: results,
	}, nil
}

func WriteFailedPredictions(failed []int) error {
	date := time.Now().Format("2025-12-12")
	outputDir := filepath.Join("bin", date)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("cannot create directory: %w", err)
	}
	outputFile := filepath.Join(outputDir, "error_query_prediction.json")
	if err := lib.WriteJSON(outputFile, failed); err != nil {
		return fmt.Errorf("cannot write failed fixtures file: %w", err)
	}
	return nil
}

func MatchResult(date string) ([]m.UpdateFixtureResultDTO, error) {

	predictions, err := PredictionByDay(date)
	if err != nil {
		return nil, err
	}

	fixtures, err := service.QueryFixtureDate(date)
	if err != nil {
		return nil, err
	}
	// เก็บ fixtureId
	fixtureMap := make(map[int]fixture_module.Response)

	// merge
	for _, fx := range fixtures {
		fixtureMap[fx.Fixture.ID] = fx
	}

	results := make([]m.UpdateFixtureResultDTO, 0, len(predictions))

	for _, p := range predictions {
		fx, ok := fixtureMap[p.FixtureID]
		if !ok {
			continue
		}

		home := 0
		away := 0
		if fx.Goals.Home != nil {
			home = *fx.Goals.Home
		}
		if fx.Goals.Away != nil {
			away = *fx.Goals.Away
		}
		results = append(results, m.UpdateFixtureResultDTO{
			FixtureID:   fx.Fixture.ID,
			MatchFinish: fx.Fixture.Status.Long,
			MatchResult: fmt.Sprintf("%d-%d", home, away),
		})
	}
	return results, nil
}

func updateBet(fixtureId string) (*m.BetPick, error) {
	// todo
	return nil, nil
}
