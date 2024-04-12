package main

import (
	"log/slog"
	"os"
	"sso-learn-project-grpc/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// init app config.
	cfg := config.MustLoad()

	// init app logger.
	log := setupLogger(cfg.Env)
	log.Info(
		"starting application",
		slog.String("env", cfg.Env),
		slog.Int("port", cfg.GRPC.Port))

	log.Debug("Debug message")

	log.Error("Error message")
	log.Warn("Warning message")

	// TODO: init app (app).

	// TODO: run gRPC app.
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
