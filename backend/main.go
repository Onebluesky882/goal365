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
		AllowOrigins: "http://localhost:5173,https://mytipster-production.up.railway.app,https://goal365.thaimongkon777.workers.dev/",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	app.Get("/ids", fixtures.GetFixtureIds)
	app.Get("/id", fixtures.GetFixtureById)
	app.Get("/odds", fixtures.Odds)
	app.Get("/date", fixtures.GetFixtureDate)

	app.Get("/tips", mytips.Service)
	app.Get("/prediction-id", predictions.Service)

	// -------------- * --------------
	// step 1
	app.Get("/get-odds-today", oddstoday.Service)
	// step 2
	app.Get("/mytips", mytips.Service)

	// upload bin/date/prediontion.json to db
	app.Get("/upload", mytips.Insert)
	// -------------- * -------------

	// api crud get

	app.Get("/today", mytips.GetPredictionByDay)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3009"
	}

	log.Println("🚀 Server running on port", port)
	log.Fatal(app.Listen(":" + port))
}
