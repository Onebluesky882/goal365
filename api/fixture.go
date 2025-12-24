package api

import "mytipster/models"

func Fixtures(method,  url string) (*models.RootFixtureREsponse , error) {

	return  ApiFootball[models.RootFixtureREsponse]( method, url)
	
}
