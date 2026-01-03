package main

import (
	"context"
	"log"
	"mytipster/internal/analytics"
	"mytipster/internal/db"
	"mytipster/internal/mybets"
	"mytipster/internal/tipsdaily"
)

func main() {
	db.InitDB()
	ctx := context.Background()
	db := db.WithContext(ctx)
	if err := tipsdaily.CreateTable(ctx, db); err != nil {
		log.Fatal(err)
	}

	if err := mybets.CreateTable(ctx, db); err != nil {
		log.Fatal(err)
	}

	if err := analytics.CreateTable(ctx, db); err != nil {
		log.Fatal(err)
	}

}
