package sportbook

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func sportBookHandler(s *SportBook) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.UserContext()
		date := c.Query("date")
		if date == "" {
			log.Println("date is not correct")
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "date is required",
			})
		}

		if err := s.SyncMarketOdds(date, ctx); err != nil {
			log.Printf("❌ SyncMarketOdds failed: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"date":    date,
			"message": "market odds synced successfully",
		})
	}

}

func ManualInsertBookMakerHandler(s *SportBook) fiber.Handler {
	return func(c *fiber.Ctx) error {

		ctx := c.UserContext()
		date := c.Query("date")
		if date == "" {
			log.Println("date is not correct")
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "date is required",
			})
		}
		if err := s.InsertBookMaker(date, ctx); err != nil {
			log.Printf("❌ SyncMarketOdds failed: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,

			"message": "insert successfully",
		})
	}
}
