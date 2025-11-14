package main

import (
	"log/slog"
	"net/http"
	"os"

	"dang.z.v.task/internal/config"
	"dang.z.v.task/internal/handlers/prreviewer"
	"dang.z.v.task/internal/handlers/pullrequest"
	"dang.z.v.task/internal/handlers/team"
	"dang.z.v.task/internal/handlers/user"
	"dang.z.v.task/internal/storage/postgresql"
	"github.com/go-chi/chi/v5"
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

	router := chi.NewRouter()

	router.Mount("/users", user.NewHandler())
	router.Mount("/team", team.NewHandler())
	router.Mount("/pullRequest", pullrequest.NewHandler())
	router.Mount("/prReviewer", prreviewer.NewHandler())

	server := &http.Server{
		Addr:         cfg.HTTPServer.ServerAddr(),
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	log.Info("starting server")
	if err := server.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}

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
