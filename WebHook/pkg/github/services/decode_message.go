package services

import "fmt"

func (s *Services) DecodeMessage(event string, body string) {

	fmt.Println("=== PAYLOAD RECIBIDO ===")
	fmt.Println(body)
	fmt.Println("========================")

	switch event {
	case "branch":
		fmt.Println("In this case the event it´s a branch")
	case "push":
		fmt.Println("In this case the event it´s a a push")
	default:
		fmt.Printf("In this case the event is: %s \n", event)
	}

}
