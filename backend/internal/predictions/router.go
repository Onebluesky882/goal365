package predictions

import (
	oddstoday "mytipster/internal/odds-today"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/today", GetPredictionByDay)
	api.Post("/predictions", InsertPredictions)

	// manual insert prediction by id
	api.Post("/insert-manual-id", insertManualPrediction)

	// -------------- * --------------
	// step 1
	api.Get("/get-odds-today", oddstoday.GetOddsToday)
	// step 2
	api.Get("/write-predictions", writePredictions)

	// upload bin/date/prediontion.json to db
	api.Get("/insert", InsertPredictions)
	// -------------- * -------------

	// api crud get

	// api.Get("/today", mytips.GetPredictionByDay)
}
