package fixtures

import (
	"fmt"
	"mytipster/api"
	m "mytipster/models/fixture"
)

func QueryFixtureById(id string) (m.Fixture, error) {

	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/fixtures?id=%s", id)

	resp, err := api.Fixtures("GET", url)

	if err != nil {
		return m.Fixture{}, err
	}

	if len(resp.Response) == 0 {
		return m.Fixture{}, fmt.Errorf("no fixture found for id %s", id)
	}

	return resp.Response[0].Fixture, err
}
