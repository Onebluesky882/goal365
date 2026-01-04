package analytic_module

import (
	odds_models "mytipster/models/odds"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type RootMyTipsAnalytics struct {
	Count int           `json:"count"`
	Items []MyAnalytics `json:"items"`
}

type MyBets struct {
	bun.BaseModel `bun:"table:my-bets,alias:mb"`

	ID uuid.UUID `bun:",pk,type:uuid,default:gen_random_uuid()" json:"id"`

	// ✅ FK ที่จำเป็น
	TipsAnalyticsID uuid.UUID `bun:"tips_analytics_id,type:uuid,notnull"`

	// belongs-to
	TipsAnalytics *MyAnalytics `bun:"rel:belongs-to,join:tips_analytics_id=id"`

	BetPick BetPick `bun:"bet_pick,type:jsonb"`
}

type TipsDaily struct {
	bun.BaseModel `bun:"table:tips-daily,alias:td"`

	// ✅ FK ที่จำเป็น
	ID uuid.UUID `bun:",pk,type:uuid,default:gen_random_uuid()" json:"id"`

	// ✅ FK (จำเป็นมาก)
	TipsAnalyticsID uuid.UUID `bun:"tips_analytics_id,type:uuid,notnull"`
	// belongs-to
	TipsAnalytics *MyAnalytics `bun:"rel:belongs-to,join:tips_analytics_id=id"`

	BetPick BetPick `bun:"bet_pick,type:jsonb"`
}

type MyAnalytics struct {
	bun.BaseModel `bun:"table:my-analytics,alias:ma"`

	ID                  uuid.UUID       `bun:",pk,type:uuid,default:gen_random_uuid()" json:"id"`
	FixtureID           int             `bun:"fixture_id,notnull" json:"fixture_id"`
	Date                string          `bun:"date" json:"date"`
	League              string          `bun:"league" json:"league"`
	TimeStamp           string          `bun:"timestamp" json:"timestamp"`
	Country             string          `bun:"country" json:"country"`
	Home                string          `bun:"team_home" json:"home"`
	Away                string          `bun:"team_away" json:"away"`
	MatchFinish         string          `bun:"match_finish" json:"match_finish"`
	Handicap            odds_models.Bet `bun:"handicap,type:jsonb" json:"handicap"`
	FormLeagueHomeCount int             `bun:"form_league_home_count" json:"form_league_home_count"`
	FormLeagueAwayCount int             `bun:"form_league_away_count" json:"form_league_away_count"`
	HomeFormScore14     int             `bun:"home_form_14" json:"home_form_14"`
	AwayFormScore14     int             `bun:"away_form_14" json:"away_form_14"`

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
	BetPick   BetPick      `bun:"bet_pick,type:jsonb" json:"bet_pick"`
	MyBets    []*MyBets    `bun:"rel:has-many,join:id=tips_analytics_id"`
	TipsDaily []*TipsDaily `bun:"rel:has-many,join:id=tips_analytics_id"`
}

type BetPick struct {
	Handicap string `json:"handicap"`
	Team     string `json:"team"`
	Odds     string `json:"odds"`
	Stake    string `json:"stake"`
	Result   string `json:"result"`
	Amount   int    `json:"amount"`
	Profit   int    `json:"profit"`
	Note     string `json:"note"`
}
type UpdateFixtureResultDTO struct {
	FixtureID   int    `json:"fixture_id"`
	MatchFinish string `json:"match_finish"`
	MatchResult string `json:"match_result"`
}
