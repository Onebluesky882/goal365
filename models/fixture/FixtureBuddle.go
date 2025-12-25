package fixture_module

import (
	odds "mytipster/models/odds"
	prediction "mytipster/models/prediction"
)

type FixturePredictionBundle struct {
	FixtureIDs []int
	Items      []FixturePrediction
}
type FixturePrediction struct {
	FixtureID   int
	Predictions []prediction.PredictionResponse
	Bookmaker   map[int][]odds.Bet
	Result      string
	Picked      bool
}
