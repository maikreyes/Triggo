package services

import "triggo/internal/config"

type Services struct {
	Config *config.Config
}

func NewSercies() *Services {
	return &Services{}
}
