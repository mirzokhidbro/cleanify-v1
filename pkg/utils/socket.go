package utils

import (
	"bw-erp/models"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// WebSocketManager handles WebSocket connections
type WebSocketManager struct {
	upgrader    websocket.Upgrader
	connections sync.Map
	mu          sync.RWMutex
}

var (
	manager *WebSocketManager
	once    sync.Once
)

// GetManager returns a singleton instance of WebSocketManager
func GetManager() *WebSocketManager {
	once.Do(func() {
		manager = &WebSocketManager{
			upgrader: websocket.Upgrader{
				ReadBufferSize:  1024,
				WriteBufferSize: 1024,
				CheckOrigin: func(r *http.Request) bool {
					return true
				},
			},
		}
	})
	return manager
}

// HandleConnection handles new WebSocket connections
func (m *WebSocketManager) HandleConnection(c *gin.Context, userID string) error {
	conn, err := m.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return fmt.Errorf("WebSocket upgrade failed: %v", err)
	}

	m.handleClient(userID, conn)
	return nil
}

// handleClient manages individual WebSocket client connections
func (m *WebSocketManager) handleClient(userID string, conn *websocket.Conn) {
	defer func() {
		conn.Close()
		m.connections.Delete(userID)
	}()

	m.connections.Store(userID, conn)

	// Start keepalive
	go m.keepAlive(conn)

	// Read messages (required to keep connection alive)
	for {
		messageType, _, err := conn.ReadMessage()
		if err != nil || messageType == websocket.CloseMessage {
			return
		}
	}
}

// keepAlive sends periodic ping messages
func (m *WebSocketManager) keepAlive(conn *websocket.Conn) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
			return
		}
	}
}

// SendMessage sends a notification to a specific client
func (m *WebSocketManager) SendMessage(userID string, notification models.GetMyNotificationsResponse) error {
	m.mu.RLock()
	defer m.mu.RUnlock()

	connVal, ok := m.connections.Load(userID)
	if !ok {
		return fmt.Errorf("client %s not connected", userID)
	}

	conn, ok := connVal.(*websocket.Conn)
	if !ok {
		return fmt.Errorf("invalid connection type for client %s", userID)
	}

	message, err := json.Marshal(notification)
	if err != nil {
		return fmt.Errorf("failed to marshal notification: %v", err)
	}

	return conn.WriteMessage(websocket.TextMessage, message)
}

// GetActiveConnections returns the count of active connections
func (m *WebSocketManager) GetActiveConnections() int {
	count := 0
	m.connections.Range(func(_, _ interface{}) bool {
		count++
		return true
	})
	return count
}
