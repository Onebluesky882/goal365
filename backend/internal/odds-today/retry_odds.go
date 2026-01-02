package oddstoday

import (
	"log"
	"mytipster/internal/fixtures"
	"mytipster/lib"
	"strconv"
	"time"
)

// RetryFixtureOdds - ใช้ RetryWithBackoff แทน manual retry
func RetryFixtureOdds(fixtureID int) error {
	idStr := strconv.Itoa(fixtureID)

	err := lib.RetryWithBackoff(func() error {
		log.Printf("⚠️  [Retry] fixture %d attempting...", fixtureID)

		time.Sleep(500 * time.Millisecond)

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

// RetryLater - เรียกใช้ในกรณีที่ต้องการ retry ภายหลัง
func RetryLater(fixtureID int) {
	go func() {
		time.Sleep(5 * time.Second) // รอก่อน retry
		if err := RetryFixtureOdds(fixtureID); err != nil {
			log.Printf("⚠️  RetryLater failed for fixture %d", fixtureID)
		}
	}()
}

// เก็บ function เดิมไว้สำหรับ backward compatibility
func retryFixtureOdds(fixtureID int) {
	idStr := strconv.Itoa(fixtureID)

	for attempt := 1; attempt <= 3; attempt++ {
		time.Sleep(1 * time.Second)

		log.Printf("[Retry] fixture %d attempt %d/3", fixtureID, attempt)

		_, err := fixtures.QueryMyTipsOdds(idStr)
		if err == nil {
			log.Printf("[Retry] fixture %d success", fixtureID)
			return
		}

		time.Sleep(2 * time.Second)
	}

	log.Printf("[Retry] fixture %d failed after 3 attempts", fixtureID)
}
