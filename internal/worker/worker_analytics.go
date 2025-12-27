package worker

import (
	"log"
	m "mytipster/models/fixture"
	"strconv"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

// AnalyticsFixture handles batch processing of fixture IDs
func AnalyticsFixtures(c *fiber.Ctx) error {
	ids := []int{
		1485798, 1486269, 1389669, 1485645, 1394567, 1379114,
		1387830, 1378002, 1389674, 1401121, 1486270, 1387326,
		1488408, 1483270, 1483852, 1451260, 1451432, 1382562,
		1379120, 1378006, 1380388, 1382726, 1394585, 1387351,
		1388439, 1347241, 1489310, 1386821,
	}

	go func() {
		log.Printf("[BG] Starting batch process for %d ids", len(ids))

		result, err := HelpersGetFixtureByIds(ids)
		if err != nil {
			log.Printf("[BG ERROR] Process failed: %v", err)
			return
		}

		// เมื่อดึงข้อมูลเสร็จค่อยส่งขึ้น BigQuery ที่นี่
		log.Printf("[BG] Process complete, items: %d. Starting Upload...", len(result.Items))
		// err = UploadToBigQuery(result)
		// if err != nil { log.Printf("[BQ ERROR] %v", err) }
	}()

	return c.JSON(fiber.Map{
		"status":  "accepted",
		"message": "Processing in background, check BigQuery soon",
		"count":   len(ids),
	})
}

func HelpersGetFixtureByIds(ids []int) (*m.RootFixtureBundle, error) {
	bundle := &m.RootFixtureBundle{
		Items: []m.FixtureBuddle{},
	}

	jobs := make(chan int, len(ids))
	results := make(chan m.FixtureBuddle, len(ids))

	// --- [ปรับจุดที่ 1] ลด Worker เหลือ 2-3 ตัว ---
	// การมี Worker เยอะเกินไปในขณะที่ API ช้า จะยิ่งทำให้ติด Rate Limit
	workerCount := 2

	var wg sync.WaitGroup

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func(workerId int) {
			defer wg.Done()
			for id := range jobs {
				idStr := strconv.Itoa(id)

				// เรียกใช้ Retry ที่เราจะปรับให้รอนานขึ้น (ข้อ 2)
				fixture, err := QueryFixtureIdRetry(idStr)
				if err != nil {
					log.Printf("[SKIP] fixture %s FAIL after all retries", idStr)
					continue
				}

				pred, _ := QueryPredictionRetry(idStr)

				// 3. ผลลัพธ์ (เช็ค nil เพื่อความปลอดภัย)
				if fixture != nil {
					results <- m.FixtureBuddle{
						FixtureID:   fixture.Fixture.ID,
						Fixture:     fixture,
						Predictions: pred,
					}
				}

				// --- [ปรับจุดที่ 2] พักระหว่างงานให้นานขึ้นเป็น 2-3 วินาที ---
				time.Sleep(2000 * time.Millisecond)
			}
		}(i)
	}

	go func() {
		for _, id := range ids {
			jobs <- id
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for item := range results {
		bundle.Items = append(bundle.Items, item)
	}

	return bundle, nil
}

func retry[T any](
	max int,
	delay time.Duration,
	fn func() (*T, error),
) (*T, error) {

	var lastErr error

	for i := 1; i <= max; i++ {
		result, err := fn()
		if err == nil {
			return result, nil
		}

		lastErr = err
		log.Printf("[RETRY %d/%d] %v", i, max, err)
		time.Sleep(delay)
	}

	return nil, lastErr
}
