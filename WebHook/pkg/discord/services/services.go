package services

import "triggo/pkg/config"

type Services struct {
	Config *config.Config
}

func NewServices(config *config.Config) *Services {
	return &Services{
		Config: config,
	}
}
