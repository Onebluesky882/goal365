package player

import (
	"mytipster/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"
)

func RegisterRoutes(app *fiber.App, db *bun.DB) {
	svc := NewPlayer(db)

	api := app.Group("/api")

	// 🌍 Public API (frontend เรียกได้)
	api.Get("/players", getPlayer(svc))
	api.Post("/new-player", createPlayerHandler(svc))

	// 🔒 Internal API (backend → backend)
	api.Post(
		"/player/login-log",
		middleware.InternalOnly,
		playerLoginLogsHandler(svc),
	)
}
