package api

import (
	odds_models "mytipster/models/odds"
)

func Odds(method string, url string) (*odds_models.RootOdds, error) {
	return ApiFootball[odds_models.RootOdds](method, url)
}
