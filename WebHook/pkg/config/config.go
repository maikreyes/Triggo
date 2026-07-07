package config

import "os"

type Config struct {
	Secret string
}

func NewConfig() *Config {
	return &Config{
		Secret: os.Getenv("GITHUB_WEBHOOK_SECRET"),
	}
}
