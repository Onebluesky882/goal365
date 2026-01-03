package mybets

import (
	"context"
	"mytipster/internal/db"
	"mytipster/lib/common"
	m "mytipster/models/mytips"
	"strconv"

	"github.com/uptrace/bun"
)

func FilterPredictionByDate(date string, items []m.TipsDaily) ([]m.TipsDaily, error) {

	var result []m.TipsDaily
	for _, item := range items {
		ts, err := strconv.ParseInt(item.Date, 10, 64)

		if err != nil {
			return nil, err
		}
		format := common.TimestampUTCDate(ts)
		if format != date {
			continue
		}
		result = append(result, item)
	}
	return result, nil

}

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

// func UpdatePicked(item *m.TipsDaily) error {
// 	ctx := context.Background()
// 	db := db.WithContext(ctx)

// }

// func delete(id int) error {
// 	ctx := context.Background()
// 	db := db.WithContext(ctx)

// }
