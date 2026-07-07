package handler

import (
	"triggo/pkg/ports"
)

type Handler struct {
	GServices ports.GithubServices
	DServices ports.DiscordServices
}

func Newhandler(gs ports.GithubServices, ds ports.DiscordServices) *Handler {
	return &Handler{
		GServices: gs,
		DServices: ds,
	}
}
