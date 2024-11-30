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
	connections sync.Map // map[string][]*websocket.Conn
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
		m.mu.Lock()
		if conns, ok := m.connections.Load(userID); ok {
			connections := conns.([]*websocket.Conn)
			// Remove this connection from the slice
			for i, c := range connections {
				if c == conn {
					connections = append(connections[:i], connections[i+1:]...)
					break
				}
			}
			if len(connections) == 0 {
				m.connections.Delete(userID)
			} else {
				m.connections.Store(userID, connections)
			}
		}
		m.mu.Unlock()
	}()

	m.mu.Lock()
	conns, _ := m.connections.LoadOrStore(userID, []*websocket.Conn{})
	connections := conns.([]*websocket.Conn)
	connections = append(connections, conn)
	m.connections.Store(userID, connections)
	m.mu.Unlock()

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

	connsVal, ok := m.connections.Load(userID)
	if !ok {
		return fmt.Errorf("client %s not connected", userID)
	}

	connections := connsVal.([]*websocket.Conn)
	if len(connections) == 0 {
		return fmt.Errorf("no active connections for client %s", userID)
	}

	message, err := json.Marshal(notification)
	if err != nil {
		return fmt.Errorf("failed to marshal notification: %v", err)
	}

	var lastErr error
	for _, conn := range connections {
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			lastErr = err
		}
	}

	return lastErr
}

// BroadcastMessage sends a notification to all connected clients
func (m *WebSocketManager) BroadcastMessage(message interface{}) error {
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("error marshalling broadcast message: %v", err)
	}

	m.mu.RLock()
	defer m.mu.RUnlock()

	var lastErr error
	m.connections.Range(func(key, value interface{}) bool {
		connections := value.([]*websocket.Conn)
		for _, conn := range connections {
			if err := conn.WriteMessage(websocket.TextMessage, jsonMessage); err != nil {
				lastErr = err
			}
		}
		return true
	})

	return lastErr
}

// GetActiveConnections returns the count of active connections
func (m *WebSocketManager) GetActiveConnections() int {
	count := 0
	m.connections.Range(func(key, value interface{}) bool {
		if conns, ok := value.([]*websocket.Conn); ok {
			count += len(conns)
		}
		return true
	})
	return count
}
