package database

import (
	"context"
	"log"
	analytic_module "mytipster/models/analytic"

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
		(*analytic_module.MyAnalytics)(nil),
		// (*bets_models.Bets)(nil),
		// (*tipsdaliy_models.TipsDaily)(nil),
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
