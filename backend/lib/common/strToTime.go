package common

import (
	"time"
)

func StringToTime(rfc3339 string) time.Time {
	t, err := time.Parse(time.RFC3339, rfc3339)
	if err != nil {
		return time.Time{}
	}

	// ตัดเวลา เหลือแค่ DATE (00:00:00 UTC)
	return time.Date(
		t.Year(), t.Month(), t.Day(),
		0, 0, 0, 0,
		time.UTC,
	)
}
