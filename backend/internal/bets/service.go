package bets

import (
	"context"
	"encoding/json"
	analytic_module "mytipster/models/analytic"
	m "mytipster/models/analytic"
	bets_models "mytipster/models/bets"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

func CreateTable(ctx context.Context, db *bun.DB) error {
	_, err := db.NewCreateTable().
		Model((*bets_models.Bets)(nil)).IfNotExists().
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
			TipsAnalyticsID: analyticsID,
			BetPick: bets_models.Bets{
				Handicap: fx.Handicap,
				Team:     fx.Team,
				Odds:     fx.Odds,
				Stake:    fx.Stake,
				Result:   fx.Result,
				Amount:   fx.Amount,
				Profit:   fx.Profit,
				Comments: fx.Comments,
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

func UpdateMyBets(
	id string,
	body mybets_models.BetPickIn,
	db *bun.DB,
	ctx context.Context,
) error {

	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	// แปลง body → map (เฉพาะ field ที่ส่งมา)
	payload := map[string]any{}

	if body.Handicap != "" {
		payload["handicap"] = body.Handicap
	}
	if body.Team != "" {
		payload["team"] = body.Team
	}
	if body.Odds != "" {
		payload["odds"] = body.Odds
	}
	if body.Stake != "" {
		payload["stake"] = body.Stake
	}
	if body.Result != "" {
		payload["result"] = body.Result
	}
	if body.Amount != 0 {
		payload["amount"] = body.Amount
	}
	if body.Profit != 0 {
		payload["profit"] = body.Profit
	}
	if body.Note != "" {
		payload["note"] = body.Note
	}

	// ❗ ต้องมีอย่างน้อย 1 field
	if len(payload) == 0 {
		return nil
	}

	// ✅ merge jsonb (ไม่ overwrite field ที่ไม่ส่งมา)
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	_, err = db.NewUpdate().
		Table("my-bets").
		Set("bet_pick = bet_pick || ?::jsonb", string(payloadBytes)).
		Where("id = ?", uid).
		Exec(ctx)
	return err
}

func DeletePicked(id string, items []analytic_module.MyBets, db *bun.DB, ctx context.Context) error {
	return nil
}
