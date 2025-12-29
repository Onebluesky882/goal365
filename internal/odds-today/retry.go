package oddstoday

import (
	"log"
	"mytipster/internal/fixtures/service"
	"strconv"
	"time"
)

func retryFixtureOdds(fixtureId int) {
	idStr := strconv.Itoa(fixtureId)

	for attempt := 1; attempt <= 3; attempt++ {
		time.Sleep(1 * time.Second)

		log.Printf("[Retry] fixture %d attempt %d/3", fixtureId, attempt)

		_, err := service.QueryMyTipsOdds(idStr)
		if err == nil {
			log.Printf("[Retry] fixture %d success", fixtureId)
			return
		}

		time.Sleep(2 * time.Second)
	}

	log.Printf("[Retry] fixture %d failed after 3 attempts", fixtureId)
}

func RetryLater(fixtureId int) {

	retryFixtureOdds(fixtureId)

}
