package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/CRASH-Tech/talos-admin/internal/talos-admin/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type DB struct {
	db  *bun.DB
	ctx context.Context
}

func New(cfg config.СonfigImpl) DB {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.DB_USER,
		cfg.DB_PASS,
		cfg.DB_HOST,
		cfg.DB_PORT,
		cfg.DB_NAME,
	)

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	result := DB{}
	result.ctx = context.TODO()
	result.db = bun.NewDB(sqldb, pgdialect.New())

	return result
}
