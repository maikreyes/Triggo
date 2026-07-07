package config

import "os"

type Config struct {
	Secret string
	DUrl   string
}

func NewConfig() *Config {
	return &Config{
		Secret: os.Getenv("GITHUB_WEBHOOK_SECRET"),
		DUrl:   os.Getenv("DISCORD_WEBHOOK_URL"),
	}
}
