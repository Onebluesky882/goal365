package prediction_models

type PredictionsRoot struct {
	Get        string               `json:"get"`
	Parameters PredParameters       `json:"parameters"`
	Errors     any                  `json:"errors"`
	Results    int                  `json:"results"`
	Paging     Paging               `json:"paging"`
	Response   []PredictionResponse `json:"response"`
	Info       string               `json:"info"`
}

type PredParameters struct {
	Fixture string `json:"fixture"`
}

type Paging struct {
	Current int `json:"current"`
	Total   int `json:"total"`
}

// Main Response
type PredictionResponse struct {
	Predictions Predictions  `json:"predictions"`
	League      PredLeague   `json:"league"`
	Teams       PredTeams    `json:"teams"`
	Comparison  Comparison   `json:"comparison"`
	H2H         []H2HFixture `json:"h2h"`
}

// Predictions
type Predictions struct {
	Winner    Winner      `json:"winner"`
	WinOrDraw bool        `json:"win_or_draw"`
	UnderOver interface{} `json:"under_over"`
	Goals     PredGoals   `json:"goals"`
	Advice    string      `json:"advice"`
	Percent   Percent     `json:"percent"`
}

type Winner struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

type PredGoals struct {
	Home string `json:"home"`
	Away string `json:"away"`
}

type Percent struct {
	Home string `json:"home"`
	Draw string `json:"draw"`
	Away string `json:"away"`
}

// League
type PredLeague struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
	Logo    string `json:"logo"`
	Flag    string `json:"flag"`
	Season  int    `json:"season"`
}

// Teams
type PredTeams struct {
	Home PredTeamDetail `json:"home"`
	Away PredTeamDetail `json:"away"`
}

type PredTeamDetail struct {
	ID     int             `json:"id"`
	Name   string          `json:"name"`
	Logo   string          `json:"logo"`
	Last5  Last5           `json:"last_5"`
	League TeamLeagueStats `json:"league"`
}

type Last5 struct {
	Played int        `json:"played"`
	Form   string     `json:"form"`
	Att    string     `json:"att"`
	Def    string     `json:"def"`
	Goals  Last5Goals `json:"goals"`
}

type Last5Goals struct {
	For     GoalsStat `json:"for"`
	Against GoalsStat `json:"against"`
}

type GoalsStat struct {
	Total   int    `json:"total"`
	Average string `json:"average"`
}

type TeamLeagueStats struct {
	Form          string        `json:"form"`
	Fixtures      FixturesStats `json:"fixtures"`
	Goals         DetailedGoals `json:"goals"`
	Biggest       Biggest       `json:"biggest"`
	CleanSheet    StatBreakdown `json:"clean_sheet"`
	FailedToScore StatBreakdown `json:"failed_to_score"`
	Penalty       PenaltyStats  `json:"penalty"`
	Lineups       []Lineup      `json:"lineups"`
	Cards         Cards         `json:"cards"`
}

type FixturesStats struct {
	Played HomeAwayTotal `json:"played"`
	Wins   HomeAwayTotal `json:"wins"`
	Draws  HomeAwayTotal `json:"draws"`
	Loses  HomeAwayTotal `json:"loses"`
}

type HomeAwayTotal struct {
	Home  int `json:"home"`
	Away  int `json:"away"`
	Total int `json:"total"`
}

type DetailedGoals struct {
	For     GoalsDetail `json:"for"`
	Against GoalsDetail `json:"against"`
}

type GoalsDetail struct {
	Total     HomeAwayTotal `json:"total"`
	Average   HomeAwayAvg   `json:"average"`
	Minute    MinuteStats   `json:"minute"`
	UnderOver UnderOverMap  `json:"under_over"`
}

type HomeAwayAvg struct {
	Home  string `json:"home"`
	Away  string `json:"away"`
	Total string `json:"total"`
}

type MinuteStats struct {
	M0_15    MinuteStat `json:"0-15"`
	M16_30   MinuteStat `json:"16-30"`
	M31_45   MinuteStat `json:"31-45"`
	M46_60   MinuteStat `json:"46-60"`
	M61_75   MinuteStat `json:"61-75"`
	M76_90   MinuteStat `json:"76-90"`
	M91_105  MinuteStat `json:"91-105"`
	M106_120 MinuteStat `json:"106-120"`
}

type MinuteStat struct {
	Total      interface{} `json:"total"`
	Percentage interface{} `json:"percentage"`
}

type UnderOverMap struct {
	Under0_5 OverUnder `json:"0.5"`
	Under1_5 OverUnder `json:"1.5"`
	Under2_5 OverUnder `json:"2.5"`
	Under3_5 OverUnder `json:"3.5"`
	Under4_5 OverUnder `json:"4.5"`
}

type OverUnder struct {
	Over  int `json:"over"`
	Under int `json:"under"`
}

type Biggest struct {
	Streak StreakStats  `json:"streak"`
	Wins   HomeAwayStat `json:"wins"`
	Loses  HomeAwayStat `json:"loses"`
	Goals  BiggestGoals `json:"goals"`
}

type StreakStats struct {
	Wins  int `json:"wins"`
	Draws int `json:"draws"`
	Loses int `json:"loses"`
}

type HomeAwayStat struct {
	Home string `json:"home"`
	Away string `json:"away"`
}

type BiggestGoals struct {
	For     HomeAwayInt `json:"for"`
	Against HomeAwayInt `json:"against"`
}

type HomeAwayInt struct {
	Home int `json:"home"`
	Away int `json:"away"`
}

type StatBreakdown struct {
	Home  int `json:"home"`
	Away  int `json:"away"`
	Total int `json:"total"`
}

type PenaltyStats struct {
	Scored PenaltyDetail `json:"scored"`
	Missed PenaltyDetail `json:"missed"`
	Total  int           `json:"total"`
}

type PenaltyDetail struct {
	Total      int    `json:"total"`
	Percentage string `json:"percentage"`
}

type Lineup struct {
	Formation string `json:"formation"`
	Played    int    `json:"played"`
}

type Cards struct {
	Yellow MinuteStats `json:"yellow"`
	Red    MinuteStats `json:"red"`
}

// Comparison
type Comparison struct {
	Form                HomeAwayStat `json:"form"`
	Att                 HomeAwayStat `json:"att"`
	Def                 HomeAwayStat `json:"def"`
	PoissonDistribution HomeAwayStat `json:"poisson_distribution"`
	H2H                 HomeAwayStat `json:"h2h"`
	Goals               HomeAwayStat `json:"goals"`
	Total               HomeAwayStat `json:"total"`
}

// H2H Fixture
type H2HFixture struct {
	Fixture H2HFixtureDetail `json:"fixture"`
	League  H2HLeague        `json:"league"`
	Teams   H2HTeams         `json:"teams"`
	Goals   H2HGoals         `json:"goals"`
	Score   H2HScore         `json:"score"`
}

type H2HFixtureDetail struct {
	ID        int        `json:"id"`
	Referee   *string    `json:"referee,omitempty"`
	Timezone  string     `json:"timezone"`
	Date      string     `json:"date"`
	Timestamp int64      `json:"timestamp"`
	Periods   H2HPeriods `json:"periods"`
	Venue     H2HVenue   `json:"venue"`
	Status    H2HStatus  `json:"status"`
}

type H2HPeriods struct {
	First  int `json:"first"`
	Second int `json:"second"`
}

type H2HVenue struct {
	ID   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	City string  `json:"city"`
}

type H2HStatus struct {
	Long    string `json:"long"`
	Short   string `json:"short"`
	Elapsed int    `json:"elapsed"`
	Extra   *int   `json:"extra,omitempty"`
}

type H2HLeague struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Country   string  `json:"country"`
	Logo      string  `json:"logo"`
	Flag      *string `json:"flag,omitempty"`
	Season    int     `json:"season"`
	Round     string  `json:"round"`
	Standings bool    `json:"standings"`
}

type H2HTeams struct {
	Home H2HTeam `json:"home"`
	Away H2HTeam `json:"away"`
}

type H2HTeam struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Logo   string `json:"logo"`
	Winner bool   `json:"winner"`
}

type H2HGoals struct {
	Home int `json:"home"`
	Away int `json:"away"`
}

type H2HScore struct {
	Halftime  H2HScoreDetail `json:"halftime"`
	Fulltime  H2HScoreDetail `json:"fulltime"`
	Extratime H2HScoreDetail `json:"extratime"`
	Penalty   H2HScoreDetail `json:"penalty"`
}

type H2HScoreDetail struct {
	Home interface{} `json:"home"`
	Away interface{} `json:"away"`
}
