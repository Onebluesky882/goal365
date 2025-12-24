package odds

import (
	"fmt"
	"mytipster/api"
	"mytipster/lib"
	m "mytipster/models/odds"
	odds_models "mytipster/models/odds"
)

func Service(id string) (map[int][]odds_models.Bet, error) {

	url := fmt.Sprintf("https://api-football-v1.p.rapidapi.com/v3/odds?fixture=%s", id)

	resp, err := api.ApiFootball[m.RootOdds]("GET", url)

	if err != nil {
		return nil, err
	}

	result := lib.FilterBookMarket(resp, "Bet365")
	return result, nil
}
