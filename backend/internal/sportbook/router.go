package sportbook

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/sports-book", sportBookHandler())
}
