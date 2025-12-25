package fixtures

import (
	"mytipster/internal/odds"
	"mytipster/internal/predictions"
	m "mytipster/models/fixture"

	"github.com/gofiber/fiber/v2"
)

func fixtureOddsPrediction(id string) (*m.FixturePrediction, error) {

	fixtureData, err := QueryFixtureById(id)
	if err != nil {
		return nil, err
	}
	oddsData, err := odds.Service(id)
	if err != nil {
		return nil, err
	}
	// 4. เรียก service ดึง predictions
	predictionsData, err := predictions.Service(id)
	if err != nil {
		return nil, err
	}

	result := &m.FixturePrediction{
		FixtureID:   fixtureData.ID,
		Predictions: predictionsData.Response,
		Bookmaker:   oddsData,
	}

	return result, nil

}

// main -> FixtureOddsPredictionHandler -> fixtureOddsPrediction

func FixtureOddsPredictionHandler(c *fiber.Ctx) error {
	id := c.Query("id")
	result, err := fixtureOddsPrediction(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(result)
}
