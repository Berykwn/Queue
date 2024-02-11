package api

import (
	"encoding/json"
	"net/http"

	"github.com/Berykwn/Queue/api/db"
)

func ProduceQueueHandler(w http.ResponseWriter, r *http.Request) {
	// Mendekode payload dari request
	type RequestData struct {
		Data string `json:"data"`
	}
	var requestData RequestData
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Simpan ke dalam database
	database := db.GetDB()
	_, err = database.Exec("INSERT INTO queue (data) VALUES (?)", requestData.Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Kirim respons
	response := map[string]string{"message": "Queue produced successfully"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func MonitorQueueHandler(w http.ResponseWriter, r *http.Request) {
	// Query database untuk mendapatkan informasi tentang antrian
	database := db.GetDB()
	var totalQueue int
	err := database.QueryRow("SELECT COUNT(*) FROM queue").Scan(&totalQueue)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Kirim respons
	response := map[string]int{"total_queue": totalQueue}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
