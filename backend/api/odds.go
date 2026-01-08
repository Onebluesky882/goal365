package api

import "mytipster/models"

func Odds(method string, url string) (*models.RootOdds, error) {

	return ApiFootball[models.RootOdds](method, url)
}
