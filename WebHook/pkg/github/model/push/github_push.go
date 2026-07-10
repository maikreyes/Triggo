package push

import (
	"triggo/pkg/github/model/installation"
	"triggo/pkg/github/model/pusher"
	"triggo/pkg/github/model/repository"
)

type GithubPush struct {
	Ref          string                    `json:"ref"`
	Repository   repository.Repository     `json:"repository"`
	Pusher       pusher.Pusher             `json:"pusher"`
	Installation installation.Installation `json:"installation"`
}
