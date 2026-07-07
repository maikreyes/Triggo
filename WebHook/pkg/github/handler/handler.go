package handler

import (
	"triggo/pkg/ports"
)

type Handler struct {
	Services ports.GithubServices
}

func Newhandler(s ports.GithubServices) *Handler {
	return &Handler{
		Services: s,
	}
}
