package mybets_models

import "github.com/google/uuid"

type InsertPickedRequest struct {
	AnalyticsID uuid.UUID   `json:"analytics_id"`
	Items       []BetPickIn `json:"items"`
}

type BetPickIn struct {
	Handicap string `json:"handicap,omitempty"`
	Team     string `json:"team,omitempty"`
	Odds     string `json:"odds,omitempty"`
	Stake    string `json:"stake,omitempty"`
	Result   string `json:"result,omitempty"`
	Amount   int    `json:"amount,omitempty"`
	Profit   int    `json:"profit,omitempty"`
	Note     string `json:"note,omitempty"`
}
