package api

import (
	"net/http"
	"triggo/pkg/config"
	"triggo/pkg/github/handler"
	"triggo/pkg/github/services"
)

func Webhook(w http.ResponseWriter, r *http.Request) {

	config := config.NewConfig()

	//GitHub configuration
	GithubServices := services.NewSercies(config)
	GithubHandler := handler.Newhandler(GithubServices)

	GithubHandler.WebhookHandler(w, r)

}
