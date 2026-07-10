package config

import "os"

type Config struct {
	Secret   string
	DSN      string
	FrontUrl string
}

func NewConfig() *Config {
	return &Config{
		Secret:   os.Getenv("GITHUB_WEBHOOK_SECRET"),
		DSN:      os.Getenv("CONNECTION_STRING"),
		FrontUrl: os.Getenv("FRONT_URL"),
	}
}
