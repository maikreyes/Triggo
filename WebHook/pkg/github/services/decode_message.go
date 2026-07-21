package services

import (
	"encoding/json"
	"fmt"
	messainfromation "triggo/pkg/github/model/messa_infromation"
	"triggo/pkg/github/model/push"
)

func (s *Services) DecodeMessage(event string, body []byte) (messainfromation.MessaInformation, string) {

	switch event {
	case "branch":
		fmt.Println("In this case the event it´s a branch")
	case "push":

		var push push.GithubPush
		var info messainfromation.MessaInformation

		err := json.Unmarshal(body, &push)

		if err != nil {
			fmt.Println("Error to decode message")
			return messainfromation.MessaInformation{}, ""
		}

		message := "Se ha hecho un cambio en la rama: " + push.Ref + "\nPor el siguiente usuario: " + push.Pusher.Name + "\n"

		info = messainfromation.MessaInformation{
			Installation: push.Installation,
			Repository:   push.Repository,
		}

		return info, message

	default:
		return messainfromation.MessaInformation{}, "In this case the event is: %s \n" + event
	}
	return messainfromation.MessaInformation{}, ""
}
