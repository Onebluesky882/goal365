package main

import (
	"context"
	"log"
	"mytipster/internal/analytics"
	"mytipster/internal/bets"
	"mytipster/internal/database"
	"mytipster/internal/fixtures"
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
		AllowOrigins: "https://mytipster-production.up.railway.app,https://goal365.thaimongkon777.workers.dev/,http://localhost:3001",
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
	analytics.RegisterRoutes(app, db)
	tipsdaily.RegisterRoutes(app)
	bets.RegisterRoutes(app, db)
	// auth.RegisterRoute(app, db)

	log.Println("📡 LINE Auth endpoints:")
	log.Println("   - GET  /api/auth/signin/line")
	log.Println("   - GET  /api/auth/callback/line")
	log.Println("   - POST /api/auth/verify-line")
	log.Println("   - GET  /api/auth/session")
	log.Println("   - POST /api/auth/signout")
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Println("🚀 Server running on port", port)
	log.Fatal(app.Listen(":" + port))
}
