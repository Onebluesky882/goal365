package main

import (
	"log"
	"mytipster/internal/db/service"
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
	service.InitDB()
	app := fiber.New()
	app.Get("/ids", fixtures.GetFixtureIds)
	app.Get("/id", fixtures.GetFixtureById)
	app.Get("/odds", fixtures.Odds)
	app.Get("/date", fixtures.GetFixtureDate)

	app.Get("/tips", mytips.Service)
	app.Get("/prediction", predictions.Service)

	// -------------- * --------------
	// step 1
	app.Get("/get-odds-today", oddstoday.Service)
	// step 2
	app.Get("/mytips", mytips.Service)

	// upload bin/date/prediontion.json to db
	app.Get("/upload", mytips.Insert)
	// -------------- * -------------

	// todo

	/*
		insert prediction to db

		post pick match

		post check result update goal match finish

	*/

	port := os.Getenv("PORT")
	if port == "" {
		port = "3009"
	}

	log.Println("🚀 Server running on port", port)
	log.Fatal(app.Listen(":" + port))
}
