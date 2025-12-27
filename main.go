package main

import (
	"log"
	"mytipster/internal/fixtures"
	"mytipster/internal/worker"
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
	app.Get("/prediction", fixtures.Predictions)
	app.Get("/date", fixtures.GetFixtureDate)

	// -------------- * --------------
	app.Get("/analytics", worker.AnalyticsFixtures)
	// app.Get("/predictions", worker.AnalyticsFixture)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3009"
	}

	log.Println("🚀 Server running on port", port)
	log.Fatal(app.Listen(":" + port))
}
