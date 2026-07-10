package ports

import (
	"triggo/pkg/repository/model"
)

type RepositoryServices interface {
	CreateRecord(model.RepositoryWebhook) error
	SearchRecord(installationId int64, repository string) (model.RepositoryWebhook, error)
	DecodeRecord(Body []byte) (model.RepositoryWebhook, error)
}
