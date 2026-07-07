package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
	"triggo/pkg/config"
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

	b, err := json.Marshal(payload)

	if err != nil {
		http.Error(w, "Error to conver payload", http.StatusInternalServerError)
		return
	}

	config := config.NewConfig()

	req, err := http.NewRequest(http.MethodPost, config.DUrl, bytes.NewBuffer(b))

	if err != nil {
		http.Error(w, "Error to try create request", http.StatusInternalServerError)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Do(req)

	if err != nil {
		http.Error(w, "Error to try send message to discrod", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		bodyBites, _ := io.ReadAll(resp.Body)
		http.Error(w, "discord webhook returned non-2xx", http.StatusBadGateway)
		w.Write(bodyBites)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Webhook received and successfully validated."))

}
