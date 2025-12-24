package fixtureoddsprediction

import (
	odds_models "mytipster/models/odds"
	prediction_models "mytipster/models/prediction"
)



type Fixture365Bet struct {
	FixtureID   int
	Predictions []prediction_models.PredictionResponse
	Bookmaker   []odds_models.Bookmaker
}

func name() {

}
