package betstransections

import (
	"mytipster/models"

	"github.com/gofiber/fiber/v2"
)

func InsertTransactionHandler(service *TransactionService) fiber.Handler {

	return func(c *fiber.Ctx) error {
		var req models.CreateTransactionRequest
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
