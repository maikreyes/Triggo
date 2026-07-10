package services

import (
	"triggo/pkg/repository/model"
)

func (s *Services) SearchRecord(installationid int64, repository string) (model.RepositoryWebhook, error) {

	var record model.RepositoryWebhook

	search := s.DB.Where("installation_id = ? AND repository = ?", installationid, repository).First(&record)

	if search.Error != nil {
		return model.RepositoryWebhook{}, search.Error
	}

	return record, nil

}
