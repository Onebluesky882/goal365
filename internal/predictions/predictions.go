package predictions

import (
	"fmt"
	"mytipster/api"
	m "mytipster/models/prediction"
)

func Service(id string) (*m.PredictionsRoot, error) {

	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/predictions?fixture=%s", id)
	resp, err := api.ApiFootball[m.PredictionsRoot]("GET", url)
	if err != nil {
		return &m.PredictionsRoot{}, err
	}
	return resp, nil
}
