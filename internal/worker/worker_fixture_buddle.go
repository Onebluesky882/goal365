package worker

// import (
// 	"fmt"
// 	"log"
// 	m "mytipster/models/fixture"
// 	"strconv"
// 	"sync"
// 	"time"

// 	"github.com/gofiber/fiber/v2"
// )

// // WorkerFixtureService handles batch processing of fixture IDs
// func WorkerFixtureService(c *fiber.Ctx) error {
// 	ids := []int{
// 		1485798, 1486269, 1389669, 1485645, 1394567, 1379114,
// 		1387830, 1378002, 1389674, 1401121, 1486270, 1387326,
// 		1488408, 1483270, 1483852, 1451260, 1451432, 1382562,
// 		1379120, 1378006, 1380388, 1382726, 1394585, 1387351,
// 		1388439, 1347241, 1489310,
// 	}

// 	fixtureBundles, err := HelperSlickFixtureBuddle(ids)
// 	if err != nil {
// 		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
// 	}

// 	return c.JSON(fixtureBundles)
// }

// // HelperSlickFixtureBuddle processes fixtures concurrently
// func HelperSlickFixtureBuddle(ids []int) ([]*m.FixturePredictionBundle, error) {
// 	result := make([]*m.FixturePredictionBundle, 0, len(ids))

// 	workerCount := 4
// 	jobs := make(chan int)
// 	results := make(chan *m.FixturePredictionBundle)
// 	var wg sync.WaitGroup

// 	log.Printf("[main] start %d workers, %d jobs\n", workerCount, len(ids))

// 	// Start workers
// 	for i := 0; i < workerCount; i++ {
// 		wg.Add(1)
// 		go func(workerID int) {
// 			defer wg.Done()
// 			for fixtureId := range jobs {
// 				log.Printf("[worker-%d] processing %d\n", workerID, fixtureId)
// 				id := strconv.Itoa(fixtureId)

// 				// Use timeout wrapper to avoid long-hanging API calls
// 				temp, err := fixtureOddsPredictionWithTimeout(id, 5*time.Second)
// 				if err != nil {
// 					log.Printf("[worker-%d] skip %d: %v\n", workerID, fixtureId, err)
// 					continue
// 				}

// 				results <- &m.FixturePredictionBundle{
// 					Items: []m.FixturePrediction{*temp},
// 				}
// 			}
// 			log.Printf("[worker-%d] stopped\n", workerID)
// 		}(i)
// 	}

// 	// Sender
// 	go func() {
// 		for _, id := range ids {
// 			jobs <- id
// 		}
// 		close(jobs)
// 		log.Println("[sender] all jobs sent")
// 	}()

// 	// Closer
// 	go func() {
// 		wg.Wait()
// 		close(results)
// 		log.Println("[main] results closed")
// 	}()

// 	// Collector
// 	for r := range results {
// 		result = append(result, r)
// 	}

// 	log.Printf("[main] finished, success %d\n", len(result))
// 	return result, nil
// }

// // Timeout wrapper for fixtureOddsPrediction
// func fixtureOddsPredictionWithTimeout(id string, timeout time.Duration) (*m.FixturePrediction, error) {
// 	ch := make(chan *m.FixturePrediction)
// 	errCh := make(chan error)

// 	go func() {
// 		res, err := fixtureOddsPrediction(id)
// 		if err != nil {
// 			errCh <- err
// 			return
// 		}
// 		ch <- res
// 	}()

// 	select {
// 	case r := <-ch:
// 		return r, nil
// 	case e := <-errCh:
// 		return nil, e
// 	case <-time.After(timeout):
// 		return nil, fmt.Errorf("timeout for fixture %s", id)
// 	}
// }
