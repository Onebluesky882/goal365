package main

import (
	"context"
	"log"
	"mytipster/internal/db"
	"mytipster/internal/mytips"
)

func main() {
	db.InitDB()
	ctx := context.Background()
	db := db.WithContext(ctx)
	if err := mytips.CreateTable(ctx, db); err != nil {

		log.Fatal(err)
	}

}
