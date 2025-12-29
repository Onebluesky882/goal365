package main

import (
	"log"
	"mytipster/internal/fixtures"
	"mytipster/internal/mytips"
	oddstoday "mytipster/internal/odds-today"
	"mytipster/internal/predictions"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()

	app := fiber.New()
	app.Get("/ids", fixtures.GetFixtureIds)
	app.Get("/id", fixtures.GetFixtureById)
	app.Get("/odds", fixtures.Odds)
	app.Get("/date", fixtures.GetFixtureDate)

	app.Get("/tips", mytips.Service)
	app.Get("/prediction", predictions.Service)

	// -------------- * --------------

	in := "/Users/onebluesky882/local_files/myjob/mytipster/bin/data.json"
	out := "/Users/onebluesky882/local_files/myjob/mytipster/bin/output.json"

	if err := oddstoday.ProcessOddsFile(in, out); err != nil {
		log.Fatal(err)
	}

	// -------------- * --------------

	port := os.Getenv("PORT")
	if port == "" {
		port = "3009"
	}

	log.Println("🚀 Server running on port", port)
	log.Fatal(app.Listen(":" + port))
}
