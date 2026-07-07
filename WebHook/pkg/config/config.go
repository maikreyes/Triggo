package config

import "os"

type Config struct {
	Secret string
}

func NewConfig() *Config {
	return &Config{
		Secret: os.Getenv("GITHHUB_WEBHOOK_SECRET"),
	}
}
