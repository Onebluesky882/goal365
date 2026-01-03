package main

import (
	"context"
	"log"
	"mytipster/internal/analytics"
	"mytipster/internal/db"
	"mytipster/internal/fixtures"
	"mytipster/internal/mybets"
	"mytipster/internal/tipsdaily"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()
	db.InitDB()
	ctx := context.Background()
	dbbun := db.WithContext(ctx)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173,https://mytipster-production.up.railway.app,https://goal365.thaimongkon777.workers.dev/,http://localhost:8787,http://localhost:3009",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))
	app.Get("/ids", fixtures.GetFixtureIds)
	app.Get("/fixture", fixtures.GetFixtureById)
	app.Get("/odds", fixtures.Odds)
	app.Get("/prediction", fixtures.Predictions)
	app.Get("/date", fixtures.GetFixtureDate)

	// register routes
	analytics.RegisterRoutes(app)
	tipsdaily.RegisterRoutes(app)
	mybets.RegisterRoutes(app, dbbun)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3009"
	}

	log.Println("🚀 Server running on port", port)
	log.Fatal(app.Listen(":" + port))
}
