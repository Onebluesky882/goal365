package mytips

import (
	"fmt"
	"log"
	"mytipster/internal/fixtures/service"
	m "mytipster/models/fixture"
	"strconv"
	"sync"

	"github.com/gofiber/fiber/v2"
)

/*
1. get all fixture ids

2. loop odds with ids filter asia-handicap [ 0.25 , 0.5 , 0.75 ] everage price 1.80 - 2.20

fixture today 84
test normal ids -> odds
3.

4. insert db
*/
func MyTips(date string) (map[int]m.FixtureAnalytics, error) {

	result := make(map[int]m.FixtureAnalytics)
	var mu sync.Mutex
	var wg sync.WaitGroup

	ids, err := service.GetIds(date)
	if err != nil {
		return nil, err
	}
	sem := make(chan struct{}, 2)
	for _, fixtureId := range ids {
		wg.Add(1)
		sem <- struct{}{}

		go func(fid int) {
			defer wg.Done()
			defer func() { <-sem }()

			idStr := strconv.Itoa(fid)
			oddsMap, err := service.QueryMyTipsOdds(idStr)
			if err != nil {
				log.Printf("[MyTips] odds error %d: %v", fid, err)
				RetryLater(fid)
				return
			}

			for _, bets := range oddsMap {
				if len(bets) == 0 {
					continue
				}
				mu.Lock()
				result[fid] = m.FixtureAnalytics{
					FixtureID: fid,

					Handicap: bets[0],
				}
				mu.Unlock()
				break
			}
		}(fixtureId)
	}
	// รอ goroutine หลัก
	wg.Wait()
	WaitRetryDone()

	fmt.Println("result :", len(result))
	return result, nil
}

func Service(c *fiber.Ctx) error {
	date := c.Query("date")
	result, err := MyTips(date)
	if err != nil {
		return err
	}

	return c.JSON(result)
}

/*

fixture, err := service.QueryFixtureId(idStr)
			if err != nil {
				log.Printf("[MyTips] fixture error %d: %v", fid, err)
				return // ❌ ไม่ retry ซ้ำซ้อน
			}

			t, err := time.Parse(time.RFC3339, fixture.Fixture.Date)

			loc, _ := time.LoadLocation("Asia/Bangkok")
			bangkokTime := t.In(loc)

			dateStr := bangkokTime.Format("2006-01-02 15:04:05")

*/
