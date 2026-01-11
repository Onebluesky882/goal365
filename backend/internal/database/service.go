package database

import (
	"context"
	"fmt"
	"log"
	"mytipster/models"
	"time"

	"github.com/uptrace/bun"
)

var (
	// global DB instance
	DBConn *bun.DB
)

func InitDB() {
	var err error
	DBConn, err = NewDB()
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}
}

func WithContext(ctx context.Context) *bun.DB {
	if DBConn == nil {
		log.Fatal("DB not initialized. Call service.InitDB() first.")
	}
	return DBConn
}

func CreateTables(ctx context.Context, db *bun.DB) error {
	models := []interface{}{
		// (*models.Player)(nil),
		// (*models.Transaction)(nil),
		(*models.BetTransaction)(nil),
		// (*models.MyAnalytics)(nil),
		// (*models.NaWinTatips)(nil),
		// (*models.Bets)(nil),
		// (*models.TipsDaily)(nil),
	}

	for _, m := range models {
		if _, err := db.NewCreateTable().
			Model(m).
			IfNotExists().
			Exec(ctx); err != nil {
			return err
		}
	}

	return nil
}

func EnsurePartition(ctx context.Context, db *bun.DB, date time.Time) error {
	if date.IsZero() {
		return fmt.Errorf("match_date is zero (0001-01-01), cannot create partition")
	}

	from := date.Format("2006-01-02")
	to := date.AddDate(0, 0, 1).Format("2006-01-02")
	table := fmt.Sprintf("sportsbooks_%s", date.Format("2006_01_02"))

	_, err := db.ExecContext(ctx, fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s
		PARTITION OF sportsbooks
		FOR VALUES FROM ('%s') TO ('%s');
	`, table, from, to))

	return err
}
