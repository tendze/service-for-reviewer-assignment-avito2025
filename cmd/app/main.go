package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"dang.z.v.task/internal/config"
	"dang.z.v.task/internal/handlers/prreviewer"
	"dang.z.v.task/internal/handlers/pullrequest"
	"dang.z.v.task/internal/handlers/team"
	"dang.z.v.task/internal/handlers/user"
	prservice "dang.z.v.task/internal/service/pullrequest"
	teamservice "dang.z.v.task/internal/service/team"
	userservice "dang.z.v.task/internal/service/user"
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

	router := chi.NewRouter()

	userService := userservice.NewUserService(storage, log)
	teamService := teamservice.NewTeamService(storage, log)
	prService := prservice.NewPullRequestService(storage, log)

	router.Mount("/users", user.NewHandler(userService))
	router.Mount("/team", team.NewHandler(teamService))
	router.Mount("/pullRequest", pullrequest.NewHandler(prService))
	router.Mount("/prReviewer", prreviewer.NewHandler())

	server := &http.Server{
		Addr:         cfg.HTTPServer.ServerAddr(),
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		log.Info("starting server")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error("failed to start server", slog.String("err", err.Error()))
		}
	}()

	<-stop
	log.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), cfg.HTTPServer.ShuttingDownTimeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Error("graceful shutdown failed", slog.String("err", err.Error()))
	} else {
		log.Info("server gracefully stopped")
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
