package helper

import "cloud.google.com/go/bigquery"

// Helper: แปลง *int เป็น NullInt64
func ToNullInt64(v *int) bigquery.NullInt64 {
	if v == nil {
		return bigquery.NullInt64{Valid: false}
	}
	return bigquery.NullInt64{Int64: int64(*v), Valid: true}
}

// Helper: แปลง *string เป็น NullString
func ToNullString(v *string) bigquery.NullString {
	if v == nil || *v == "" {
		return bigquery.NullString{Valid: false}
	}
	return bigquery.NullString{StringVal: *v, Valid: true}
}

// Helper: แปลง *bool เป็น NullBool
func ToNullBool(v *bool) bigquery.NullBool {
	if v == nil {
		return bigquery.NullBool{Valid: false}
	}
	return bigquery.NullBool{Bool: *v, Valid: true}
}
