package api

import (
	"net/http"
	"triggo/pkg/config"
	DServices "triggo/pkg/discord/services"
	"triggo/pkg/github/handler"
	GServices "triggo/pkg/github/services"
	"triggo/pkg/middleware"
	RepoHandler "triggo/pkg/repository/handler"
	RServices "triggo/pkg/repository/services"
)

var webhookHandler *handler.Handler
var setupHandler *RepoHandler.Handler
var middle *middleware.Middleware

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
	setupHandler = RepoHandler.NewHandler(RepositoryServices)
	middle = middleware.NewMiddleware(cfg)
}

func Webhook(w http.ResponseWriter, r *http.Request) {

	webhookHandler.WebhookHandler(w, r)

}
