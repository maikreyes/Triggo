package api

import "net/http"

func GetRepositories(w http.ResponseWriter, r *http.Request) {

	err := initApp()

	if err != nil {
		http.Error(w, "Error interno iniciando la aplicación: "+err.Error(), http.StatusInternalServerError)
		return
	}

	isOptions := middle.Cors(w, r)

	if isOptions {
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	webhookHandler.GetRepositoriesHandler(w, r)

}
