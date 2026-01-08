package tipsdaily

import (
	"context"
	"mytipster/models"

	"github.com/uptrace/bun"
)

func CreateTable(ctx context.Context, db *bun.DB) error {
	_, err := db.NewCreateTable().
		Model((*models.TipsDaily)(nil)).IfNotExists().
		Exec(ctx)
	return err

}
