package api

import (
	"net/http"
	"triggo/internal/config"
	"triggo/internal/github/handler"
	"triggo/internal/github/services"
)

func Webhook(w http.ResponseWriter, r *http.Request) {

	config := config.NewConfig()

	//GitHub configuration
	GithubServices := services.NewSercies(config)
	GithubHandler := handler.Newhandler(GithubServices)

	GithubHandler.WebhookHandler(w, r)

}
