package fixtures

import (
	"log"
	m "mytipster/models/fixture"
	"runtime"
	"strconv"
	"sync"

	"github.com/gofiber/fiber/v2"
)

func WorkerFixtureService(c *fiber.Ctx) error {
	// date := c.Query("date")

	// url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/fixtures?date=%s", date)
	// resp, err := api.Fixtures("GET", url)
	// if err != nil {
	// 	return err
	// }

	// step
	// got slice of fixture Ids
	// result, err := getFixtureIds(resp)
	// if err != nil {
	// 	return err
	// }

	// slice call function fixtureOddsPrediction

	ids := []int{
		1485798,
		1486269,
		1389669,
		1485645,
		1394567,
		1379114,
		1387830,
		1378002,
		1389674,
		1401121,
		1486270,
		1387326,
		1488408,
		1483270,
		1483852,
		1451260,
		1451432,
		1382562,
		1379120,
		1378006,
		1380388,
		1382726,
		1394585,
		1387351,
		1388439,
		1347241,
		1489310,
	}

	fixtureBuddle, err := HelperSlickFixtureBuddle(ids)
	if err != nil {
		return err
	}

	return c.JSON(fixtureBuddle)

	// next filter match tips Football with strategy

	// upload big query
}

/*  เหตุผลที่ใช้ []*Bundle

•	ไม่ copy struct ใหญ่ ๆ
•	ทำงานเร็ว
•	เป็น pattern มาตรฐานใน Go backend

*/

func HelperSlickFixtureBuddle(ids []int) ([]*m.FixturePredictionBundle, error) {
	result := make([]*m.FixturePredictionBundle, 0, len(ids))

	worker := runtime.NumCPU()
	jobs := make(chan int)
	results := make(chan *m.FixturePredictionBundle)

	var wg sync.WaitGroup

	log.Printf("[main] start %d workers, %d jobs\n", worker, len(ids))

	// 1️⃣ start workers
	for i := 0; i < worker; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			for fixtureId := range jobs {
				log.Printf("[worker-%d] processing %d\n", workerID, fixtureId)

				id := strconv.Itoa(fixtureId)
				temp, err := fixtureOddsPrediction(id)
				if err != nil {
					log.Printf("[worker-%d] skip %d: %v\n", workerID, fixtureId, err)
					continue
				}

				results <- &m.FixturePredictionBundle{
					FixtureIDs: []int{fixtureId},
					Items:      []m.FixturePrediction{*temp},
				}
			}
			log.Printf("[worker-%d] stopped\n", workerID)
		}(i)
	}

	// 2️⃣ sender (มีแค่ตัวเดียว)
	go func() {
		for _, id := range ids {
			log.Printf("[sender] send %d\n", id)
			jobs <- id
		}
		close(jobs)
		log.Println("[sender] all jobs sent")
	}()

	// 3️⃣ closer (มีแค่ตัวเดียว)
	go func() {
		wg.Wait()
		close(results)
		log.Println("[main] results closed")
	}()

	// 4️⃣ collector (อยู่ main goroutine)
	for r := range results {
		log.Printf("[collector] got %d\n", r.FixtureIDs[0])
		result = append(result, r)
	}

	log.Printf("[main] finished, success %d\n", len(result))
	return result, nil
}
