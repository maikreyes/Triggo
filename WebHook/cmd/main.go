package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"triggo/api"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		fmt.Println("Dont found .env %w", err)
	}

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
