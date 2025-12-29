package odds_models

type RootOdds struct {
	Get        string     `json:"get"`
	Parameters Parameters `json:"parameters"`
	Errors     any        `json:"errors"`
	Results    int        `json:"results"`
	Paging     Paging     `json:"paging"`
	Response   []Response `json:"response"`
	Info       string     `json:"info"`
}

type Parameters struct {
	Fixture string `json:"fixture"`
}

type Paging struct {
	Current int `json:"current"`
	Total   int `json:"total"`
}

type Response struct {
	League     League      `json:"league"`
	Fixture    Fixture     `json:"fixture"`
	Update     string      `json:"update"`
	Bookmakers []Bookmaker `json:"bookmakers"`
}

type League struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
	Logo    string `json:"logo"`
	Flag    string `json:"flag"`
	Season  int    `json:"season"`
}

type Fixture struct {
	ID        int    `json:"id"`
	Timezone  string `json:"timezone"`
	Date      string `json:"date"`
	Timestamp int64  `json:"timestamp"`
}

type Bookmaker struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Bets []Bet  `json:"bets"`
}

type Bet struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Values []Value `json:"values"`
}

type Value struct {
	Value any    `json:"value"`
	Odd   string `json:"odd"`
}

type OddsMap map[string][]Bet

type SimplifiedOdds struct {
	Data OddsMap
}
