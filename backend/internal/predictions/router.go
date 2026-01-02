package predictions

import (
	oddstoday "mytipster/internal/odds-today"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/today", GetPredictionByDay)
	api.Post("/predictions", InsertPredictions)

	// --------------  get daliy prodiction* --------------
	// step 1
	api.Get("/get-odds-today", oddstoday.GetOddsToday)
	// step 2
	api.Get("/write-predictions", writePredictions)

	// 2.1 retry
	api.Get("/retry-predictions", insertRetryPrediction)

	// upload bin/date/prediontion.json to db
	api.Get("/insert", InsertPredictions)
	// -------------- * -------------

	// api crud get

	// api.Get("/today", mytips.GetPredictionByDay)
}
