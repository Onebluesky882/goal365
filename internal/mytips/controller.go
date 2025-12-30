package mytips

import (
	mytips_db "mytipster/internal/db/mytips"

	"github.com/gofiber/fiber/v2"
)

func GetPredictionByDay(c *fiber.Ctx) error {

	
	date := c.Query("date")

	predictions, err := mytips_db.GetPredictionByDay(date)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(predictions)
}
