package player

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"
)

func RegisterRoutes(app *fiber.App, db *bun.DB) {
	svc := NewPlayer(db)
	api := app.Group("/api")

	api.Post("/new-player", CreatePlayerHandler(svc))
}
