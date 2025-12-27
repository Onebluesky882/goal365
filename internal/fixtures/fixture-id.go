package fixtures

import (
	"mytipster/internal/fixtures/service"

	"github.com/gofiber/fiber/v2"
)

func GetFixtureById(c *fiber.Ctx) error {

	id := c.Query("id")
	res, err := service.QueryFixtureId(id)
	if err != nil {
		return err
	}

	return c.JSON(res)
}
