package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type ServerConfig struct {
	ServerPort      uint          `env:"SERVER_PORT" env-default:"8080"`
	ServerKeepAlive time.Duration `env:"SERVER_KEEP_ALIVE" env-default:"15s"`
	ClientRps       uint64        `env:"CLIENT_RPS" env-default:"10"`
}

type ClientConfig struct {
	ServerHost string `env:"SERVER_HOST" env-default:"localhost"`
	ServerPort uint   `env:"SERVER_PORT" env-default:"8080"`
}

func Load[C any](config C) (*C, error) {
	err := cleanenv.ReadEnv(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
