package common

import "time"

func Timestamp(ts int64) string {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	return time.Unix(ts, 0).In(loc).Format("15:04:05")
}
func TimestampDate(ts int64) string {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	return time.Unix(ts, 0).In(loc).Format("2006-01-02")
}