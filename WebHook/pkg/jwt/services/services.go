package services

import "triggo/pkg/config"

type Services struct {
	Config *config.Config
}

func NewServices(ctg *config.Config) *Services {
	return &Services{
		Config: ctg,
	}
}
