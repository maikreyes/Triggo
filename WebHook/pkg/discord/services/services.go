package services

import (
	"triggo/pkg/config"
	"triggo/pkg/ports"
)

type Services struct {
	Config     *config.Config
	Repository ports.RepositoryServices
}

func NewServices(config *config.Config, repo ports.RepositoryServices) *Services {
	return &Services{
		Config:     config,
		Repository: repo,
	}
}
