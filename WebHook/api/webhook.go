package api

import (
	"net/http"
	"triggo/pkg/config"
	DServices "triggo/pkg/discord/services"
	"triggo/pkg/github/handler"
	GServices "triggo/pkg/github/services"
)

func Webhook(w http.ResponseWriter, r *http.Request) {

	config := config.NewConfig()

	//GitHub configuration
	GithubServices := GServices.NewServices(config)

	//Discord Cinfiguration
	DiscordServices := DServices.NewServices(config)

	Handler := handler.Newhandler(GithubServices, DiscordServices)

	Handler.WebhookHandler(w, r)

}
