package main

import (
	"context"
	"log"
	"mytipster/internal/db/analytics"
	mytips_db "mytipster/internal/db/mytips"
	"mytipster/internal/db/service"
)

func main() {
	service.InitDB()
	ctx := context.Background()
	db := service.WithContext(ctx)
	if err := mytips_db.CreateTable(ctx, db); err != nil {

		log.Fatal(err)
	}

	if err := analytics.CreateTable(ctx, db); err != nil {
		log.Fatal(err)
	}

}
