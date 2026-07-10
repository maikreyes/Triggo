package api

import (
	"net/http"
)

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
