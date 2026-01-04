package bets_models

import (
	analytic_module "mytipster/models/analytic"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type InsertPickedRequest struct {
	AnalyticsID uuid.UUID `json:"analytics_id"`
	Items       []Bets    `json:"items"`
}

type Bets struct {
	bun.BaseModel `bun:"table:my-bets,alias:mb"`

	ID uuid.UUID `bun:",pk,type:uuid,default:gen_random_uuid()" json:"id"`

	// ✅ FK ที่จำเป็น
	TipsAnalyticsID uuid.UUID `bun:"tips_analytics_id,type:uuid,notnull"`

	Handicap string `bun:"handicap" json:"handicap"`
	Team     string `bun:"team" json:"team"`
	Odds     string `bun:"odds" json:"odds"`
	Stake    string `bun:"stake" json:"stake"`
	Result   string `bun:"result" json:"result"`
	Amount   int    `bun:"amount" json:"amount"`
	Profit   int    `bun:"profit" json:"profit"`
	Comments string `bun:"comments" json:"comments"`

	// belongs-to
	TipsAnalytics *analytic_module.MyAnalytics `bun:"rel:belongs-to,join:tips_analytics_id=id"`
}
