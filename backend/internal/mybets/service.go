package mybets

import (
	"context"
	"mytipster/internal/db"
	m "mytipster/models/mytips"

	"github.com/uptrace/bun"
)

func CreateTable(ctx context.Context, db *bun.DB) error {
	_, err := db.NewCreateTable().
		Model((*m.MyBets)(nil)).IfNotExists().
		Exec(ctx)
	return err

}

func getPickedByDate(date string, items []m.TipsDaily) error {

	ctx := context.Background()
	db := db.WithContext(ctx)

	result, err := FilterPredictionByDate(date, items)

	_, err = db.NewSelect().Model(&result).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func InsertPicked(items []m.TipsDaily) error {
	ctx := context.Background()
	db := db.WithContext(ctx)

	_, err := db.NewInsert().Model(&items).Exec(ctx)

	if err != nil {
		return err
	}
	return nil
}

func UpdatePicked(id string, items []m.TipsDaily, db *bun.DB) error {
	ctx := context.Background()

	item, err := FindId(id, items)
	if err != nil {
		return err
	}
	_, err = db.NewUpdate().Model(item).Set("picked = ?", true).Exec(ctx)
	return err
}

func DeletePicked(id string, items []m.TipsDaily, db *bun.DB) error {
	ctx := context.Background()
	item, err := FindId(id, items)

	if err != nil {
		return err
	}

	_, err = db.NewDelete().Model(item).Where("id = ?", item.ID).Exec(ctx)
	return err
}
