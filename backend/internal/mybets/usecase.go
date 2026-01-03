package mybets

import (
	"fmt"
	"mytipster/lib/common"
	m "mytipster/models/mytips"
	"strconv"

	"github.com/google/uuid"
)

func FindId(id string, items []m.TipsDaily) (*m.TipsDaily, error) {

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

func FilterPredictionByDate(date string, items []m.TipsDaily) ([]m.TipsDaily, error) {

	var result []m.TipsDaily
	for _, item := range items {
		ts, err := strconv.ParseInt(item.Date, 10, 64)

		if err != nil {
			return nil, err
		}
		format := common.TimestampUTCDate(ts)
		if format != date {
			continue
		}
		result = append(result, item)
	}
	return result, nil

}
