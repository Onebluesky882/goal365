package fixtures

import (
	"github.com/gofiber/fiber/v2"
		"mytipster/internal/fixtures/service"
)

func GetFixtureDate(c *fiber.Ctx) error {

	date := c.Query("date")
	res, err := service.QueryFixtureDate(date)
	if err != nil {
		return err
	}

	return c.JSON(res)
}
