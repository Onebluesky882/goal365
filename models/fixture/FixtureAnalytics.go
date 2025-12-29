package fixture_module

import "github.com/uptrace/bun"

type RootFixtureAnalytics struct {
	Items []FixtureAnalytics `json:"items"`
}

type FixtureAnalytics struct {
	bun.BaseModel `bun:"table:nawin_analytics,alias:fa"`

	ID        int64  `bun:",pk,autoincrement" json:"-"`
	FixtureID int    `bun:"fixture_id,notnull" json:"fixture_id"`
	Date      string `bun:"date" json:"date"`

	TeamHome string `bun:"team_home" json:"team_home"`
	TeamAway string `bun:"team_away" json:"team_away"`

	HomeFormScore14 int `bun:"home_form_14" json:"home_form_14"`
	AwayFormScore14 int `bun:"away_form_14" json:"away_form_14"`

	HomeFormScore12 int `bun:"home_form_12" json:"home_form_12"`
	AwayFormScore12 int `bun:"away_form_12" json:"away_form_12"`

	HomeFormScore10 int `bun:"home_form_10" json:"home_form_10"`
	AwayFormScore10 int `bun:"away_form_10" json:"away_form_10"`

	HomeFormScore7 int `bun:"home_form_7" json:"home_form_7"`
	AwayFormScore7 int `bun:"away_form_7" json:"away_form_7"`

	HomeFormScore5 int `bun:"home_form_5" json:"home_form_5"`
	AwayFormScore5 int `bun:"away_form_5" json:"away_form_5"`

	HomeScore   string `bun:"home_score" json:"home_score"`
	AwayScore   string `bun:"away_score" json:"away_score"`
	MatchResult string `bun:"match_result" json:"match_result"`

	// ⭐ JSONB column
	BetPick BetPick `bun:"bet_pick,type:jsonb" json:"bet_pick"`
}

type BetPick struct {
	Handicap string `json:"handicap"`
	Picked   string `json:"picked"`
	Stake    string `json:"stake"`
}
