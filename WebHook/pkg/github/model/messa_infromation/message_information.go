package messainfromation

import (
	"triggo/pkg/github/model/installation"
	"triggo/pkg/github/model/repository"
)

type MessaInformation struct {
	Installation installation.Installation `json:"installation"`
	Repository   repository.Repository     `json:"repository"`
}
