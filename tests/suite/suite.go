package suite

import (
	ssov1 "github.com/assilbekov/grpc-app-course-protoes/gen/go/sso"
	"sso-learn-project-grpc/internal/config"
	"testing"
)

type Suite struct {
	*testing.T
	Cfg        *config.Config
	AuthClient ssov1.AuthClient
}
