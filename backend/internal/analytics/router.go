package analytics

import (
	oddstoday "mytipster/internal/odds-today"

	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"
)

func RegisterRoutes(app *fiber.App, db *bun.DB) {
	api := app.Group("/api") // inject db into service
	svc := NewAnalyticService(db)

	api.Post("/nawinta", naWinTaTipsHandler(svc))

	api.Post("/predictions", InsertPredictions(svc))
	// --------------  get daliy prodiction* --------------
	// step 1
	api.Get("/get-odds-today", oddstoday.GetOddsToday)
	// step 2
	api.Get("/write-predictions.json", WritePredictionsHandler(svc))

	// 2.1 retry
	api.Get("/retry-predictions", InsertRetryPrediction(svc))

	// 3.0 upload http://localhost:3009/api/insert?date=2026-01-05
	// bin/date/prediontion.json to db
	api.Get("/insert", InsertPredictions(svc))
	// -------------- * -------------

	// for Frontedn get api

	api.Get("/analytics", GetPredictionByDayHandler(svc))

	// match result
	api.Patch("/match-result", MatchResultHandler(svc))

}
