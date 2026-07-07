package ports

import (
	embed "triggo/pkg/discord/model/Embed"
	"triggo/pkg/discord/model/payload"
)

type DiscordServices interface {
	CreateEmbed(event string, message string) embed.Embed
	CreateDiscordPayload(e embed.Embed) payload.Payload
}
