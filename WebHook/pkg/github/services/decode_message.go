package services

import (
	"encoding/json"
	"fmt"
	"log"
	"triggo/pkg/github/model/push"
)

func (s *Services) DecodeMessage(event string, body []byte) {

	switch event {
	case "branch":
		fmt.Println("In this case the event it´s a branch")
	case "push":
		log.Println("In this case the event it´s a a push")

		var push push.GithubPush

		err := json.Unmarshal(body, &push)

		if err != nil {
			fmt.Println("Error to decode message")
			return
		}

		message := "Se ha hecho un cambio en la rama: " + push.Ref + "\nPor el siguiente usuario: " + push.Pusher.Name + "\n"

		log.Println(message)

	default:
		fmt.Printf("In this case the event is: %s \n", event)
	}

}
