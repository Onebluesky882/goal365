package main

import (
	"context"
	"log"
	"mytipster/internal/analytics"
	"mytipster/internal/database"
	"mytipster/internal/fixtures"
	"mytipster/internal/sportbook"
	"mytipster/internal/tipsdaily"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()

	database.InitDB()
	ctx := context.Background()
	db := database.WithContext(ctx)

	// Run migrations (creates necessary auth tables)
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://mytipster-production.up.railway.app,https://goal365.thaimongkon777.workers.dev/,http://localhost:3001,https://goal365.club",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// Public routes
	app.Get("/ids", fixtures.GetFixtureIds)
	app.Get("/fixture", fixtures.GetFixtureById)
	app.Get("/odds", fixtures.Odds)
	app.Get("/prediction", fixtures.Predictions)
	app.Get("/date", fixtures.GetFixtureDate)

	// register routes
	fixtures.RegisterRoutes(app)
	analytics.RegisterRoutes(app, db)
	tipsdaily.RegisterRoutes(app)
	// mybets.RegisterRoutes(app, db)

	// market auth
	sportbook.RegisterRoutes(app, db)

	port := os.Getenv("PORT")

	if port == "" {
		port = "3009"
	}

	log.Println("🚀 Server running on port", port)
	log.Fatal(app.Listen(":" + port))
}
