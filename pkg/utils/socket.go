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

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var connections sync.Map

func HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}
	defer conn.Close()

	clientID := c.Query("user_id")

	connections.Store(clientID, conn)
	defer connections.Delete(clientID)

	go func() {
		for {
			time.Sleep(30 * time.Second)
			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				fmt.Println("Ping failed:", err)
				return
			}
		}
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Printf("Received from %s: %s\n", clientID, string(msg))
	}
}

func SendMessageToClient(clientID string, notification models.GetMyNotificationsResponse) error {
	message, _ := json.Marshal(notification)
	if conn, ok := connections.Load(clientID); ok {
		wsConn := conn.(*websocket.Conn)
		return wsConn.WriteMessage(websocket.TextMessage, []byte(message))
	}
	return fmt.Errorf("client %s not connected", clientID)
}
