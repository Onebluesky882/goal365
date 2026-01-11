package player

import (
	"mytipster/models"

	"github.com/gofiber/fiber/v2"
)

func CreatePlayerHandler(service *PlayerService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req models.CreatePlayerRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid request body",
			})
		}
		ctx := c.Context()

		player, err := service.CreatePlayer(ctx, req.Name, req.UserID)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(201).JSON(player)

	}
}
