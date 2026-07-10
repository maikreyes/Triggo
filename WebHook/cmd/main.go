package main

import (
	"log"
	"net/http"
	"os"

	// Importas tu paquete api (ajusta "triggo/api" según el nombre real de tu módulo)
	"triggo/api"
)

func main() {

	// 2. Crear un enrutador local
	mux := http.NewServeMux()

	// 3. Conectar tus funciones de Vercel a las rutas locales
	mux.HandleFunc("/api/setup", api.SetupHandler)
	mux.HandleFunc("/api/webhook", api.Webhook)

	// 4. Iniciar el servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Servidor local corriendo en http://localhost:%s\n", port)

	// La magia de Go: esto mantiene el servidor vivo escuchando peticiones
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
