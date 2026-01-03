package oddstoday

import (
	"log"
	"mytipster/internal/fixtures"
	"mytipster/lib"
	"strconv"
	"time"
)

// RetryLater - เรียกใช้ในกรณีที่ต้องการ retry ภายหลัง
func RetryLater(fixtureID int) {
	go func() {
		if err := retryFixtureOdds(fixtureID); err != nil {
			log.Printf("⚠️  RetryLater failed for fixture %d", fixtureID)
		}
	}()
}

// RetryFixtureOdds - ใช้ RetryWithBackoff แทน manual retry
func retryFixtureOdds(fixtureID int) error {
	idStr := strconv.Itoa(fixtureID)
	time.Sleep(500 * time.Millisecond)
	err := lib.RetryWithBackoff(func() error {
		log.Printf("⚠️  [Retry] fixture %d attempting...", fixtureID)

		_, err := fixtures.QueryFixtureOdds(idStr)
		if err != nil {
			return err
		}

		log.Printf("✅ [Retry] fixture %d success", fixtureID)
		return nil
	}, 3, 1*time.Second)

	if err != nil {
		log.Printf("❌ [Retry] fixture %d failed after 3 attempts: %v", fixtureID, err)
		return err
	}

	return nil
}
