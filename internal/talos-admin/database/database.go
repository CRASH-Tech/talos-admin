package database

import (
	"database/sql"
	"fmt"

	"github.com/CRASH-Tech/talos-admin/internal/talos-admin/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func New(cfg config.Сonfig) *bun.DB {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.DB_USER,
		cfg.DB_PASS,
		cfg.DB_HOST,
		cfg.DB_PORT,
		cfg.DB_NAME,
	)

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	return bun.NewDB(sqldb, pgdialect.New())
}
