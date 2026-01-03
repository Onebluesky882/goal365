package mybets

import (
	"context"
	analytic_module "mytipster/models/analytic"
	m "mytipster/models/analytic"
	mybets_models "mytipster/models/mybets"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func CreateTable(ctx context.Context, db *bun.DB) error {
	_, err := db.NewCreateTable().
		Model((*m.MyBets)(nil)).IfNotExists().
		Exec(ctx)
	return err

}

func GetPickedByDate(date string, items []m.MyAnalytics, ctx context.Context, db *bun.DB) error {

	result, err := FilterPredictionByDate(date, items)

	_, err = db.NewSelect().Model(&result).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func InsertPicked(items []mybets_models.BetPickIn, analyticsID uuid.UUID, db *bun.DB, ctx context.Context) error {
	results := make([]analytic_module.MyBets, 0, len(items))
	for _, fx := range items {
		results = append(results, m.MyBets{
			TipsAnalyticsID: analyticsID,
			BetPick: m.BetPick{
				Picked: fx.Picked,
				Team:   fx.Team,
				Odds:   fx.Odds,
				Stake:  fx.Stake,
			},
		})
	}
	if len(results) == 0 {
		return nil
	}

	_, err := db.NewInsert().Model(&results).Returning("id").Exec(ctx)

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
