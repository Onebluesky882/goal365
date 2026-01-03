package mybets

import (
	"context"
	"fmt"
	"mytipster/internal/db"
	"mytipster/lib"
	m "mytipster/models/analytic"
	odds_models "mytipster/models/odds"

	"github.com/uptrace/bun"
)

func CreateTable(ctx context.Context, db *bun.DB) error {
	_, err := db.NewCreateTable().
		Model((*m.MyBets)(nil)).IfNotExists().
		Exec(ctx)
	return err

}

func getPickedByDate(date string, items []m.MyAnalytics) error {

	ctx := context.Background()
	db := db.WithContext(ctx)

	result, err := FilterPredictionByDate(date, items)

	_, err = db.NewSelect().Model(&result).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func InsertPicked(date string, handicap odds_models.Bet, bet m.BetPick, items []m.MyAnalytics, db *bun.DB) error {
	ctx := context.Background()

	path := fmt.Sprintf("bin/%s/predictions.json", date)

	var results []m.MyBets

	temp, err := lib.ReadJson[[]m.MyBets](path)
	if err != nil {
		return err
	}

	for _, fx := range temp {

		for _, item := range items {

			if fx.ID == item.ID {
				prediction := m.MyBets{
					Handicap: handicap,
					BetPick:  bet,
				}
				results = append(results, prediction)
			}
		}
	}

	_, err = db.NewInsert().Model(&results).Exec(ctx)

	if err != nil {
		return err
	}
	return nil
}

func UpdatePicked(id string, items []m.MyAnalytics, db *bun.DB) error {
	ctx := context.Background()

	item, err := FindId(id, items)
	if err != nil {
		return err
	}
	_, err = db.NewUpdate().Model(item).Set("picked = ?", true).Exec(ctx)
	return err
}

func DeletePicked(id string, items []m.MyAnalytics, db *bun.DB) error {
	ctx := context.Background()
	item, err := FindId(id, items)

	if err != nil {
		return err
	}

	_, err = db.NewDelete().Model(item).Where("id = ?", item.ID).Exec(ctx)
	return err
}
