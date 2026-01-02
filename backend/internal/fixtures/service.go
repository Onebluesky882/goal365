package fixtures

import (
	"fmt"
	"mytipster/api"
	"mytipster/lib"
	m "mytipster/models/fixture"
	odds_models "mytipster/models/odds"
	prediction_models "mytipster/models/prediction"
)

// get ids

func GetIds(date string) ([]int, error) {
	var result = make([]int, 0)

	resp, err := QueryFixtureDate(date)
	if err != nil {
		return nil, err
	}

	for _, id := range resp {
		result = append(result, id.Fixture.ID)
	}
	return result, nil
}

// fixture by id
func QueryFixtureId(id string) (*m.Response, error) {

	url := fmt.Sprintf(
		"https://v3.football.api-sports.io/fixtures?id=%s",
		id,
	)

	resp, err := api.Fixtures("GET", url)
	if err != nil {
		return nil, err
	}

	if len(resp.Response) == 0 {
		return nil, fmt.Errorf("no fixture found in api response")
	}

	return &resp.Response[0], nil
}

// odds FilterAsianHandicap

func QueryFixtureOdds(id string) (map[int][]odds_models.Bet, error) {

	url := fmt.Sprintf(
		"https://v3.football.api-sports.io/odds?fixture=%s",
		id,
	)

	resp, err := api.Odds("GET", url)
	if err != nil {
		return nil, err
	}

	if len(resp.Response) == 0 {
		return nil, fmt.Errorf("no prediction found for fixture %s", id)
	}
	result := lib.FilterAsianHandicap(resp, "Bet365")

	return result, err
}

// fixture date
func QueryFixtureDate(date string) ([]m.Response, error) {

	url := fmt.Sprintf(
		"https://v3.football.api-sports.io/fixtures?id=%s",
		date,
	)

	resp, err := api.Fixtures("GET", url)
	if err != nil {
		return nil, err
	}

	if len(resp.Response) == 0 {
		return nil, fmt.Errorf("no fixtures found for date %s", date)
	}

	return resp.Response, nil
}

// predictions

func QueryPrediction(id string) (*prediction_models.PredictionResponse, error) {

	url := fmt.Sprintf(
		"https://v3.football.api-sports.io/predictions?fixture=%s",
		id,
	)

	resp, err := api.Predictions("GET", url)
	if err != nil {
		return nil, err
	}

	if len(resp.Response) == 0 {
		return nil, fmt.Errorf("no fixture found in api response")
	}
	return &resp.Response[0], nil
}

// my tips
func QueryMyTipsOdds(id string) (map[int][]odds_models.Bet, error) {

	url := fmt.Sprintf(
		"https://v3.football.api-sports.io/odds?fixture=%s",
		id,
	)

	resp, err := api.Odds("GET", url)
	if err != nil {
		return nil, err
	}

	result := lib.FilterOdds(resp, "Bet365")

	return result, err
}
