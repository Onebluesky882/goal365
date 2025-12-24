package odds

import (
	"fmt"
	"mytipster/api"
	odds_models "mytipster/models/odds"

	"github.com/gofiber/fiber/v2"
)

func Service(c *fiber.Ctx) error {

	fixture := c.Query("fixture")
	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/odds?fixture=%s", fixture)

	resp, err := api.ApiFootball[odds_models.RootOdds]("GET", url)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(resp)
}
