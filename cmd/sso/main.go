package main

import (
	"fmt"
	"sso-learn-project-grpc/internal/config"
)

func main() {
	// TODO: init app config.
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// TODO: init app logger.

	// TODO: init app (app).

	// TODO: run gRPC app.
}
