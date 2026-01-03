package predictions

import (
	"context"
	"fmt"
	"log"
	"mytipster/internal/db"
	m "mytipster/models/mytips"

	"github.com/uptrace/bun"
)

func CreateTable(ctx context.Context, db *bun.DB) error {
	_, err := db.NewCreateTable().
		Model((*m.MyTipsAnalytics)(nil)).IfNotExists().
		Exec(ctx)
	return err

}

func InsertManual(item *m.MyTipsAnalytics) error {
	ctx := context.Background()
	db := db.WithContext(ctx)

	_, err := db.NewInsert().Model(item).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func insertMany(items []m.MyTipsAnalytics) error {
	ctx := context.Background()
	db := db.WithContext(ctx)
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

func PredictionByDay(date string) ([]m.MyTipsAnalytics, error) {
	ctx := context.Background()
	db := db.WithContext(ctx)
	var result []m.MyTipsAnalytics
	err := db.NewSelect().Model(&result).Where("date = ?", date).Scan(ctx)
	if err != nil {
		log.Fatalf("query error: %v", err)
	}
	return result, nil
}
