package betstransections

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"
)

func RegisterRoutes(app *fiber.App, db *bun.DB) {
	svc := NewTransaction(db)
	api := app.Group("/api")

	api.Get("/transactions", getTransactionHandler(svc))
	api.Post("/bet-slips", InsertTransactionHandler(svc))

}
