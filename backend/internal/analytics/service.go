package analytics

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mytipster/internal/fixtures"
	analytic_module "mytipster/models/analytic"
	m "mytipster/models/analytic"
	fixture_module "mytipster/models/fixture"
	pred "mytipster/models/prediction"
	"strconv"

	"github.com/uptrace/bun"
)

// AnalyticService คือ interface (สัญญา / abstraction) abstract class
type AnalyticService interface {
	InsertManual(ctx context.Context, item *m.MyAnalytics) error
	InsertMany(ctx context.Context, items []m.MyAnalytics) error
	PredictionByDay(ctx context.Context, date string) ([]m.MyAnalytics, error)
	naWinTaTips(ctx context.Context, id string) (*pred.NaWinTatips, error)
	MatchResult(ctx context.Context, date string) ([]m.UpdateFixtureResultDTO, error)
}

// constructor / factory
// (เชื่อม interface ↔ struct) คือค่า interface
func NewAnalyticService(db *bun.DB) AnalyticService {
	return &analyticsService{
		db: db,
	}
}

// inform receiver type
// คือ concrete implementation
type analyticsService struct {
	db *bun.DB
}

func (s *analyticsService) InsertManual(ctx context.Context, item *m.MyAnalytics) error {
	_, err := s.db.NewInsert().Model(item).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *analyticsService) InsertMany(ctx context.Context, items []m.MyAnalytics) error {
	var filtered []m.MyAnalytics

	for _, item := range items {
		if item.FormLeagueHomeCount < 5 {
			continue
		}
		filtered = append(filtered, item)
	}

	if len(filtered) == 0 {
		fmt.Println("No records passed the filter, nothing to insert")
		return nil
	}

	_, err := s.db.NewInsert().Model(&filtered).Exec(ctx)
	if err != nil {
		log.Fatalf("insert many error %v", err)
	}

	fmt.Printf("✅ Inserted %d records\n", len(filtered))
	return nil
}

func (s *analyticsService) PredictionByDay(ctx context.Context, date string) ([]m.MyAnalytics, error) {
	var result []m.MyAnalytics
	err := s.db.NewSelect().Model(&result).Where("date = ?", date).Scan(ctx)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *analyticsService) naWinTaTips(ctx context.Context, fixtureId string) (*pred.NaWinTatips, error) {

	var analytic analytic_module.MyAnalytics

	data, err := fixtures.QueryPrediction(fixtureId)
	if err != nil {
		return nil, err
	}
	// 2️⃣ marshal prediction → jsonb
	raw, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	transform, err := strconv.Atoi(fixtureId)

	// 3 create model
	tip := &pred.NaWinTatips{
		TipsAnalyticsID: analytic.ID,
		FixtureID:       transform,
		Payload:         raw,
	}

	// 4.
	_, err = s.db.NewInsert().Model(tip).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return tip, err

}

func (s *analyticsService) MatchResult(ctx context.Context, date string) ([]m.UpdateFixtureResultDTO, error) {

	fixtures, err := fixtures.QueryFixtureDate(date)
	if err != nil {
		return nil, err
	}
	// เก็บ fixtureId
	fixtureMap := make(map[int]fixture_module.Response)

	// merge
	for _, fx := range fixtures {
		fixtureMap[fx.Fixture.ID] = fx
	}

	predictions, err := s.PredictionByDay(ctx, date)
	if err != nil {
		return nil, err
	}
	results := make([]m.UpdateFixtureResultDTO, 0, len(fixtures))

	for _, p := range predictions {
		fx, ok := fixtureMap[p.FixtureID]
		if !ok {
			continue
		}

		home := 0
		away := 0
		if fx.Goals.Home != nil {
			home = *fx.Goals.Home
		}
		if fx.Goals.Away != nil {
			away = *fx.Goals.Away
		}
		results = append(results, m.UpdateFixtureResultDTO{
			FixtureID:   fx.Fixture.ID,
			MatchFinish: fx.Fixture.Status.Long,
			MatchResult: fmt.Sprintf("%d-%d", home, away),
		})
	}
	return results, nil
}
