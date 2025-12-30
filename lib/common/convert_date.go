package common

import "time"

func ToBangkokDate(dateStr string) (string, error) {
	// parse ISO8601
	t, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		return "", err
	}

	// load Bangkok location
	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		return "", err
	}

	// convert timezone
	tBKK := t.In(loc)

	// format yyyy-mm-dd
	return tBKK.Format("2006-01-02"), nil
}
