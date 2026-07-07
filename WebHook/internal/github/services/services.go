package services

import "triggo/internal/config"

type Services struct {
	Config *config.Config
}

func NewSercies(config *config.Config) *Services {
	return &Services{
		Config: config,
	}
}
