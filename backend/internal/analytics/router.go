package analytics

import (
	oddstoday "mytipster/internal/odds-today"

	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"
)

func RegisterRoutes(app *fiber.App, db *bun.DB) {
	api := app.Group("/api") // inject db into service
	svc := NewAnalyticService(db)

	api.Post("/predictions", InsertPredictions(svc))
	// --------------  get daliy prodiction* --------------
	// step 1
	api.Get("/get-odds-today", oddstoday.GetOddsToday)
	// step 2
	api.Get("/write-predictions", writePredictions)

	// 2.1 retry
	api.Get("/retry-predictions", InsertRetryPrediction(db, svc))

	// upload bin/date/prediontion.json to db
	api.Get("/insert", InsertPredictions(svc))
	// -------------- * -------------

	// api crud get

	// api.Get("/today", mytips.GetPredictionByDay)
}
