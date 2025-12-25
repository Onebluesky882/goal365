package fixture_module

import (
	odds "mytipster/models/odds"
	prediction "mytipster/models/prediction"
)

type FixturePredictionBundle struct {
	Items []FixturePrediction
}
type FixturePrediction struct {
	FixtureID   int
	Fixture     *Response
	Predictions *prediction.PredictionResponse
	Bookmaker   map[int][]odds.Bet
	Team        string
	Result      string
	Picked      bool
}
