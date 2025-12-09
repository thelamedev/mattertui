package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/thelamedev/mattertui/internal/config"
)

var pool *pgxpool.Pool

func InitDatabase(config config.DatabaseConfig) error {
	var err error
	poolCfg, err := pgxpool.ParseConfig(config.Url)
	if err != nil {
		return err
	}
	pool, err = pgxpool.NewWithConfig(context.Background(), poolCfg)
	if err != nil {
		return err
	}
	return nil
}

func GetPool() *pgxpool.Pool {
	return pool
}
