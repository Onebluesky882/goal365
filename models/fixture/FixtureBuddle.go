package fixture_module

import (
	odds "mytipster/models/odds"
	prediction "mytipster/models/prediction"
)

type RootFixtureBundle struct {
	Items []FixtureBuddle
}
type FixtureBuddle struct {
	FixtureID   int
	Fixture     *Response
	Predictions *prediction.PredictionResponse
	Bookmaker   map[int][]odds.Bet
	Team        string
	Odds        string
	Result      string
	Picked      bool
}
