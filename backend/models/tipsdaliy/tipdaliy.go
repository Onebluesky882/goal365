package tipsdaliy_models

import (
	analytic_module "mytipster/models/analytic"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type TipsDaily struct {
	bun.BaseModel `bun:"table:tips-daily,alias:td"`

	// ✅ FK ที่จำเป็น
	ID uuid.UUID `bun:",pk,type:uuid,default:gen_random_uuid()" json:"id"`

	// ✅ FK (จำเป็นมาก)
	TipsAnalyticsID uuid.UUID `bun:"tips_analytics_id,type:uuid,notnull"`
	// belongs-to
	TipsAnalytics *analytic_module.MyAnalytics `bun:"rel:belongs-to,join:tips_analytics_id=id"`
}
