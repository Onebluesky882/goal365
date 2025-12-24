package main

import (
	"log"
	newFixture "mytipster/internal/newfixture"
	"mytipster/internal/odds"
	"mytipster/internal/predictions"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()

	app := fiber.New()

	app.Get("/fixtures", newFixture.Service)
	app.Get("/predictions", predictions.Service)
	app.Get("/odds", odds.Service)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3009"
	}

	log.Println("🚀 Server running on port", port)
	log.Fatal(app.Listen(":" + port))
}
