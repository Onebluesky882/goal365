package mytips_db

import (
	"context"
	"fmt"
	"log"
	"mytipster/internal/db/service"
	m "mytipster/models/mytips"
	mytips_module "mytipster/models/mytips"

	"github.com/uptrace/bun"
)

func CreateTable(ctx context.Context, db *bun.DB) error {
	_, err := db.NewCreateTable().
		Model((*m.MyTipsAnalytics)(nil)).IfNotExists().
		Exec(ctx)
	return err

}

func InsertMany(items []mytips_module.MyTipsAnalytics) error {

	ctx := context.Background()
	db := service.WithContext(ctx)

	var filtered []mytips_module.MyTipsAnalytics

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
