package mybets_test

import (
	"mytipster/internal/mybets"
	m "mytipster/models/mytips"
	"testing"
)

func TestFilterTipsDailyByDate(t *testing.T) {
	date := "2026-01-02"

	items := []m.TipsDaily{
		{
			FixtureID: 1,
			Date:      "1767369600", // 2026-01-02 UTC
		},
		{
			FixtureID: 2,
			Date:      "1767283200", // 2026-01-01 UTC
		},
	}

	result, err := mybets.FilterPredictionByDate(date, items)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(result) != 1 {
		t.Fatalf("expected 1 item, got %d", len(result))
	}

	if result[0].FixtureID != 1 {
		t.Errorf("expected FixtureID 1, got %d", result[0].FixtureID)
	}
}
