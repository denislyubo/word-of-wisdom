package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type ServerConfig struct {
	ServerPort      uint          `env:"SERVER_PORT" env-default:"8080"`
	ServerKeepAlive time.Duration `env:"SERVER_KEEP_ALIVE" env-default:"15s"`
	ServerDeadline  time.Duration `env:"SERVER_DEAD_LINE" env-default:"10s"`
	Difficulty      uint8         `env:"DIFFICULTY" env-default:"6"`
}

type ClientConfig struct {
	ServerHost string `env:"SERVER_HOST" env-default:"localhost"`
	ServerPort uint   `env:"SERVER_PORT" env-default:"8080"`
	ClientRps  uint64 `env:"CLIENT_RPS" env-default:"10"`
	Difficulty uint8  `env:"DIFFICULTY" env-default:"6"`
}

func Load[C any](config *C) error {
	err := cleanenv.ReadEnv(config)
	if err != nil {
		return err
	}

	return nil
}
