package services

import (
	embed "triggo/pkg/discord/model/Embed"
)

func (s *Services) CreateEmbed(event string, message string) embed.Embed {

	switch event {
	case "push":
		return embed.Embed{
			Title:       "New Push",
			Description: message,
			Color:       0x33FF57,
		}
	default:
		return embed.Embed{
			Title:       "New " + event,
			Description: message,
			Color:       0xFFFF,
		}

	}

}
