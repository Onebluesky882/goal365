package mytips

import (
	"log"
	"mytipster/internal/fixtures/service"
	"strconv"
	"sync"
	"time"
)

var retryWG sync.WaitGroup
var sem = make(chan struct{}, 1) // semaphore ขนาด 1 = ทำงานทีละตัว

func retryFixtureOdds(fixtureId int) {
	idStr := strconv.Itoa(fixtureId)

	for attempt := 1; attempt <= 3; attempt++ {
		log.Printf("[Retry] fixture %d attempt %d/3", fixtureId, attempt)

		_, err := service.QueryMyTipsOdds(idStr)
		if err == nil {
			log.Printf("[Retry] fixture %d success", fixtureId)
			return
		}

		time.Sleep(4 * time.Second)
	}

	log.Printf("[Retry] fixture %d failed after 3 attempts", fixtureId)
}

func RetryLater(fixtureId int) {
	retryWG.Add(1)

	go func(id int) {
		defer retryWG.Done()

		// ขอสิทธิ์ก่อนเข้า critical section
		sem <- struct{}{}
		defer func() { <-sem }() // ปล่อยสิทธิ์เมื่อเสร็จ

		retryFixtureOdds(id)
	}(fixtureId)
}

func WaitRetryDone() {
	retryWG.Wait()
}
