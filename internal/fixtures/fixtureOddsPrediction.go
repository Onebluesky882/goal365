package fixtures

import (
	"log"
	"mytipster/internal/predictions"
	m "mytipster/models/fixture"

	"github.com/gofiber/fiber/v2"
)

func fixtureOddsPrediction(id string) (*m.FixturePrediction, error) {

	// 1. Fetch fixture data
	fixtureData, err := QueryFixtureById(id)
	if err != nil {
		// log missing fixture and skip
		log.Printf("[fixtureOddsPrediction] no fixture for id %s: %v", id, err)
		return nil, err
	}

	// 2. Fetch full fixture response
	fixtureResponse, err := fixtureDataService(id)
	if err != nil {
		log.Printf("[fixtureOddsPrediction] failed to get fixture response for id %s: %v", id, err)
		return nil, err
	}

	// 3. Fetch predictions
	predictionsData, err := predictions.Service(id)
	if err != nil {
		log.Printf("[fixtureOddsPrediction] failed to get predictions for id %s: %v", id, err)
		return nil, err
	}
	// 4. Build result safely
	result := &m.FixturePrediction{
		FixtureID: fixtureData.ID,
		Fixture:   fixtureResponse,
	}

	if len(predictionsData.Response) > 0 {
		result.Predictions = &predictionsData.Response[0]
	} else {
		log.Printf("[fixtureOddsPrediction] no predictions for fixture %d", fixtureData.ID)
		result.Predictions = nil
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
