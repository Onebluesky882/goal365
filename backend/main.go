package main

import (
	"log"
	"mytipster/internal/db"
	"mytipster/internal/fixtures"
	"mytipster/internal/mytips"
	"mytipster/internal/predictions"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()
	db.InitDB()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173,https://mytipster-production.up.railway.app,https://goal365.thaimongkon777.workers.dev/,http://localhost:8787,http://localhost:3009",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))
	app.Get("/ids", fixtures.GetFixtureIds)
	app.Get("/id", fixtures.GetFixtureById)
	app.Get("/odds", fixtures.Odds)
	app.Get("/date", fixtures.GetFixtureDate)

	app.Get("/prediction-id", predictions.Service)

	// register routes
	mytips.RegisterRoutes(app)

	// todo
	/*
		  1. post new file  match result   bin/2025-12-30    today is 31 - 1  for guest
		  		- make json match_result (2025-12-30)
				- update app on db where id = id

	*/

	/*

		own file


	*/
	// post new file  bet pick handicap with new file /bin20-12-30.  for only me ที่จะบันทึกผลการเล่นส่วนตัว
	// frontend
	port := os.Getenv("PORT")
	if port == "" {
		port = "3009"
	}

	log.Println("🚀 Server running on port", port)
	log.Fatal(app.Listen(":" + port))
}
