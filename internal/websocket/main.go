package ws

import (
	"encoding/json"
	"net/http"
	"sync"

	"splitflap-backend/internal/logger"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebSocket struct {
	clients          map[*websocket.Conn]bool
	HandleNewMessage *func(obj []byte)
	mutex            sync.Mutex
	upgrader         websocket.Upgrader
}

func NewWebsocket(allowedOrigins []string) *WebSocket {
	return &WebSocket{
		clients:          map[*websocket.Conn]bool{},
		HandleNewMessage: nil,
		mutex:            sync.Mutex{},
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				origin := r.Header.Get("Origin")
				for _, allowed := range allowedOrigins {
					if origin == allowed {
						return true
					}
				}
				return false
			},
		},
	}
}

func (w *WebSocket) BroadcastMessageAsText(msg []byte) {
	w.broadcastMessage(msg, 1)
}

func (w *WebSocket) BroadcastMessageAsBinary(msg []byte) {
	w.broadcastMessage(msg, 2)
}

func (w *WebSocket) broadcastMessage(msg []byte, msgType int) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	for client := range w.clients {
		err := client.WriteMessage(msgType, msg)
		if err != nil {
			logger.Error().Err(err).Msg("WebSocket: error sending message to client")
			client.Close()
			delete(w.clients, client)
		}
	}
}

// WebSocket handler for upgrades and real-time messaging
func (w *WebSocket) HandleWebSocket(c *gin.Context) {
	conn, err := w.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
			logger.Error().Err(err).Msg("WebSocket: failed to set upgrade")
			return
		}
	defer conn.Close()

	// Register new client
	w.mutex.Lock()
	w.clients[conn] = true
	w.mutex.Unlock()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			logger.Error().Err(err).Msg("WebSocket: read error")
			w.mutex.Lock()
			delete(w.clients, conn)
			w.mutex.Unlock()
			break
		}

		text := string(msg)
		logger.Info().Str("message", text).Msg("WebSocket: received")
	}
}

func ToBytes(msg any) []byte {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return []byte{}
	}
	return bytes
}

func (w *WebSocket) ClientCount() int {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	return len(w.clients)
}
