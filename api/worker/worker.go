package worker

import (
	"log"
	"time"

	"github.com/Berykwn/Queue/api/db"
)

func StartWorker() {
	go func() {
		for {
			// Ambil antrian dari database
			data, err := db.GetQueueData()
			if err != nil {
				log.Printf("Error fetching queue data: %v\n", err)
				continue
			}

			// Jika ada antrian yang belum diproses, proses
			if data != "" {
				processQueue(data)
				err := db.UpdateQueueProcessedAt(data)
				if err != nil {
					log.Printf("Error updating queue processed at time: %v\n", err)
				}
			} else {
				log.Println("No queue data found, waiting for the next check...")
				// Tunggu sebelum mengambil antrian berikutnya
				time.Sleep(3 * time.Second)
			}
		}
	}()
}

func processQueue(data string) {
	// Implementasi logika pemrosesan antrian di sini
	log.Printf("Processing queue data: %s\n", data)

	// Catat waktu selesai pemrosesan antrian
	err := db.UpdateQueueProcessedAt(data)
	if err != nil {
		log.Printf("Error updating queue processed at time: %v\n", err)
	}
}
