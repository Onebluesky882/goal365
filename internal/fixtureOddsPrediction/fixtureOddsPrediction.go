package fixtureoddsprediction

import (
	fixturebyid "mytipster/internal/fixtureById"
	"mytipster/internal/odds"
	"mytipster/internal/predictions"
	odds_models "mytipster/models/odds"
	prediction_models "mytipster/models/prediction"

	"github.com/gofiber/fiber/v2"
)

type FixtureOddsPredictionType struct {
	FixtureID   int
	Predictions []prediction_models.PredictionResponse
	Bookmaker   map[int][]odds_models.Bet
	Result      string
	Picked      bool
}

func service(id string) (*FixtureOddsPredictionType, error) {

	fixtureData, err := fixturebyid.Service(id)
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

	result := &FixtureOddsPredictionType{
		FixtureID:   fixtureData.ID,
		Predictions: predictionsData.Response,
		Bookmaker:   oddsData,
	}

	return result, nil

}

func FixtureOddsPredictionHandler(c *fiber.Ctx) error {
	id := c.Query("id")
	result, err := service(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(result)
}
