package main

import (
	"context"
	"log"
	"mytipster/internal/db"
	"mytipster/internal/db/analytics"
	mytips_db "mytipster/internal/db/mytips"
)

func main() {
	ctx := context.Background()

	database, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	if err := mytips_db.CreateTable(ctx, database); err != nil {
		log.Fatal(err)
	}

	if err := analytics.CreateTable(ctx, database); err != nil {
		log.Fatal(err)
	}

}
