package services

import (
	"triggo/pkg/repository/model"
)

func (s *Services) CreateRecord(record model.RepositoryWebhook) error {

	result := s.DB.Create(record)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
