package sportbook

import (
	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"
)

func RegisterRoutes(app *fiber.App, db *bun.DB) {
	svc := NewSportBook(db)
	api := app.Group("/api")
	api.Get("/sports-book", sportBookHandler(svc))
	api.Get("/manual-insert-sports-book", ManualInsertBookMakerHandler(svc))
}
