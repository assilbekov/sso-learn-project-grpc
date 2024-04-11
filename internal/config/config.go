package config

import (
	"flag"
	"os"
	"time"
)

type Config struct {
	Env         string        `yaml:"env" env-default:"local"`
	StoragePath string        `yaml:"storage_path" env-required:"true"`
	TokenTTL    time.Duration `yaml:"token_ttl" env-default:"1h"`
	GRPC        GRPCConfig    `yaml:"grpc"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

// fetchConfigPath fetches config path from command line arguments or environment variables.
// Priority: flag > env > default.
// Default value is empty string.
func fetchConfigPath() string {
	var res string

	// --config="/path/to/config.yaml"
	flag.StringVar(&res, "config", "", "config path")
	flag.Parse()

	if res != "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
