package mybets_models

import "github.com/google/uuid"

type InsertPickedRequest struct {
	AnalyticsID uuid.UUID   `json:"analytics_id"`
	Items       []BetPickIn `json:"items"`
}

type BetPickIn struct {
	Picked string `json:"picked"`
	Team   string `json:"team"`
	Odds   string `json:"odds"`
	Stake  string `json:"stake"`
}
