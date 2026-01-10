package oddstoday

import (
	"fmt"
	"log"
	"mytipster/internal/fixtures"
	"mytipster/lib"
	m "mytipster/models"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

var oneMillisecond = 100 * time.Millisecond

func processSingleFixtureOdds(fixtureID int) ([]m.Bet, error) {
	var oddsMaps []m.Bookmaker
	var err error
	idStr := strconv.Itoa(fixtureID)

	// Query odds พร้อม retry
	err = lib.RetryWithBackoff(func() error {
		oddsMaps, err := fixtures.QueryFixtureOdds(idStr)
		fmt.Printf("DEBUG fixture %d => %v, err=%v\n", fixtureID, oddsMaps, err)
		if err != nil {
			log.Printf("⚠️ fixture %d query error: %v\n", fixtureID, err)
			return err
		}

		log.Printf("DEBUG fixture %d got oddsMaps count: %d\n", fixtureID, len(oddsMaps))
		if len(oddsMaps) == 0 {
			return fmt.Errorf("odds map is empty")
		}
		return nil
	}, 3, oneMillisecond)

	if err != nil {
		return nil, fmt.Errorf("odds error for fixture %d: %w", fixtureID, err)
	}

	// filter
	filtered := lib.FilterBookmaker(oddsMaps, []string{"Betano"})
	for _, bm := range filtered {
		fmt.Println("BOOKMAKER name =>", fmt.Sprintf("[%s]", bm.Name))
	}
	oddsFilter := lib.FilterBetType(filtered, []string{"Asian Handicap"})
	for _, bm := range oddsFilter {
		fmt.Println("BOOKMAKER oddsFilter =>", fmt.Sprintf("[%s]", bm.Name))
	}
	return oddsFilter, nil
}

func QueryOdds(date string) ([]m.Bet, error) {
	ids, err := fixtures.GetIds(date)
	if err != nil {
		return nil, err
	}

	const maxConcurrent = 4
	sem := make(chan struct{}, maxConcurrent)

	var (
		mu     sync.Mutex
		result []m.Bet
		wg     sync.WaitGroup
	)

	for i, fixtureID := range ids {
		wg.Add(1)
		defer wg.Done()
		go func(id int, idx int) {

			sem <- struct{}{}
			defer func() { <-sem }()

			log.Printf("⏳ [%d/%d] fixture %d\n", idx+1, len(ids), id)

			bets, err := processSingleFixtureOdds(id)
			if err != nil {
				log.Printf("❌ fixture %d error: %v\n", id, err)
				return
			}

			mu.Lock()
			result = append(result, bets...) // ✅ merge slice
			mu.Unlock()

		}(fixtureID, i)

		time.Sleep(oneMillisecond)
	}

	wg.Wait()
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
