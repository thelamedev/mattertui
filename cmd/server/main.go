package main

import (
	"log/slog"
	"os"

	"github.com/thelamedev/mattertui/internal/config"
	"github.com/thelamedev/mattertui/internal/server"
	"github.com/thelamedev/mattertui/internal/server/database"
)

func main() {
	slogOpts := slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slogOpts)))

	cfg, err := config.LoadConfig("..", ".")
	if err != nil {
		slog.Error("failed to load config", "error", err.Error())
		return
	}
	slog.Info("config loaded")

	err = database.InitDatabase(cfg.Database)
	if err != nil {
		slog.Error("failed to init database", "error", err.Error())
		return
	}

	pool := database.GetPool()
	slog.Info("database initialized", "host", pool.Config().ConnConfig.Host, "port", pool.Config().ConnConfig.Port)

	srv := server.NewServer(cfg)

	errCh := srv.Run()
	if !cfg.Server.Debug {
		slog.Info("server started", "bind_addr", cfg.Server.BindAddr)
	}
	err = <-errCh

	if err != nil {
		slog.Error("failed to run server", "error", err.Error())
		return
	}
}
