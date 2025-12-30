package service

import (
	"context"
	"log"
	"mytipster/internal/db"

	"github.com/uptrace/bun"
)

var (
	// global DB instance
	DBConn *bun.DB
)

func InitDB() {
	var err error
	DBConn, err = db.NewDB()
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
