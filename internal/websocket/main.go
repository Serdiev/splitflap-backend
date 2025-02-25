package ws

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"splitflap-backend/internal/logger"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebSocket struct {
	clients          map[*websocket.Conn]bool // Track active WebSocket clients
	HandleNewMessage *func(obj []byte)
	mutex            sync.Mutex
	upgrader         websocket.Upgrader
}

func NewWebsocket() *WebSocket {
	return &WebSocket{
		clients:          map[*websocket.Conn]bool{},
		HandleNewMessage: nil,
		mutex:            sync.Mutex{},
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

func (w *WebSocket) BroadcastMessage(msg []byte) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	for client := range w.clients {
		err := client.WriteMessage(1, msg)
		if err != nil {
			log.Printf("Error sending message to client: %v", err)
			client.Close()
			delete(w.clients, client)
		}
	}
}

// WebSocket handler for upgrades and real-time messaging
func (w *WebSocket) HandleWebSocket(c *gin.Context) {
	conn, err := w.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Failed to set WebSocket upgrade:", err)
		return
	}
	defer conn.Close()

	// Register new client
	w.mutex.Lock()
	w.clients[conn] = true
	w.mutex.Unlock()

	for {
		var msg any
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error: %v", err)
			w.mutex.Lock()
			delete(w.clients, conn)
			w.mutex.Unlock()
			break
		}

		bytes, err := json.Marshal(msg)
		if err != nil {
			return
		}

		if w.HandleNewMessage != nil {
			(*w.HandleNewMessage)(bytes)
		} else {
			logger.Error().Msg("Do not have any handler for incoming websocket messages")
		}
	}
}

func ParseMessage[T any](msg []byte) (res *T, err error) {
	err = json.Unmarshal(msg, &res)
	return res, err
}

func ToBytes(msg any) []byte {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return []byte{}
	}
	return bytes
}
