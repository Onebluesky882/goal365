package main

import (
	"context"
	"log"
	"mytipster/internal/db"
	"mytipster/internal/db/analytics"
)

func main() {
	ctx := context.Background()

	database, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	if err := analytics.CreateTable(ctx, database); err != nil {
		log.Fatal(err)
	}

	log.Println("migration finished ✅")
}
