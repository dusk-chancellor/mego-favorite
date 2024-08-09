package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DBUser 	   string `env:"DB_USER" required:"true"`
	DBPassword string `env:"DB_PASSWORD" required:"true"`
	DBHost 	   string `env:"DB_HOST" required:"true"`
	DBPort 	   string `env:"DB_PORT" required:"true"`
	DBName 	   string `env:"DB_NAME" required:"true"`

	GRPCPort   string `env:"FAVORITE_SERVICE_SERVER_PORT" required:"true"`
}

func LoadConfig() *Config {
	path := "./.env"
	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		log.Fatalf("failed to read config: %v", err)
	}
	return &cfg
}
