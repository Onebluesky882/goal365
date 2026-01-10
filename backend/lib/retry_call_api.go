package lib

import (
	"fmt"
	"log"
	"time"
)

// RetryConfig กำหนดค่า retry
type RetryConfig struct {
	MaxRetries   int
	InitialDelay time.Duration
	MaxDelay     time.Duration
	Multiplier   float64
}

// DefaultRetryConfig ค่า default สำหรับ retry
var DefaultRetryConfig = RetryConfig{
	MaxRetries:   3,
	InitialDelay: 500 * time.Millisecond,
	MaxDelay:     10 * time.Second,
	Multiplier:   2.0,
}

// Retry ลอง execute function หลายครั้งจนสำเร็จ

func Retry(fn func() error, config RetryConfig) error {
	var err error
	delay := config.InitialDelay

	for attempt := 0; attempt < config.MaxRetries; attempt++ {
		err = fn()
		if err == nil {
			return nil
		}

		// ถ้ายังไม่ถึง max retries ให้รอก่อน retry
		if attempt < config.MaxRetries-1 {
			log.Printf("⚠️  Retry attempt %d/%d after %v: %v",
				attempt+1, config.MaxRetries, delay, err)
			time.Sleep(delay)

			// Exponential backoff
			delay = time.Duration(float64(delay) * config.Multiplier)
			if delay > config.MaxDelay {
				delay = config.MaxDelay
			}
		}
	}

	return fmt.Errorf("failed after %d retries: %w", config.MaxRetries, err)
}

// RetryWithBackoff ใช้ exponential backoff (shorthand)
func RetryWithBackoff(fn func() error, maxRetries int, initialDelay time.Duration) error {
	return Retry(fn, RetryConfig{
		MaxRetries:   maxRetries,
		InitialDelay: initialDelay,
		MaxDelay:     30 * time.Second,
		Multiplier:   2.0,
	})
}

// SimpleRetry retry แบบง่าย ไม่มี backoff
func SimpleRetry(fn func() error, maxRetries int, delay time.Duration) error {
	var err error

	for attempt := 0; attempt < maxRetries; attempt++ {
		err = fn()
		if err == nil {
			return nil
		}

		if attempt < maxRetries-1 {
			time.Sleep(delay)
		}
	}

	return fmt.Errorf("failed after %d retries: %w", maxRetries, err)
}
