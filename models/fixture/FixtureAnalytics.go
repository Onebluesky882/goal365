package fixture_module

type RootFixtureAnalytics struct {
	Items []FixtureAnalytics `json:"items"`
}

type FixtureAnalytics struct {
	FixtureID int    `json:"fixture_id"`
	Date      string `json:"date"`
	TeamHome  string `json:"team_home"`
	TeamAway  string `json:"team_away"`

	HomeFormScore14 int `json:"home_form_14"`
	AwayFormScore14 int `json:"away_form_14"`

	HomeFormScore12 int `json:"home_form_12"`
	AwayFormScore12 int `json:"away_form_12"`

	HomeFormScore10 int `json:"home_form_10"`
	AwayFormScore10 int `json:"away_form_10"`

	HomeFormScore7 int `json:"home_form_7"`
	AwayFormScore7 int `json:"away_form_7"`

	HomeFormScore5 int `json:"home_form_5"`
	AwayFormScore5 int `json:"away_form_5"`

	HomeScore   string `json:"home_score"`
	AwayScore   string `json:"away_score"`
	MatchResult string `json:"match_result"`
	BetPick     BetPick
}

type BetPick struct {
	Handicap string
	Picked   string
	Stake    string
}
