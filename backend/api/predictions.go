package api

import "mytipster/models"

func Predictions(method, url string) (*models.PredictionsRoot, error) {
	return ApiFootball[models.PredictionsRoot](method, url)
}
