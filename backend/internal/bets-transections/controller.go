package betstransections

import (
	"mytipster/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func getTransactionHandler(service *TransactionService) fiber.Handler {

	return func(c *fiber.Ctx) error {
		var req models.UpdateTransactionRequest

		playerNo := c.Query("player_no")

		if playerNo == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "player_no is required",
			})
		}

		player, err := strconv.ParseInt(playerNo, 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid bill_id",
			})
		}
		req.PlayerNo = player
		tx, err := service.getTransaction(c.Context(), req)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if tx == nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "transaction not found",
			})
		}
		return c.Status(fiber.StatusOK).JSON(tx)
	}
}

func InsertTransactionHandler(service *TransactionService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req models.CreateTransactionRequest

		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid request body",
			})
		}

		// ---- map DTO -> DB model ----
		bets := make([]models.BetTransaction, len(req.Bets))
		for i, b := range req.Bets {
			bets[i] = models.BetTransaction{
				FixtureID: b.FixtureID,
				Market:    b.Market,
				Selection: b.Selection,
				Odds:      b.Odds,
				Amount:    b.Amount,
			}
		}

		tx, err := service.InsertTransaction(
			c.Context(),
			bets,
			req.PlayerNo,
		)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(tx)
	}
}
