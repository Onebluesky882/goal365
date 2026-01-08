package bets

import (
	"context"
	m "mytipster/models"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func GetBetListsByDate(date string, items []m.MyAnalytics, db *bun.DB, ctx context.Context) ([]m.Bets, error) {

	// --- 1️⃣ query analytics ของวันนั้นโดยตรง ---
	var analyticsItems []m.MyAnalytics
	if err := db.NewSelect().
		Model(&analyticsItems).
		Where("date = ?", date).
		Scan(ctx); err != nil {
		return nil, err
	}
	if len(analyticsItems) == 0 {
		return []m.Bets{}, nil
	}

	// --- 2️⃣ เก็บ IDs ของ analytics ที่ match ---
	analyticsIDs := make([]uuid.UUID, len(analyticsItems))
	for i, a := range analyticsItems {
		analyticsIDs[i] = a.ID
	}

	// --- 3️⃣ query my-bets ที่ relation กับ analytics ---
	var bets []m.Bets
	if err := db.NewSelect().
		Model(&bets).
		Relation("TipsAnalytics").
		Where("tips_analytics_id IN (?)", bun.In(analyticsIDs)).
		Scan(ctx); err != nil {
		return nil, err
	}

	return bets, nil
}

func InsertPicked(items []m.Bets, analyticsID uuid.UUID, db *bun.DB, ctx context.Context) error {
	results := make([]m.Bets, 0, len(items))
	for _, fx := range items {
		results = append(results, m.Bets{
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
	body m.Bets,
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
