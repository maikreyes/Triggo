package api

import (
	"log"
	"net/http"
	"triggo/pkg/config"
	"triggo/pkg/github/handler"
	"triggo/pkg/github/services"

	"github.com/joho/godotenv"
)

func Webhook(w http.ResponseWriter, r *http.Request) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := config.NewConfig()

	//GitHub configuration
	GithubServices := services.NewSercies(config)
	GithubHandler := handler.Newhandler(GithubServices)

	GithubHandler.WebhookHandler(w, r)

}
