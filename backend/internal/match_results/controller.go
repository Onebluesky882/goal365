package matchresults

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func MatchResultHandler(service MatchResultService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		ctx := c.Context()
		date := c.Query("date")
		results, err := service.MatchResult(ctx, date)

		if err != nil {
			log.Println("❌ MatchResult error:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		if len(results) == 0 {
			return c.JSON(fiber.Map{
				"success": true,
				"updated": 0,
				"message": "no fixtures to update",
			})
		}

		return c.JSON(fiber.Map{
			"success": true,
			"updated": len(results),
		})
	}
}
