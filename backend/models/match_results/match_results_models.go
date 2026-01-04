package matchresults_models

type UpdateFixtureResultDTO struct {
	FixtureID   int    `json:"fixture_id"`
	MatchFinish string `json:"match_finish"`
	MatchResult string `json:"match_result"`
}
