package oddstoday

import (
	"fmt"
	"log"
	"mytipster/internal/fixtures"
	"mytipster/lib"
	odds_models "mytipster/models/odds"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

var oneMillisecond = 100 * time.Millisecond

/*
  ✅ สำเร็จ: 207
    ❌ ล้มเหลว: 223
*/

/* 2026/01/02 17:58:15    ✅ สำเร็จ: 207
2026/01/02 17:58:15    ❌ ล้มเหลว: 223 */

// Process single fixture odds with retry
func processSingleFixtureOdds(fixtureID int) (map[int][]odds_models.Bet, error) {
	var oddsMap map[int][]odds_models.Bet
	var err error
	idStr := strconv.Itoa(fixtureID)

	// Query odds พร้อม retry
	err = lib.RetryWithBackoff(func() error {
		oddsMap, err = fixtures.QueryFixtureOdds(idStr)
		if err != nil {
			return err
		}
		if len(oddsMap) == 0 {
			return fmt.Errorf("odds map is empty")
		}
		return nil
	}, 3, oneMillisecond)

	if err != nil {
		return nil, fmt.Errorf("odds error for fixture %d: %w", fixtureID, err)
	}

	return oddsMap, nil
}

func QueryOdds(date string) (map[int][]odds_models.Bet, error) {
	// ดึง fixture IDs
	ids, err := fixtures.GetIds(date)
	if err != nil {
		return nil, fmt.Errorf("failed to get fixture IDs: %w", err)
	}

	log.Printf("🚀 เริ่มดึง odds สำหรับ %d fixtures (date: %s)\n", len(ids), date)

	// ใช้ concurrent processing แบบเดียวกับ Predictions
	const maxConcurrent = 1 // ทำทีละตัวเพื่อหลีกเลี่ยง rate limit
	sem := make(chan struct{}, maxConcurrent)

	var mu sync.Mutex
	result := make(map[int][]odds_models.Bet)
	var wg sync.WaitGroup

	successCount := 0
	errorCount := 0

	var failedFixturesMu sync.Mutex
	failedFixtures := make([]int, 0)
	for i, fixtureID := range ids {
		wg.Add(1)
		go func(id int, idx int) {
			defer wg.Done()

			// Rate limiting
			sem <- struct{}{}
			defer func() { <-sem }()

			log.Printf("⏳ [%d/%d] Processing id : fixture %d \n", idx+1, len(ids), id)

			// Process with retry
			oddsMap, err := processSingleFixtureOdds(id)

			mu.Lock()
			defer mu.Unlock()

			if err != nil {
				errorCount++
				log.Printf("❌ [%d/%d] Failed fixture %d: %v\n",
					idx+1, len(ids), id, err,
				)
				failedFixturesMu.Lock()
				failedFixtures = append(failedFixtures, id)
				failedFixturesMu.Unlock()

				return
			}

			// Merge odds data
			for k, bets := range oddsMap {
				result[k] = append(result[k], bets...)
			}

			successCount++
			log.Printf("✅ [%d/%d] Success fixture %d (%d bet types)\n",
				idx+1, len(ids), id, len(oddsMap))
		}(fixtureID, i)

		// Rate limiting between goroutine starts
		time.Sleep(oneMillisecond)
	}

	// รอให้ทุก goroutine เสร็จ
	wg.Wait()

	if len(failedFixtures) > 0 {
		errFile := "error_query_odds.json"
		if err := lib.WriteJSONWithCustomDate(date, errFile, failedFixtures); err != nil {
			log.Println("❌ write error_query_odds.json failed:", err)
		} else {
			log.Printf("🧾 wrote %d failed fixtures to %s\n",
				len(failedFixtures), errFile)
		}
	}
	log.Printf("\n📊 สรุปผลลัพธ์:\n")
	log.Printf("   ✅ สำเร็จ: %d\n", successCount)
	log.Printf("   ❌ ล้มเหลว: %d\n", errorCount)
	log.Printf("   📦 รวมทั้งหมด: %d fixtures with odds\n\n", len(result))

	return result, nil
}

func GetOddsToday(c *fiber.Ctx) error {
	date := c.Query("date")
	if date == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "date parameter is required",
		})
	}

	log.Printf("📅 Query odds for date: %s\n", date)

	// ดึง odds หลัก
	result, err := QueryOdds(date)
	if err != nil {
		log.Printf("❌ Error in QueryOdds: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	outputDir := filepath.Join("bin", date)
	outputFile := filepath.Join(outputDir, "odds_data.json")
	errFile := filepath.Join(outputDir, "error_query_odds.json")

	// สร้างโฟลเดอร์ถ้ายังไม่มี
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		log.Printf("❌ cannot create directory: %v\n", err)
	}

	// เขียน odds data
	if err := lib.WriteJSON(outputFile, result); err != nil {
		log.Printf("⚠️  Failed to save odds_data.json: %v\n", err)
	} else {
		log.Printf("💾 Saved odds data: %s\n", outputFile)
	}

	// --- บันทึก failed fixtures ---
	var failedFixtures []int
	for id, bets := range result {
		if len(bets) == 0 {
			failedFixtures = append(failedFixtures, id)
		}
	}

	if len(failedFixtures) > 0 {
		if err := lib.WriteJSON(errFile, failedFixtures); err != nil {
			log.Printf("❌ Failed to write error_query_odds.json: %v\n", err)
		} else {
			log.Printf("📝 Wrote failed fixtures file: %s (%d failed)\n", errFile, len(failedFixtures))
		}
	} else {
		log.Printf("✅ No failed fixtures for date %s\n", date)
	}

	// กรอง odds
	if err := lib.ProcessOddsFile(outputFile); err != nil {
		log.Printf("❌ Error processing filtered odds: %v\n", err)
	} else {
		log.Printf("✅ Filtered odds saved successfully\n")
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"updated": len(result),
	})
}
