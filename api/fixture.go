package api

import models "mytipster/models/fixture"

func Fixtures(method, url string) (*models.RootFixtureResponse, error) {
	return ApiFootball[models.RootFixtureResponse](method, url)
}
