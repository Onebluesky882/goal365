package fixtures

import (
	"fmt"
	"mytipster/api"
	"mytipster/lib"
	m "mytipster/models"
)

// get ids

func GetIds(date string) ([]int, error) {
	var result = make([]int, 0)

	resp, err := QueryFixtureDate(date)
	if err != nil {
		return nil, err
	}

	for _, id := range resp {
		result = append(result, id.FixtureInfo.ID)
	}
	return result, nil
}

func GetIdsWithFilterCountry(date string, countries []string) ([]int, error) {
	var result = make([]int, 0)

	resp, err := QueryFixtureDate(date)
	if err != nil {
		return nil, err
	}

	for _, fx := range resp {
		if len(countries) == 0 || lib.FilterCountry(&fx, countries) {
			result = append(result, fx.FixtureInfo.ID)
		}
	}
	return result, nil
}

// fixture by id
func QueryFixtureId(id string) (*m.FixtureResponse, error) {

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

func QueryFixtureOdds(id string) ([]m.Bookmaker, error) {
	url := fmt.Sprintf(
		"https://v3.football.api-sports.io/odds?fixture=%s",
		id,
	)

	resp, err := api.Odds("GET", url)
	if err != nil {
		return nil, err
	}

	if len(resp.Response) == 0 {
		return nil, fmt.Errorf("no odds found for fixture %s", id)
	}

	firstResponse := resp.Response[0]

	if len(firstResponse.Bookmakers) == 0 {
		return nil, fmt.Errorf("no bookmakers found for fixture %s", id)
	}

	return firstResponse.Bookmakers, nil
}

// fixture date
func QueryFixtureDate(date string) ([]m.FixtureResponse, error) {
	if date == "" {
		return nil, fmt.Errorf("date is required (YYYY-MM-DD)")
	}
	url := fmt.Sprintf(
		"https://v3.football.api-sports.io/fixtures?date=%s",
		date,
	)

	resp, err := api.Fixtures("GET", url)
	if err != nil {
		return nil, err
	}

	if len(resp.Response) == 0 {
		return nil, fmt.Errorf("no fixtures found for date %s", date)
	}

	fmt.Println(" resp.Response", len(resp.Response))

	return resp.Response, nil
}

// predictions

func QueryPrediction(id string) (*m.PredictionResponse, error) {

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
func QueryMyTipsOdds(id string) (map[int][]m.Bet, error) {

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
