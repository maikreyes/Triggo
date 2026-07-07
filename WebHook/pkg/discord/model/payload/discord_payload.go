package payload

import embed "triggo/pkg/discord/model/Embed"

type Payload struct {
	Username  string        `json:"username,omitempty"`
	AvatarUrl string        `json:"avatar_url,omitempty"`
	Embeds    []embed.Embed `json:"embeds,omitempty"`
}
