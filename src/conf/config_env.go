package conf

import (
	"github.com/caarlos0/env/v11"
)

type AppConfig struct {
	Database struct {
		PostgresHost     string `env:"POSTGRES_HOST"`
		PostgresPort     string `env:"POSTGRES_PORT"`
		PostgresUser     string `env:"POSTGRES_USER"`
		PostgresPassword string `env:"POSTGRES_PASSWORD"`
		PostgresDB       string `env:"POSTGRES_DB"`
	}
}

var config AppConfig

func SetEnv() error {
	return env.Parse(&config)
}

func LoadEnv() AppConfig {
	return config
}
