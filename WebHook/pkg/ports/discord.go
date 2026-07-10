package ports

import (
	"triggo/pkg/discord/model/embed"
	"triggo/pkg/discord/model/payload"
	messainfromation "triggo/pkg/github/model/messa_infromation"
)

type DiscordServices interface {
	CreateEmbed(event string, message string) embed.Embed
	CreateDiscordPayload(e embed.Embed) payload.Payload
	SendPayload(p payload.Payload, f messainfromation.MessaInformation) error
}
