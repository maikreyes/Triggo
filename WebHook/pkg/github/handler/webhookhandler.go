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

	err = h.GServices.ValidatedHash(signature, bodyBites)

	if err != nil {
		http.Error(w, "Error to validated", http.StatusUnauthorized)
		log.Println(err)
		return
	}

	message := h.GServices.DecodeMessage(event, bodyBites)

	embed := h.DServices.CreateEmbed(event, message)
	payload := h.DServices.CreateDiscordPayload(embed)

	err = h.DServices.SendPayload(payload)

	if err != nil {
		log.Println("Discord Service Error:", err)
		http.Error(w, "Failed to deliver message to Discord", http.StatusBadGateway)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Webhook received and successfully validated."))

}
