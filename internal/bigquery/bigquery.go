package bigquery

import (
	"context"
	"fmt"
	"mytipster/models"
	"os"

	"cloud.google.com/go/bigquery"
	"github.com/joho/godotenv"
)

func Service(rows []models.FixtureBigQuery) {



	_ = godotenv.Load()
	ctx := context.Background()

	client, err := bigquery.NewClient(ctx, os.Getenv("PROJECTID"))
	if err != nil {
		panic(err)
	}

	inserter := client.Dataset(os.Getenv("DATASET")).Table(os.Getenv("TABLE")).Inserter()

items := make([]*models.FixtureBigQuery, len(rows))
	for i := range rows {
		items[i] = &rows[i]
	}
	if err := inserter.Put(ctx, items); err != nil {
		if putErr, ok := err.(bigquery.PutMultiError); ok {
			for _, e := range putErr {
				fmt.Println("Row error:", e.Errors)
			}
		} else {
			fmt.Println("Insert error:", err)
		}
		return
	}

	fmt.Println("✅ Inserted", len(items), "row(s) into BigQuery")
}