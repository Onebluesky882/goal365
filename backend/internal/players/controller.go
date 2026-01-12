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

func PlayerLoginLogsHandler(service *PlayerService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var req models.PlayerLoginLogRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "invalid request body",
			})
		}
		// 🔒 IP & UA จาก server เท่านั้น
		ip := c.IP()
		ua := c.Get("User-Agent")

		err := service.LogPlayerLogin(
			c.Context(),
			&req,
			ip,
			ua,
		)

		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.SendStatus(fiber.StatusCreated)

	}

}
 