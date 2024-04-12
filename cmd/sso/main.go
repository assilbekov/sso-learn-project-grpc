package main

import (
	"log/slog"
	"os"
	"sso-learn-project-grpc/internal/app"
	"sso-learn-project-grpc/internal/config"
	"sso-learn-project-grpc/internal/lib/logger/handlers/slogpretty"
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

	// init app (app).
	application := app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)
	application.GRPCSrv.MustRun()

	// TODO: run gRPC app.
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{Level: slog.LevelDebug},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
