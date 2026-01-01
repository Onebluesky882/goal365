package mytips

import (
	"context"
	"fmt"
	"mytipster/internal/db"
	m "mytipster/models/mytips"

	"github.com/uptrace/bun"
)

func CreateTable(ctx context.Context, db *bun.DB) error {
	_, err := db.NewCreateTable().
		Model((*m.MyTipsAnalytics)(nil)).IfNotExists().
		Exec(ctx)
	return err

}

func UpdateFixtureResult(req []m.UpdateFixtureResultDTO) error {

	ctx := context.Background()
	db := db.WithContext(ctx)
	for _, v := range req {
		_, err := db.NewUpdate().
			Model((*m.MyTipsAnalytics)(nil)).
			Set("match_finish = ?", v.MatchFinish).
			Set("match_result = ?", v.MatchResult).
			Where("fixture_id = ?", v.FixtureID).
			Exec(ctx)
		if err != nil {
			return fmt.Errorf(
				"update failed fixture_id=%d: %w",
				v.FixtureID,
				err,
			)
		}
	}

	return nil
}
