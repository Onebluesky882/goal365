package betstransections

import (
	"mytipster/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CreateTransactionRequest struct {
	PlayerId uuid.UUID               `json:"player_id"`
	Bets     []models.BetTransaction `json:"bets"`
}

func InsertTransactionHandler(service *TransactionService) fiber.Handler {

	return func(c *fiber.Ctx) error {
		var req CreateTransactionRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid request body",
			})
		}

		tx, err := service.InsertTransaction(c.Context(), req)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(tx)
	}
}
