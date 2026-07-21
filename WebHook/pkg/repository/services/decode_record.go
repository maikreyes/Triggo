package services

import (
	"encoding/json"
	"fmt"
	"triggo/pkg/repository/model"
)

func (s *Services) DecodeRecord(b []byte) (model.RepositoryWebhook, error) {

	var record model.RepositoryWebhook

	err := json.Unmarshal(b, &record)

	if err != nil {
		return model.RepositoryWebhook{}, fmt.Errorf("Errot to try decode record: %w", err)
	}

	return record, nil
}
