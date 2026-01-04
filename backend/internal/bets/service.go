package bets

import (
	"context"
	analytic_module "mytipster/models/analytic"
	m "mytipster/models/analytic"
	bets_models "mytipster/models/bets"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func GetBetListsByDate(date string, items []m.MyAnalytics, db *bun.DB, ctx context.Context) ([]bets_models.Bets, error) {

	// --- 1️⃣ query analytics ของวันนั้นโดยตรง ---
	var analyticsItems []analytic_module.MyAnalytics
	if err := db.NewSelect().
		Model(&analyticsItems).
		Where("date = ?", date).
		Scan(ctx); err != nil {
		return nil, err
	}
	if len(analyticsItems) == 0 {
		return []bets_models.Bets{}, nil
	}

	// --- 2️⃣ เก็บ IDs ของ analytics ที่ match ---
	analyticsIDs := make([]uuid.UUID, len(analyticsItems))
	for i, a := range analyticsItems {
		analyticsIDs[i] = a.ID
	}

	// --- 3️⃣ query my-bets ที่ relation กับ analytics ---
	var bets []bets_models.Bets
	if err := db.NewSelect().
		Model(&bets).
		Relation("TipsAnalytics").
		Where("tips_analytics_id IN (?)", bun.In(analyticsIDs)).
		Scan(ctx); err != nil {
		return nil, err
	}

	return bets, nil
}

func InsertPicked(items []bets_models.Bets, analyticsID uuid.UUID, db *bun.DB, ctx context.Context) error {
	results := make([]bets_models.Bets, 0, len(items))
	for _, fx := range items {
		results = append(results, bets_models.Bets{
			ID:              fx.ID,
			TipsAnalyticsID: analyticsID,
			BaseModel:       fx.BaseModel,
			Handicap:        fx.Handicap,
			Team:            fx.Team,
			Odds:            fx.Odds,
			Stake:           fx.Stake,
			Result:          fx.Result,
			Amount:          fx.Amount,
			Profit:          fx.Profit,
			Comments:        fx.Comments,
		},
		)

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

func UpdateMyBets(
	id string,
	body bets_models.Bets,
	db *bun.DB,
	ctx context.Context,
) error {

	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	_, err = db.NewUpdate().
		Table("my-bets").
		Where("id = ?", uid).
		Exec(ctx)
	return err
}
