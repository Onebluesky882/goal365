package fixtures

import (
	"mytipster/internal/fixtures/service"

	"github.com/gofiber/fiber/v2"
)

func Predictions(c *fiber.Ctx) error {
	id := c.Query("fixture")
	resp, err := service.QueryPrediction(id)
	if err != nil {
		return err
	}
	return c.JSON(resp)
}
