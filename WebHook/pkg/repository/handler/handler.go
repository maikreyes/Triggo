package handler

import (
	"triggo/pkg/ports"
)

type Handler struct {
	Services ports.RepositoryServices
}

func NewHandler(s ports.RepositoryServices) *Handler {
	return &Handler{
		Services: s,
	}
}
