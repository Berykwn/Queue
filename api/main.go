package main

import (
	"log"
	"net/http"

	"github.com/Berykwn/Queue/api/api"
	"github.com/Berykwn/Queue/api/worker"

	"github.com/gorilla/handlers"
	"github.com/rs/cors"
)

func main() {
	// Start Worker
	worker.StartWorker()

	// Inisialisasi router
	router := api.NewRouter()

	corsHandler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(corsHandler)))
}
