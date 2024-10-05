package config

import (
	"fmt"

	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddr string `env:"SERVER_ADDRESS" env-required:"true"`
}

func LoadConfig() (*Config, error) {
	godotenv.Load() //don't handle errors because we can upload via docker

	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("configuration reading error: %w", err)
	}

	return cfg, nil
}
