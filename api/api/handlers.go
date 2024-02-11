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
	_, err = database.Exec("INSERT INTO queue (data, status) VALUES (?, ?)", requestData.Data, "pending") // Menambahkan nilai status "pending"
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
	rows, err := database.Query("SELECT id, data, processed_at FROM queue")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type Queue struct {
		ID          int    `json:"id"`
		Data        string `json:"data"`
		ProcessedAt string `json:"processed_at"`
	}

	var queues []Queue
	for rows.Next() {
		var queue Queue
		err := rows.Scan(&queue.ID, &queue.Data, &queue.ProcessedAt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		queues = append(queues, queue)
	}

	// Kirim respons
	jsonResponse, err := json.Marshal(queues)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
