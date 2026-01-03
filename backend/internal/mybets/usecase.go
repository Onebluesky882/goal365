package mybets

import (
	"fmt"
	m "mytipster/models/analytic"
	"time"

	"github.com/google/uuid"
)

func FindId(id string, items []m.MyAnalytics) (*m.MyAnalytics, error) {

	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err // หรือ handle error
	}

	for _, i := range items {
		if i.ID == uid {
			return &i, nil
		}
	}
	return nil, fmt.Errorf("id %s not found", id)
}

func FilterPredictionByDate(date string, items []m.MyAnalytics) ([]m.MyAnalytics, error) {
	filterDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, err
	}
	var result []m.MyAnalytics
	for _, item := range items {
		itemDate, err := time.Parse("2006-01-02", item.Date)
		if err != nil {
			continue // skip invalid date
		}
		if itemDate.Equal(filterDate) {
			result = append(result, item)
		}
	}
	return result, nil

}
