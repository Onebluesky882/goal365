package bets

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"
)

func RegisterRoutes(app *fiber.App, db *bun.DB) {

	api := app.Group("/api")
	api.Post("/bets", InsertPickedHandler(db))
	api.Get("/bets", GetBetListsByDateHandler(db))
	api.Patch("/bets", UpdateMyBetsHandler(db))
}
