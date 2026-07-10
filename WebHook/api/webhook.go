package api

import (
	"net/http"
	"triggo/pkg/config"
	DServices "triggo/pkg/discord/services"
	"triggo/pkg/github/handler"
	GServices "triggo/pkg/github/services"
	RServices "triggo/pkg/repository/services"
)

var webhookHandler *handler.Handler

func init() {
	cfg := config.NewConfig()

	//Repository configuration
	RepositoryServices := RServices.NewServices(cfg)
	_, err := RepositoryServices.ConnectDatabase()

	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	//GitHub configuration
	GithubServices := GServices.NewServices(cfg)

	//Discord configuration
	DiscordServices := DServices.NewServices(cfg, RepositoryServices)

	webhookHandler = handler.Newhandler(GithubServices, DiscordServices)
}

func Webhook(w http.ResponseWriter, r *http.Request) {

	webhookHandler.WebhookHandler(w, r)

}
