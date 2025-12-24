package api

import models "mytipster/models/fixture"

func Fixtures(method, url string) (*models.RootFixtureREsponse, error) {

	return ApiFootball[models.RootFixtureREsponse](method, url)

}
