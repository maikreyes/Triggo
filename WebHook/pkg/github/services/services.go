package services

import (
	"triggo/pkg/config"
	"triggo/pkg/ports"
)

type Services struct {
	Config      *config.Config
	JWTServices ports.JWTServices
}

func NewServices(config *config.Config, jwtS ports.JWTServices) *Services {
	return &Services{
		Config:      config,
		JWTServices: jwtS,
	}
}
