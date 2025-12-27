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
	result, err := HelpersGetFixtureByIds(ids)
	if err != nil {
		return err
	}

	return c.JSON(result)
}

func HelpersGetFixtureByIds(ids []int) (*m.RootFixtureBundle, error) {
	bundle := &m.RootFixtureBundle{
		Items: []m.FixtureBuddle{},
	}

	jobs := make(chan int, len(ids))
	results := make(chan m.FixtureBuddle, len(ids))

	// --- [ปรับจุดที่ 1] ลด Worker เหลือ 2-3 ตัว ---
	// การมี Worker เยอะเกินไปในขณะที่ API ช้า จะยิ่งทำให้ติด Rate Limit
	workerCount := 4
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

				odds, _ := QueryFixtureOddsSafe(idStr)
				pred, _ := QueryPredictionSafe(idStr)

				results <- m.FixtureBuddle{
					FixtureID:   fixture.Fixture.ID,
					Fixture:     fixture,
					Predictions: pred,
					Bookmaker:   odds,
				}

				// --- [ปรับจุดที่ 2] พักระหว่างงานให้นานขึ้นเป็น 2-3 วินาที ---
				time.Sleep(2500 * time.Millisecond)
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
