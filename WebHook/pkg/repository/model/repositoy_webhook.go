package model

import "github.com/google/uuid"

type RepositoryWebhook struct {
	Id             uuid.UUID `json:"id"`
	InstallationId int64     `json:"installation_id"`
	Repository     string    `json:"repository"`
	DiscordUrl     string    `json:"discord_url"`
}
