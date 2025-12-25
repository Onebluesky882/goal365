package fixtures

import (
	"fmt"
	"mytipster/api"
	m "mytipster/models/fixture"
)

func fetchFixtureResponse(id string) (*m.RootFixtureResponse, error) {
	url := fmt.Sprintf(
		"https://api-football-v1.p.rapidapi.com/v3/fixtures?id=%s",
		id,
	)

	return api.Fixtures("GET", url)
}

func QueryFixtureById(id string) (m.Fixture, error) {

	resp, err := fetchFixtureResponse(id)
	if err != nil {
		return m.Fixture{}, err
	}

	return resp.Response[0].Fixture, nil
}

func fixtureDataService(id string) (*m.Response, error) {
	resq, err := fetchFixtureResponse(id)
	if err != nil {
		return nil, err
	}
	return &resq.Response[0], nil

}
