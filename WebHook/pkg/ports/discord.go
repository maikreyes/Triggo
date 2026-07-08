package ports

import (
	"triggo/pkg/discord/model/embed"
	"triggo/pkg/discord/model/payload"
)

type DiscordServices interface {
	CreateEmbed(event string, message string) embed.Embed
	CreateDiscordPayload(e embed.Embed) payload.Payload
	SendPayload(p payload.Payload) error
}
