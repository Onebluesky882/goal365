package mybets

import (
	"context"
	m "mytipster/models/mytips"

	"github.com/uptrace/bun"
)

func CreateTable(ctx context.Context, db *bun.DB) error {
	_, err := db.NewCreateTable().
		Model((*m.MyBets)(nil)).IfNotExists().
		Exec(ctx)
	return err

}
