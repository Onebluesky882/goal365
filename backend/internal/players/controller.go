package player

import (
	"mytipster/models"

	"github.com/gofiber/fiber/v2"
)

func createPlayerHandler(service *PlayerService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req models.CreatePlayerRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid request body",
			})
		}
		ctx := c.Context()

		player, err := service.createPlayer(ctx, req.Name, req.Bio, req.UserID)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(201).JSON(player)

	}
}

func playerLoginLogsHandler(service *PlayerService) fiber.Handler {
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

func getPlayersHandler(s *PlayerService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := c.Query("user_id")
		if userId == "" {
			return c.Status(400).JSON(fiber.Map{
				"error": "userId is required",
			})
		}

		res, err := s.getPlayers(c.Context(), userId)
		if err != nil {

			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(201).JSON(res)
	}
}
func getPlayersByNoHandler(s *PlayerService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		playerNo := c.Query("player_no")

		if playerNo == "" {
			return c.Status(400).JSON(fiber.Map{
				"error": "playerNo is required",
			})
		}

		res, err := s.getPlayerByNo(c.Context(), playerNo)
		if err != nil {

			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(201).JSON(res)
	}
}
