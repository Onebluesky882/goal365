package main

import (
	"log"
	"mytipster/internal/fixtures"
	oddstoday "mytipster/internal/get-odds-today"
	"mytipster/internal/mytips"
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

	// todo
	// -------------- * --------------
	app.Get("/get-odds-today", oddstoday.Service)
	app.Get("/mytips", mytips.Service)
	// -------------- * --------------

	port := os.Getenv("PORT")
	if port == "" {
		port = "3009"
	}

	log.Println("🚀 Server running on port", port)
	log.Fatal(app.Listen(":" + port))
}
