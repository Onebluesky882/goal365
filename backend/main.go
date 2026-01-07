// package main

// import (
// 	"context"
// 	"log"
// 	"mytipster/internal/analytics"
// 	"mytipster/internal/auth"
// 	"mytipster/internal/bets"
// 	"mytipster/internal/database"
// 	"mytipster/internal/fixtures"
// 	"mytipster/internal/tipsdaily"
// 	"net/http"
// 	"os"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/gofiber/fiber/v2/middleware/cors"
// 	"github.com/joho/godotenv"
// )

// func main() {

// 	_ = godotenv.Load()

// 	database.InitDB()
// 	ctx := context.Background()
// 	db := database.WithContext(ctx)

// 	// Configure GoBetterAuth
// 	authModule := auth.New()

// 	// Initialize auth

// 	// Run migrations (creates necessary auth tables)
// 	app := fiber.New()

// 	app.Use(cors.New(cors.Config{
// 		AllowOrigins: "http://localhost:5173,https://mytipster-production.up.railway.app,https://goal365.thaimongkon777.workers.dev/,http://localhost:8787,http://localhost:3009",
// 		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
// 		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
// 	}))

// 	// Mount GoBetterAuth handler using Fiber's adaptor
// 	// This handles all /auth/* routes
// 	http.Handle("/auth/", authModule.Auth.Handler())

// 	// Public routes
// 	app.Get("/ids", fixtures.GetFixtureIds)
// 	app.Get("/fixture", fixtures.GetFixtureById)
// 	app.Get("/odds", fixtures.Odds)
// 	app.Get("/prediction", fixtures.Predictions)
// 	app.Get("/date", fixtures.GetFixtureDate)

// 	// register routes
// 	analytics.RegisterRoutes(app, db)
// 	tipsdaily.RegisterRoutes(app)
// 	bets.RegisterRoutes(app, db)
// 	// auth.RegisterRoute(app, db)

// 	log.Println("📡 LINE Auth endpoints:")
// 	log.Println("   - GET  /api/auth/signin/line")
// 	log.Println("   - GET  /api/auth/callback/line")
// 	log.Println("   - POST /api/auth/verify-line")
// 	log.Println("   - GET  /api/auth/session")
// 	log.Println("   - POST /api/auth/signout")
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "8080"
// 	}

// 	log.Println("🚀 Server running on port", port)
// 	log.Fatal(app.Listen(":" + port))
// }

package main

import (
	"log"
	"mytipster/internal/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	authModule := auth.New() // returns *Auth wrapper
	app := fiber.New()

	// CORS or middlewares
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true, // สำคัญถ้าใช้ cookie / session
	}))

	auth.RegisterAuthRoutes(app, authModule.Auth)
	// your own routes
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	log.Fatal(app.Listen(":3000"))
}
