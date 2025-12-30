package mytips_db

import (
	"context"
	"fmt"
	"log"
	"mytipster/internal/db/service"
	m "mytipster/models/mytips"

	"github.com/uptrace/bun"
)

func CreateTable(ctx context.Context, db *bun.DB) error {
	_, err := db.NewCreateTable().
		Model((*m.MyTipsAnalytics)(nil)).IfNotExists().
		Exec(ctx)
	return err

}

func InsertMany(items []m.MyTipsAnalytics) error {
	ctx := context.Background()
	db := service.WithContext(ctx)
	var filtered []m.MyTipsAnalytics

	for _, item := range items {
		if item.FormLeagueHomeCount < 5 {
			continue
		}
		filtered = append(filtered, item)
	}

	if len(filtered) == 0 {
		fmt.Println("No records passed the filter, nothing to insert")
		return nil
	}

	_, err := db.NewInsert().Model(&filtered).Exec(ctx)
	if err != nil {
		log.Fatalf("insert many error %v", err)
	}

	fmt.Printf("✅ Inserted %d records\n", len(filtered))
	return nil
}

func GetPredictionByDay(date string) ([]m.MyTipsAnalytics, error) {
	ctx := context.Background()
	db := service.WithContext(ctx)
	var result []m.MyTipsAnalytics
	err := db.NewSelect().Model(&result).Where("date = ?", date).Scan(ctx)
	if err != nil {
		log.Fatalf("query error: %v", err)
	}
	return result, nil
}

// func UpdatePredictionAll(items []m.MyTipsAnalytics) error{}
