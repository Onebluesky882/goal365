package bets

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"
)

func RegisterRoutes(app *fiber.App, db *bun.DB) {

	api := app.Group("/api")
	api.Post("/mybets/insert", InsertPickedHandler(db))
	api.Get("/mybets", GetBetListsByDateHandler(db))
	api.Patch("/mybets/update", UpdateMyBetsHandler(db))
}
