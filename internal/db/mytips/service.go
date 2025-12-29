package mytips

import (
	"context"
	m "mytipster/models/mytips"

	"github.com/uptrace/bun"
)

func CreateTable(ctx context.Context, db *bun.DB) error {
	_, err := db.NewCreateTable().
		Model((*m.MyTipsAnalytics)(nil)).IfNotExists().
		Exec(ctx)
	return err

}

func InsertData(ctx context.Context, db *bun.DB, item *m.MyTipsAnalytics) error {
	_, err := db.NewInsert().Model(item).Exec(ctx)
	return err
}
