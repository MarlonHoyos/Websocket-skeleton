package main

import (
	"WEBSOCKET-SKELETON/internal/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ws", handlers.WebSocketHandler)

	log.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}
