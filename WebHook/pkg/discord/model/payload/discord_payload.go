package payload

import "triggo/pkg/discord/model/embed"

type Payload struct {
	Username  string        `json:"username,omitempty"`
	AvatarUrl string        `json:"avatar_url,omitempty"`
	Embeds    []embed.Embed `json:"embeds,omitempty"`
}
