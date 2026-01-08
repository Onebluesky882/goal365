package api

import "mytipster/models"

func Fixtures(method, url string) (*models.RootFixtureResponse, error) {
	return ApiFootball[models.RootFixtureResponse](method, url)
}
