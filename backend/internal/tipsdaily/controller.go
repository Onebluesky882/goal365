package tipsdaily

import (
	"context"
	"log"
	"mytipster/internal/db"

	"github.com/gofiber/fiber/v2"
)

func UpdateMatchResult(c *fiber.Ctx) error {
	// func MatchResult(date string) ([]m.UpdateFixtureResultDTO, error) {
	ctx := context.Background()
	db := db.WithContext(ctx)
	date := c.Query("date")
	results, err := MatchResult(date, db, ctx)

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

	if err := UpdateFixtureResult(results); err != nil {
		log.Println("❌ UpdateFixtureResult error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"success": true,
		"updated": len(results),
	})
}
