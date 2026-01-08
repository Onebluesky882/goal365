package fixtures

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/fixture-ids", GetFixtureIds)
	api.Get("/fixture-date", GetFixtureDate)
	api.Get("/odds-id", Odds)
	api.Get("/fixture-id", GetFixtureById)
}
