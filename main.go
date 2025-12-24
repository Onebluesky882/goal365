package main

import (
	"log"
	m "mytipster/internal/fixtureOddsPrediction"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()

	app := fiber.New()

	app.Get("/odds", m.FixtureOddsPredictionHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3009"
	}

	log.Println("🚀 Server running on port", port)
	log.Fatal(app.Listen(":" + port))
}
