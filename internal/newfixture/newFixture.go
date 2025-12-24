package newFixture

import (
	"fmt"
	"mytipster/api"

	"github.com/gofiber/fiber/v2"
)

func Service(c *fiber.Ctx) error {
	date := c.Query("date")
	if date == "" {
		c.Status(500).JSON(fiber.Map{
			"messess": "wrong date",
		})
	}

	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/fixtures?date=%s", date)

	resp, err := api.Fixtures("GET", url)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(resp)
}
