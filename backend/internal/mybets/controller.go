package mybets

import (
	mybets_models "mytipster/models/mybets"

	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"
)

func InsertPickedHandler(db *bun.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req mybets_models.InsertPickedRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid request body",
			})
		}

		ctx := c.Context()

		// เรียก service ตาม signature ใหม่
		if err := InsertPicked(req.Items, req.AnalyticsID, db, ctx); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"success":  true,
			"inserted": len(req.Items),
		})
	}
}
