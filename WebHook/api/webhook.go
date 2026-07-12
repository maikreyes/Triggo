package api

import (
	"net/http"
	"triggo/pkg/config"
	DServices "triggo/pkg/discord/services"
	"triggo/pkg/github/handler"
	GServices "triggo/pkg/github/services"
	JServices "triggo/pkg/jwt/services"
	"triggo/pkg/middleware"
	RepoHandler "triggo/pkg/repository/handler"
	RServices "triggo/pkg/repository/services"
)

var webhookHandler *handler.Handler
var setupHandler *RepoHandler.Handler
var middle *middleware.Middleware

func initApp() error {
	cfg := config.NewConfig()

	//Repository configuration
	RepositoryServices := RServices.NewServices(cfg)
	_, err := RepositoryServices.ConnectDatabase()

	if err != nil {
		return err
	}

	//JWT configuration
	JWTServices := JServices.NewServices(cfg)

	//GitHub configuration
	GithubServices := GServices.NewServices(cfg, JWTServices)

	//Discord configuration
	DiscordServices := DServices.NewServices(cfg, RepositoryServices)

	webhookHandler = handler.Newhandler(GithubServices, DiscordServices)
	setupHandler = RepoHandler.NewHandler(RepositoryServices)
	middle = middleware.NewMiddleware(cfg)

	return nil
}

func Webhook(w http.ResponseWriter, r *http.Request) {

	err := initApp()

	if err != nil {
		http.Error(w, "Error interno iniciando la aplicación: "+err.Error(), http.StatusInternalServerError)
		return
	}

	webhookHandler.WebhookHandler(w, r)

}
