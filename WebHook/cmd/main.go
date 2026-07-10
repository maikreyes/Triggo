package main

import (
	"log"
	"net/http"
	"os"

	"triggo/api"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/api/setup", api.SetupHandler)
	mux.HandleFunc("/api/webhook", api.Webhook)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Servidor local corriendo en http://localhost:%s\n", port)

	log.Fatal(http.ListenAndServe(":"+port, mux))
}
