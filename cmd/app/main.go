package main

import (
	"log/slog"
	"os"

	"dang.z.v.task/internal/config"
	"dang.z.v.task/internal/domain"
	"dang.z.v.task/internal/storage/postgresql"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	dsn := cfg.DB.DSN()

	storage, err := postgresql.New(dsn)
	if err != nil {
		log.Error("failed to init storage:", slog.String("errormsg", err.Error()))
		return
	}
	_ = storage 
	
	log.Info("Инициализирована бд, написан мигратор, добавлены миграции")
	log.Info("Данг Зуй Ву написал сервис")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
