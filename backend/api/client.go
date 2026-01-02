package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func ApiFootball[T any](method string, url string) (*T, error) {
	_ = godotenv.Load()

	apiKey := os.Getenv("API_KEY")

	if apiKey == "" {
		return nil, errors.New("missing API_KEY")
	}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("x-apisports-key", apiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var data T

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
