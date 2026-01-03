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

func GetBetListsByDate(date string, items []m.MyAnalytics, db *bun.DB, ctx context.Context) ([]analytic_module.MyBets, error) {

	// --- 1️⃣ query analytics ของวันนั้นโดยตรง ---
	var analyticsItems []analytic_module.MyAnalytics
	if err := db.NewSelect().
		Model(&analyticsItems).
		Where("date = ?", date).
		Scan(ctx); err != nil {
		return nil, err
	}
	if len(analyticsItems) == 0 {
		return []analytic_module.MyBets{}, nil
	}

	// --- 2️⃣ เก็บ IDs ของ analytics ที่ match ---
	analyticsIDs := make([]uuid.UUID, len(analyticsItems))
	for i, a := range analyticsItems {
		analyticsIDs[i] = a.ID
	}

	// --- 3️⃣ query my-bets ที่ relation กับ analytics ---
	var bets []analytic_module.MyBets
	if err := db.NewSelect().
		Model(&bets).
		Relation("TipsAnalytics").
		Where("tips_analytics_id IN (?)", bun.In(analyticsIDs)).
		Scan(ctx); err != nil {
		return nil, err
	}

	return bets, nil
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
