package fixtures

import (
	"fmt"
	"log"
	"mytipster/api"
	m "mytipster/models/fixture"
	"runtime"
	"strconv"
	"sync"

	"github.com/gofiber/fiber/v2"
)

func WorkerFixtureService(c *fiber.Ctx) error {
	date := c.Query("date")

	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/fixtures?date=%s", date)
	resp, err := api.Fixtures("GET", url)
	if err != nil {
		return err
	}

	// step
	// got slice of fixture Ids
	result, err := getFixtureIds(resp)
	if err != nil {
		return err
	}

	// slice call function fixtureOddsPrediction

	fixtureBuddle, err := HelperSlickFixtureBuddle(result)
	if err != nil {
		return err
	}

	return c.JSON(fixtureBuddle)

	// next filter match tips Football with strategy

	// upload big query
}

func HelperSlickFixtureBuddle(ids []int) ([]*m.FixturePredictionBundle, error) {
	result := make([]*m.FixturePredictionBundle, 0, len(ids))
	/*  เหตุผลที่ใช้ []*Bundle

	•	ไม่ copy struct ใหญ่ ๆ
	•	ทำงานเร็ว
	•	เป็น pattern มาตรฐานใน Go backend

	*/

	worker := runtime.NumCPU()
	jobs := make(chan int)
	results := make(chan *m.FixturePredictionBundle)

	var wg sync.WaitGroup

	log.Printf("[main] start with %d workers, %d jobs\n", worker, len(ids))
	for i := 0; i < worker; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			for fixtureId := range jobs {
				log.Printf("[worker-%d] processing fixture %d\n", workerID, fixtureId)
				id := strconv.Itoa(fixtureId)

				temp, err := fixtureOddsPrediction(id)
				if err != nil {
					log.Printf("[worker-%d] skip fixture %d: %v\n", workerID, fixtureId, err)
					continue
				}
				results <- &m.FixturePredictionBundle{
					FixtureIDs: []int{fixtureId},
					Items:      []m.FixturePrediction{*temp},
				}
			}
			log.Printf("[worker-%d] stopped\n", workerID)
		}(i)
		// 2️⃣ sender
		go func() {
			for _, id := range ids {
				log.Printf("[sender] send job %d\n", id)
				jobs <- id
			}
			close(jobs)
			log.Println("[sender] all jobs sent")
		}()
		// 3️⃣ closer
		go func() {
			wg.Wait()
			close(results)
			log.Println("[main] all workers done, results closed")
		}()
		// 4️⃣ collector
		for r := range results {
			log.Printf("[collector] got result fixture %d\n", r.FixtureIDs[0])
			result = append(result, r)
		}
	}
	log.Printf("[main] finished, total success %d\n", len(result))
	return result, nil
}
