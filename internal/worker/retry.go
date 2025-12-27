package worker

import (
	"fmt"
	"log"
	"mytipster/internal/fixtures/service"
	m "mytipster/models/fixture"
	odds_models "mytipster/models/odds"
	prediction_models "mytipster/models/prediction"
	"time"
)

func Retry[T any](
	max int,
	delay time.Duration,
	fn func() (T, error),
) (T, error) {

	var lastErr error
	var zero T

	for i := 1; i <= max; i++ {
		result, err := fn()
		if err == nil {
			return result, nil
		}

		lastErr = err
		log.Printf("[RETRY %d/%d] %v", i, max, err)

		if i < max {
			// สูตร: รอเพิ่มขึ้นตามจำนวนรอบ เช่น รอบแรก 2วิ, รอบสอง 4วิ...
			// หรือจะใช้ delay คงที่ตามที่ระบุก็ได้
			wait := delay * time.Duration(i)
			time.Sleep(wait)
		}
	}

	return zero, lastErr
}
func QueryFixtureIdRetry(id string) (*m.Response, error) {
	// --- [ปรับจุดที่ 3] เพิ่มเป็น 5 ครั้ง และรอนานขึ้นรอบละ 3 วินาที ---
	return Retry[*m.Response](5, 3000*time.Millisecond, func() (*m.Response, error) {
		resp, err := service.QueryFixtureId(id)

		if err != nil {
			return nil, err
		}
		// ถ้า API คืนค่าว่าง หรือหาไม่เจอ ให้คืน error เพื่อให้เกิดการ Retry
		if resp == nil || resp.Fixture.ID == 0 {
			return nil, fmt.Errorf("fixture %s data not ready yet", id)
		}

		return resp, nil
	})
}
func QueryFixtureOddsSafe(id string) (map[int][]odds_models.Bet, error) {
	odds, err := service.QueryFixtureOdds(id)
	if err != nil {
		return nil, nil
	}
	return odds, nil
}
func QueryPredictionSafe(id string) (*prediction_models.PredictionResponse, error) {
	pred, err := service.QueryPrediction(id)
	if err != nil {
		// no prediction = normal
		return nil, nil
	}
	return pred, nil
}
