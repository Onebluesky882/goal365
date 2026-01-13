package player

import (
	"mytipster/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"
)

func RegisterRoutes(app *fiber.App, db *bun.DB) {
	svc := NewPlayer(db)
	api := app.Group("/api", middleware.InternalOnly)

	api.Post("/new-player", createPlayerHandler(svc))
	api.Post(
		"/player/login-log",
		playerLoginLogsHandler(svc),
	)
	api.Get("/players", getPlayer(svc))
}
