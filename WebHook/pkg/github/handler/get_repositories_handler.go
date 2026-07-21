package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (h *Handler) GetRepositoriesHandler(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()

	installationID := query.Get("installation_id")

	installationAccesToken, err := h.GServices.RequestAccessToken(installationID)

	if err != nil {
		http.Error(w, "Error to try create token", http.StatusInternalServerError)
		return
	}

	repositories, err := h.GServices.RequestInstallationRepositories(installationAccesToken.Token)

	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Error to try respositories", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(&repositories.Repositories)

	log.Println()

	if err != nil {
		http.Error(w, "error to try encode response", http.StatusInternalServerError)
		return
	}
}
