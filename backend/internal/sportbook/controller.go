package sportbook

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func sportBookHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		date := c.Query("date")
		if date == "" {
			log.Println("date is not correct")
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "date is required",
			})
		}
		resp, err := GetMarketOdds(date)
		if err != nil {
			log.Printf("❌ GetMarketOdds failed: %v\n", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(resp)
	}

}
