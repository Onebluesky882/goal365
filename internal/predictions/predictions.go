package predictions

import (
	"fmt"
	"mytipster/api"
	prediction_models "mytipster/models/prediction"

	"github.com/gofiber/fiber/v2"
)


func Service(c *fiber.Ctx) (error) {
fixture := c.Query("fixture")
	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/predictions?fixture=%s" , fixture)
	resp , err := api.ApiFootball[prediction_models.PredictionsRoot]("GET" , url)
		if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return  c.JSON(resp)
}
