package services

import (
	"triggo/pkg/config"

	"gorm.io/gorm"
)

type Services struct {
	Config *config.Config
	DB     *gorm.DB
}

func NewServices(cfg *config.Config) *Services {
	return &Services{
		Config: cfg,
	}
}
