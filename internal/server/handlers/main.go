package handlers

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/thelamedev/mattertui/internal/config"
	"github.com/thelamedev/mattertui/internal/server/store"
)

var queries *store.Queries
var cfg *config.Config

func InitStore(pool *pgxpool.Pool) {
	queries = store.New(pool)
}

func InitConfig(c *config.Config) {
	cfg = c
}
