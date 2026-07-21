package services

import (
	"triggo/pkg/discord/model/embed"
	"triggo/pkg/discord/model/payload"
)

func (s *Services) CreateDiscordPayload(e embed.Embed) payload.Payload {

	payload := payload.Payload{
		Username:  "Informante Triggo",
		AvatarUrl: "https://ysqz0oydi7thsqmt.public.blob.vercel-storage.com/Triggo.png",
	}

	payload.Embeds = []embed.Embed{e}

	return payload

}
