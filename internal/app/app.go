package app

import (
	"log/slog"
	grpcapp "sso-learn-project-grpc/internal/app/grpc"
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
	// TODO: Initialize the storage

	// TODO: Initialize auth service

	// TODO: Initialize the gRPC server
	grpcApp := grpcapp.New(log, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
