package handler

import (
	"io"
	"net/http"

	"github.com/google/uuid"
)

func (h *Handler) RepositoryHandler(w http.ResponseWriter, r *http.Request) {

	b, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Error try to read request body", http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	record, err := h.Services.DecodeRecord(b)

	if err != nil {
		http.Error(w, "Error to try decode record", http.StatusInternalServerError)
		return
	}

	record.Id = uuid.New()

	err = h.Services.CreateRecord(record)

	if err != nil {
		http.Error(w, "Error to try create record", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("record created successfully."))

}
