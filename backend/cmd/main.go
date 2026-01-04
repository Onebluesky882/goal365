package main

import (
	"context"
	"log"
	"mytipster/internal/database"
)

func main() {
	database.InitDB()
	ctx := context.Background()
	db := database.WithContext(ctx)

	if err := database.CreateTables(ctx, db); err != nil {
		log.Fatal("fail create tables", err)
	}

	// Generate SQL DDL จาก struct
	/*

	   func main() {
	   	db := bun.NewDB(nil, pgdialect.New())

	   	query := db.NewCreateTable().
	   		Model((*MyBets)(nil)).
	   		IfNotExists()

	   	sql, _ := query.AppendQuery(nil, nil)q
	   	fmt.Println(string(sql))
	   }

	*/
}
