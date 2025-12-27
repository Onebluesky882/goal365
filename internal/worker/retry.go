package worker

import (
	"fmt"
	"log"
	"mytipster/internal/fixtures/service"
	m "mytipster/models/fixture"
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
	return Retry[*m.Response](5, 3000*time.Millisecond, func() (*m.Response, error) {
		resp, err := service.QueryFixtureId(id)

		// 1. ถ้า service คืน error (ซึ่งตอนนี้เราดัก len=0 ไว้แล้ว)
		if err != nil {
			return nil, err
		}

		// 2. เช็คกันเหนียว ถ้า resp หลุดมาเป็น nil
		if resp == nil {
			return nil, fmt.Errorf("fixture %s: response is nil", id)
		}

		// 3. เช็คค่า ID (ถ้าโครงสร้างข้างในเป็น pointer ต้องระวัง)
		if resp.Fixture.ID == 0 {
			return nil, fmt.Errorf("fixture %s: data incomplete", id)
		}

		return resp, nil
	})
}
func QueryPredictionRetry(id string) (*prediction_models.PredictionResponse, error) {
	// รอ 1 วิ ก่อนเรียกครั้งแรก (Initial delay)
	time.Sleep(1 * time.Second)

	// ลอง 3 ครั้ง ห่างกันครั้งละ 2 วินาที
	return Retry[*prediction_models.PredictionResponse](3, 2000*time.Millisecond, func() (*prediction_models.PredictionResponse, error) {
		pred, err := service.QueryPrediction(id)

		if err != nil {
			return nil, err
		}

		// ❌ ถ้า Prediction เป็น nil หรือข้อมูลข้างในไม่มี
		if pred == nil || pred.Predictions.Advice == "" {
			return nil, fmt.Errorf("prediction for %s is still empty", id)
		}

		return pred, nil
	})
}
