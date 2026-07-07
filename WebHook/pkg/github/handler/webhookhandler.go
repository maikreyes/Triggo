package handler

import (
	"io"
	"log"
	"net/http"
)

func (h *Handler) WebhookHandler(w http.ResponseWriter, r *http.Request) {

	//Headers
	event := r.Header.Get("X-GitHub-Event")
	signature := r.Header.Get("X-Hub-Signature-256")

	bodyBites, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Error try to read request body", http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	err = h.Services.ValidatedHash(signature, bodyBites)

	if err != nil {
		http.Error(w, "Error to validated", http.StatusUnauthorized)
		log.Println(err)
		return
	}

	h.Services.DecodeMessage(event, string(bodyBites))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Webhook received and successfully validated."))

}
