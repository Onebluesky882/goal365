package db

import (
	"context"
	m "mytipster/models/fixture"
)

func Migrate(ctx context.Context) error {
	_, err := DB.NewCreateTable().
		Model((*m.FixtureAnalytics)(nil)).
		IfNotExists().
		Exec(ctx)

	return err
}
