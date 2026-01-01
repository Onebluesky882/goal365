package predictions

import (
	"fmt"
	"mytipster/internal/fixtures/service"
	"mytipster/lib"
	"mytipster/lib/common"
	m "mytipster/models/mytips"

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

func Predictions(fixtureId string) (*m.MyTipsAnalytics, error) {

	// ดึง prediction
	pred, err := service.QueryPrediction(fixtureId)
	if err != nil {
		return nil, fmt.Errorf("QueryPrediction error: %w", err)
	}
	if pred == nil {
		return nil, fmt.Errorf("prediction is nil fixtureId=%s", fixtureId)
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

	item := &m.MyTipsAnalytics{
		FixtureID:           fx.Fixture.ID,
		Date:                common.TimestampUTCDate(fx.Fixture.Timestamp),
		TimeStamp:           common.Timestamp(fx.Fixture.Timestamp),
		Country:             fx.League.Country,
		League:              fx.League.Name,
		Home:                pred.Teams.Home.Name,
		Away:                pred.Teams.Away.Name,
		HomeScore:           pred.Teams.Home.Last5.Form,
		AwayScore:           pred.Teams.Away.Last5.Form,
		FormLeagueHomeCount: len(pred.Teams.Home.League.Form),
		FormLeagueAwayCount: len(pred.Teams.Away.League.Form),
		HomeFormScore14:     lib.FormScore(14, pred.Teams.Home.League.Form),
		AwayFormScore14:     lib.FormScore(14, pred.Teams.Away.League.Form),
		HomeFormScore12:     lib.FormScore(12, pred.Teams.Home.League.Form),
		AwayFormScore12:     lib.FormScore(12, pred.Teams.Away.League.Form),
		HomeFormScore10:     lib.FormScore(10, pred.Teams.Home.League.Form),
		AwayFormScore10:     lib.FormScore(10, pred.Teams.Away.League.Form),
		HomeFormScore7:      lib.FormScore(7, pred.Teams.Home.League.Form),
		AwayFormScore7:      lib.FormScore(7, pred.Teams.Away.League.Form),
		HomeFormScore5:      lib.FormScore(5, pred.Teams.Home.League.Form),
		AwayFormScore5:      lib.FormScore(5, pred.Teams.Away.League.Form),
		MatchFinish:         fx.Fixture.Status.Long,
		MatchResult:         fmt.Sprintf("%d-%d", home, away),
		BetPick: m.BetPick{
			Picked: "",
			Team:   "",
			Odds:   "",
			Stake:  "",
		},
	}

	return item, nil

}
