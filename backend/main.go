package main

import (
	"log"
	"mytipster/internal/db/service"
	"mytipster/internal/fixtures"
	oddstoday "mytipster/internal/get-odds-today"
	"mytipster/internal/mytips"
	"mytipster/internal/predictions"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()
	service.InitDB()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173,https://mytipster-production.up.railway.api,https://goal365.thaimongkon777.workers.dev/,http://localhost:8787",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))
	api := app.Group("/api")
	api.Get("/ids", fixtures.GetFixtureIds)
	api.Get("/id", fixtures.GetFixtureById)
	api.Get("/odds", fixtures.Odds)
	api.Get("/date", fixtures.GetFixtureDate)

	api.Get("/tips", mytips.Service)
	api.Get("/prediction-id", predictions.Service)

	// -------------- * --------------
	// step 1
	api.Get("/get-odds-today", oddstoday.Service)
	// step 2
	api.Get("/mytips", mytips.Service)

	// upload bin/date/prediontion.json to db
	api.Get("/upload", mytips.Insert)
	// -------------- * -------------

	// api crud get

	api.Get("/today", mytips.GetPredictionByDay)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3009"
	}

	log.Println("🚀 Server running on port", port)
	log.Fatal(app.Listen(":" + port))
}
