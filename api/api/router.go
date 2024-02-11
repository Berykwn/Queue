// api/router.go

package api

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/produce", ProduceQueueHandler).Methods("POST")
	router.HandleFunc("/monitor", MonitorQueueHandler).Methods("GET") // Endpoint untuk memantau antrian
	return router
}
