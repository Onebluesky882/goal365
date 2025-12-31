package oddstoday

import (
	"fmt"
	"log"
	"mytipster/internal/fixtures/service"
	"mytipster/lib"
	odds_models "mytipster/models/odds"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Process single fixture odds with retry
func processSingleFixtureOdds(fixtureID int) (map[int][]odds_models.Bet, error) {
	var oddsMap map[int][]odds_models.Bet
	var err error
	idStr := strconv.Itoa(fixtureID)

	// Query odds พร้อม retry
	err = lib.RetryWithBackoff(func() error {
		time.Sleep(500 * time.Millisecond) // delay ก่อนเรียก API
		oddsMap, err = service.QueryFixtureOdds(idStr)
		if err != nil {
			return err
		}
		if len(oddsMap) == 0 {
			return fmt.Errorf("odds map is empty")
		}
		return nil
	}, 3, 500*time.Millisecond)

	if err != nil {
		return nil, fmt.Errorf("odds error for fixture %d: %w", fixtureID, err)
	}

	return oddsMap, nil
}

func QueryOdds(date string) (map[int][]odds_models.Bet, error) {
	// ดึง fixture IDs
	ids, err := service.GetIds(date)
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

	for i, fixtureID := range ids {
		wg.Add(1)
		go func(id int, idx int) {
			defer wg.Done()

			// Rate limiting
			sem <- struct{}{}
			defer func() { <-sem }()

			// เพิ่ม delay ก่อนเริ่มทำงาน
			time.Sleep(800 * time.Millisecond)

			log.Printf("⏳ [%d/%d] Processing fixture %d...\n", idx+1, len(ids), id)

			// Process with retry
			oddsMap, err := processSingleFixtureOdds(id)

			mu.Lock()
			defer mu.Unlock()

			if err != nil {
				errorCount++
				log.Printf("❌ [%d/%d] Failed fixture %d: %v\n", idx+1, len(ids), id, err)
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
		time.Sleep(800 * time.Millisecond)
	}

	// รอให้ทุก goroutine เสร็จ
	wg.Wait()

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

	result, err := QueryOdds(date)
	if err != nil {
		log.Printf("❌ Error in QueryOdds: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	log.Printf("🎉 ส่งผลลัพธ์ %d fixtures\n", len(result))
	outputPath := filepath.Join("bin", date, "odds_data.json")
	if err := lib.WriteJSONWithCustomDate(date, "odds_data.json", result); err != nil {
		log.Printf("⚠️  Warning: Failed to save JSON file: %v\n", err)
	} else {
		log.Printf("💾 บันทึกไฟล์: %s\n", outputPath)

		// ประมวลผลและกรองข้อมูล
		log.Println("\n🔄 เริ่มกรองข้อมูล...")
		if err := lib.ProcessOddsFile(outputPath); err != nil {
			log.Printf("⚠️  Warning: Failed to filter odds: %v\n", err)
		}
	}
	return c.JSON(result)
}
