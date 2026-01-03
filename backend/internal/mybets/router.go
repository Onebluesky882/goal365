package mybets

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"
)

func RegisterRoutes(app *fiber.App, db *bun.DB) {

	api := app.Group("/api")
	api.Post("/mybets/insert", InsertPickedHandler(db))
}
