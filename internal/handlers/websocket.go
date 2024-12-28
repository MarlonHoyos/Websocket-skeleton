package handlers

import (
	"log"
	"net/http"
	"WEBSOCKET-SKELETON/internal/infrastructure"
	"github.com/gorilla/websocket"
)

// Configuración del upgrader WebSocket

var (
	upgrader       = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
	webSocketManager = infrastructure.NewWebSocketManager() // Gestor de conexiones
)


func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}
	defer conn.Close()

	// Simula obtener un userID (autenticación real debería hacerse aquí)
	userID := r.URL.Query().Get("userID")
	webSocketManager.AddConnection(userID, conn)
	defer webSocketManager.RemoveConnection(userID)

	log.Printf("User %s connected!", userID)


	for {
		var event struct {
			Type string                 `json:"type"`
			Data map[string]interface{} `json:"data"`
		}
		if err := conn.ReadJSON(&event); err != nil {
			log.Printf("Error reading event: %v", err)
			break
		}

		switch event.Type {
		case "on-connected-user":
			handleConnectedUser(event.Data, conn)
		case "on-send-message":
			handleSendMessage(event.Data)
		default:
			log.Printf("Unknown event type: %s", event.Type)
		}
	}

}

func handleConnectedUser(data map[string]interface{}, conn *websocket.Conn) {
	log.Printf("User connected: %v", data)
	_ = conn.WriteJSON(map[string]string{"message": "Welcome!"})
}

func handleSendMessage(data map[string]interface{}) {
	log.Printf("Message sent: %v", data)
	// Implementa lógica adicional aquí (persistencia, notificación, etc.)
}
