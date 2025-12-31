package mytips

import (
	oddstoday "mytipster/internal/get-odds-today"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/today", GetPredictionByDay)
	api.Post("/predictions", InsertPredictions)
	// api.Put("/mytips/result", UpdateMatchResult)

	// -------------- * --------------
	// step 1
	api.Get("/get-odds-today", oddstoday.Service)
	// step 2
	api.Get("/mytips", WritePrediction)

	// upload bin/date/prediontion.json to db
	api.Get("/insert", InsertPredictions)
	// -------------- * -------------

	// api crud get

	// api.Get("/today", mytips.GetPredictionByDay)
}
