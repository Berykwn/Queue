// main.go

package main

import (
	"log"
	"net/http"

	"github.com/Berykwn/Queue/api/api"

	"github.com/Berykwn/Queue/api/worker"
)

func main() {
	// Start Worker
	worker.StartWorker()

	// Inisialisasi router
	router := api.NewRouter()

	// Mulai server HTTP
	log.Fatal(http.ListenAndServe(":8080", router))
}
