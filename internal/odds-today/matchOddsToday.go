package oddstoday

import (
	"log"
	"mytipster/internal/fixtures/service"
	odds_models "mytipster/models/odds"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func QueryOdds(date string) (map[int][]odds_models.Bet, error) {
	result := make(map[int][]odds_models.Bet)
	retried := make(map[int]bool)
	ids, err := service.GetIds(date)
	if err != nil {
		return nil, err
	}

	for _, fixtureId := range ids {
		idStr := strconv.Itoa(fixtureId)

		oddsMap, err := service.QueryFixtureOdds(idStr)
		if err != nil {
			log.Printf("[OddMatchToday] skip fixture %d: %v", fixtureId, err)
			if !retried[fixtureId] {
				retried[fixtureId] = true
				retryFixtureOdds(fixtureId)
			}
			continue
		}

		for k, bets := range oddsMap {
			result[k] = append(result[k], bets...)
		}

		time.Sleep(2 * time.Second) // rate limit
	}

	log.Println("[OddMatchToday] finished")
	log.Println("total :", len(result))

	return result, nil
}

func Service(c *fiber.Ctx) error {
	date := c.Query("date")
	result, err := QueryOdds(date)
	if err != nil {
		return err
	}

	return c.JSON(result)
}
