package predictions

import (
	"context"
	"fmt"
	"mytipster/internal/db"
	"mytipster/internal/db/analytics"
	"mytipster/internal/fixtures/service"
	m "mytipster/models/fixture"

	"github.com/gofiber/fiber/v2"
)

func Service(c *fiber.Ctx) error {
	//1386834
	fixture := c.Query("fixture")
	resp, err := Predictions(fixture)

	if err != nil {
		return err
	}
	return c.JSON(resp)

}

func Predictions(fixtureId string) (*m.RootFixtureAnalytics, error) {

	ctx := context.Background()
	db, err := db.NewDB()
	if err != nil {
		return nil, err
	}

	results := &m.RootFixtureAnalytics{
		Items: m.FixtureAnalytics{},
	}

	// ดึง prediction
	pred, err := service.QueryPrediction(fixtureId)
	if err != nil || pred == nil {
		return nil, fmt.Errorf("prediction not found")
	}

	// ดึง fixture
	fx, err := service.QueryFixtureId(fixtureId)
	if err != nil {
		return nil, err
	}
	home := 0
	away := 0

	if fx.Goals.Home != nil {
		home = *fx.Goals.Home
	}
	if fx.Goals.Away != nil {
		away = *fx.Goals.Away
	}

	item := m.FixtureAnalytics{
		FixtureID:       fx.Fixture.ID,
		Date:            fx.Fixture.Date,
		Country:         fx.League.Country,
		League:          fx.League.Name,
		Home:            pred.Teams.Home.Name,
		Away:            pred.Teams.Away.Name,
		HomeScore:       pred.Teams.Home.Last5.Form,
		AwayScore:       pred.Teams.Away.Last5.Form,
		MatchFinish:     fx.Fixture.Status.Long,
		MatchResult:     fmt.Sprintf("%d-%d", home, away),
		HomeFormScore14: FormScore(14, pred.Teams.Home.League.Form),
		AwayFormScore14: FormScore(14, pred.Teams.Away.League.Form),
		HomeFormScore12: FormScore(12, pred.Teams.Home.League.Form),
		AwayFormScore12: FormScore(12, pred.Teams.Away.League.Form),
		HomeFormScore10: FormScore(10, pred.Teams.Home.League.Form),
		AwayFormScore10: FormScore(10, pred.Teams.Away.League.Form),
		HomeFormScore7:  FormScore(7, pred.Teams.Home.League.Form),
		AwayFormScore7:  FormScore(7, pred.Teams.Away.League.Form),
		HomeFormScore5:  FormScore(5, pred.Teams.Home.League.Form),
		AwayFormScore5:  FormScore(5, pred.Teams.Away.League.Form),
		BetPick: m.BetPick{
			Odds:   "0.5",
			Picked: "derby",
			Stake:  "", // win or lose
		},
	}

	if err = analytics.InsertData(ctx, db, &item); err != nil {
		return nil, err
	}

	results = &m.RootFixtureAnalytics{
		Items: item,
	}

	return results, nil

}

func FormScore(match int, form string) int {

	// set
	score := 0

	// running
	count := 0

	for i := len(form) - 1; i >= 0 && count < match; i-- {
		switch form[i] {
		case 'W':
			score += 3
		case 'D':
			score += 1
		case 'L':
		}
		count++
	}
	return score
}
