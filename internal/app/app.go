package app

import (
	"log/slog"
	grpcapp "sso-learn-project-grpc/internal/app/grpc"
	"sso-learn-project-grpc/internal/services/auth"
	"sso-learn-project-grpc/internal/storage/sqlite"
	"time"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	// Initialize the storage
	storage, err := sqlite.New(storagePath)
	if err != nil {
		panic(err)
	}

	authService := auth.New(log, storage, storage, storage, tokenTTL)

	// TODO: Initialize auth service

	// TODO: Initialize the gRPC server
	grpcApp := grpcapp.New(log, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
