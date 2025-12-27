package fixtures

import (
	"mytipster/internal/fixtures/service"

	"github.com/gofiber/fiber/v2"
)

func Odds(c *fiber.Ctx) error {
	id := c.Query("fixture")
	res, err := service.QueryFixtureOdds(id)
	if err != nil {
		return err
	}

	return c.JSON(res)

}
