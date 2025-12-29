package analytics

import (
	"context"
	fixture_module "mytipster/models/fixture"

	"github.com/uptrace/bun"
)

func CreateTable(ctx context.Context, db *bun.DB) error {
	_, err := db.NewCreateTable().
		Model((*fixture_module.FixtureAnalytics)(nil)).IfNotExists().
		Exec(ctx)
	return err

}

func InsertData(ctx context.Context, db *bun.DB, item *fixture_module.FixtureAnalytics) error {
	_, err := db.NewInsert().Model(item).Exec(ctx)
	return err
}
