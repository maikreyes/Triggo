package ports

import (
	"triggo/pkg/github/model/installation"
	messainfromation "triggo/pkg/github/model/messa_infromation"
	"triggo/pkg/github/model/repository"
)

type GithubServices interface {
	ValidatedHash(signature string, payload []byte) error
	DecodeMessage(event string, body []byte) (messainfromation.MessaInformation, string)
	RequestAccessToken(installationId string) (installation.InstallationAccesToken, error)
	RequestInstallationRepositories(token string) (repository.InstallationRepositories, error)
}
