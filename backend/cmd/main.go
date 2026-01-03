package main

import (
	"context"
	"log"
	"mytipster/internal/db"
	"mytipster/internal/mybets"
	tipsdaily "mytipster/internal/tips-daily"
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

}
