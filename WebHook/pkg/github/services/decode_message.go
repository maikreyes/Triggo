package services

import (
	"encoding/json"
	"fmt"
	"triggo/pkg/github/model/push"
)

func (s *Services) DecodeMessage(event string, body []byte) string {

	switch event {
	case "branch":
		fmt.Println("In this case the event it´s a branch")
	case "push":

		var push push.GithubPush

		err := json.Unmarshal(body, &push)

		if err != nil {
			fmt.Println("Error to decode message")
			return ""
		}

		message := "Se ha hecho un cambio en la rama: " + push.Ref + "\nPor el siguiente usuario: " + push.Pusher.Name + "\n"

		return message

	default:
		return "In this case the event is: %s \n" + event
	}
	return ""
}
