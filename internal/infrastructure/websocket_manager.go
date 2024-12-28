package infrastructure

import (
	"github.com/gorilla/websocket"
	"sync"
)

type WebSocketManager struct {
	connections map[string]*websocket.Conn
	mu          sync.RWMutex
}

func NewWebSocketManager() *WebSocketManager  {
	return &WebSocketManager{
		connections: make(map[string]*websocket.Conn),
	}
}

func (wm *WebSocketManager) AddConnection(userID string, conn *websocket.Conn) {
	wm.mu.Lock()
	defer wm.mu.Unlock()
	wm.connections[userID] = conn
}

func (wm *WebSocketManager) RemoveConnection(userID string) {
	wm.mu.Lock()
	defer wm.mu.Unlock()
	delete(wm.connections, userID)
}

func (wm *WebSocketManager) SendMessage(userID string, message interface{}) error {
	wm.mu.RLock()
	defer wm.mu.RUnlock()

	conn, exists := wm.connections[userID]
	if !exists {
		return nil
	}

	return conn.WriteJSON(message)
}
