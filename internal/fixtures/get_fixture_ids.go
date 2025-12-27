package fixtures

import (
	"mytipster/internal/fixtures/service"

	"github.com/gofiber/fiber/v2"
)

func GetFixtureIds(c *fiber.Ctx) error {
	date := c.Query("date")

	resp, err := service.GetIds(date)
	if err != nil {
		return err
	}
	return c.JSON(resp)
}
