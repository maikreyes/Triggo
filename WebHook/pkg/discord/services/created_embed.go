package services

import "triggo/pkg/discord/model/embed"

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
