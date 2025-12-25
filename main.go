package main

import (
	"log"
	"mytipster/internal/fixtures"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	

	_ = godotenv.Load()

	app := fiber.New()

	app.Get("/analytics" , fixtures.WorkerFixtureService)
	app.Get("/predictions", fixtures.WorkerFixtureService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3009"
	}

	log.Println("🚀 Server running on port", port)
	log.Fatal(app.Listen(":" + port))
}
