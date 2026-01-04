package matchresults

import (
	"context"
	"fmt"
	"mytipster/internal/analytics"
	"mytipster/internal/fixtures"
	fixture_module "mytipster/models/fixture"
	m "mytipster/models/match_results"
)

func MatchResult(service analytics.AnalyticService, date string) ([]m.UpdateFixtureResultDTO, error) {
	ctx := context.Background()
	predictions, err := service.PredictionByDay(ctx, date)
	if err != nil {
		return nil, err
	}

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

	results := make([]m.UpdateFixtureResultDTO, 0, len(predictions))

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
