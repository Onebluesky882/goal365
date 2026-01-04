package database

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

var DB *bun.DB

func NewDB() (*bun.DB, error) {
	_ = godotenv.Load()
	dsn := os.Getenv("DATABASE_URL")

	if strings.Contains(dsn, "railway.internal") && os.Getenv("ENV") != "production" {
		return nil, fmt.Errorf("railway internal db used outside railway")
	}

	sqldb := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithDSN(dsn),
	))
	if os.Getenv("DATABASE_URL") == "" {
		fmt.Println("no db_url")
	}

	db := bun.NewDB(sqldb, pgdialect.New())

	// เปิด debug SQL (dev only)
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
	))

	return db, nil
}
