package api

import (
	"net/http"
	"triggo/pkg/config"
	"triggo/pkg/middleware"
	"triggo/pkg/repository/handler"
	"triggo/pkg/repository/services"
)

var setupHandler *handler.Handler
var middle *middleware.Middleware

func init() {

	ctg := config.NewConfig()

	RepositoryServices := services.NewServices(ctg)
	middle = middleware.NewMiddleware(ctg)

	_, err := RepositoryServices.ConnectDatabase()

	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	setupHandler = handler.NewHandler(RepositoryServices)

}

func SetupHandler(w http.ResponseWriter, r *http.Request) {

	isOptions := middle.Cors(w, r)

	if isOptions {
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	setupHandler.RepositoryHandler(w, r)
}
