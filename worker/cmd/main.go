package main

import (
	"log"
	"time"
)

func main() {
	log.Println("Starting worker process...")
	log.Println("Worker is running. Press Ctrl+C to stop.")

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			log.Printf("Worker heartbeat: %s", time.Now().Format(time.RFC3339))
		}
	}
}
