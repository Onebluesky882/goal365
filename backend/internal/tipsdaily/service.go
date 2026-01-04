package tipsdaily

import (
	"context"
	tipsdaliy_models "mytipster/models/tipsdaliy"

	"github.com/uptrace/bun"
)

func CreateTable(ctx context.Context, db *bun.DB) error {
	_, err := db.NewCreateTable().
		Model((*tipsdaliy_models.TipsDaily)(nil)).IfNotExists().
		Exec(ctx)
	return err

}
