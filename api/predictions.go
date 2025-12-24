package api

import prediction_models "mytipster/models/prediction"

func Predictions(method, url string) (*prediction_models.PredictionsRoot, error) {
	return ApiFootball[prediction_models.PredictionsRoot](method, url)
}
