package fixtures

import (
	"context"
	"fmt"
	"mytipster/internal/db"
	"mytipster/internal/db/analytics"
	"mytipster/internal/fixtures/service"
	m "mytipster/models/fixture"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// pred "mytipster/models/prediction"

/*
get predictions by id


output want
FixtureID       int
TeamHome        string
	TeamAway        string
	HomeFormScore14 int
	AwayFormScore14 int
	HomeFormScore12 int
	AwayFormScore12 int
	HomeFormScore10 int
	AwayFormScore10 int
	HomeFormScore7  int
	AwayFormScore7  int
	HomeFormScore5  int
	AwayFormScore5  int    // 12
	Last5form       string // "54%"
	MatchResult     string // "2-1"
	Picked          string // home
	Handicap        string // 0.5
*/

func Analytics(data map[int]m.BetPick) (*m.RootFixtureAnalytics, error) {

	ctx := context.Background()
	db, err := db.NewDB()
	if err != nil {
		return nil, err
	}

	result := &m.RootFixtureAnalytics{
		Items: []m.FixtureAnalytics{},
	}
	for fixtureID, betPick := range data {
		idStr := strconv.Itoa(fixtureID)
		pred, err := service.QueryPrediction(idStr)

		fixture, err := service.QueryFixtureId(idStr)
		if err != nil {
			continue // ข้าม fixture นี้
		}
		// === ตัวอย่าง filter ===
		if pred == nil {
			continue
		}

		home := 0
		away := 0

		if fixture.Goals.Home != nil {
			home = *fixture.Goals.Home
		}
		if fixture.Goals.Away != nil {
			away = *fixture.Goals.Away
		}

		item := m.FixtureAnalytics{
			FixtureID:       fixtureID,
			Date:            fixture.Fixture.Date,
			TeamHome:        pred.Teams.Home.Name,
			TeamAway:        pred.Teams.Away.Name,
			HomeScore:       pred.Teams.Home.Last5.Form,
			AwayScore:       pred.Teams.Away.Last5.Form,
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
			BetPick:         betPick,
		}

		result.Items = append(result.Items, item)

		err = analytics.InsertData(ctx, db, &item)
		time.Sleep(2000 * time.Millisecond)
		fmt.Println("current count:", len(result.Items))
	}
	return nil, nil

}

func Service(c *fiber.Ctx) error {
	resp, err := Analytics(service.Data)

	if err != nil {
		return err
	}
	return c.JSON(resp)

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
