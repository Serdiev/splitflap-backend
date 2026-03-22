package ws

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

func (w *WebSocket) BroadcastMessageAsText(msg []byte) {
	fmt.Println("braoadcasting text message")
	fmt.Println(len(msg))
	w.broadcastMessage(msg, 1)
}

func (w *WebSocket) BroadcastMessageAsBinary(msg []byte) {
	fmt.Println("braoadcasting binary message")
	fmt.Println(len(msg))
	w.broadcastMessage(msg, 2)
}

func (w *WebSocket) broadcastMessage(msg []byte, msgType int) {
	// 42 size of the json to webpage splitflap current text message
	// if len(msg) != 42 {
	// 	fmt.Println("broadcasting", len(msg))

	// 	// print only first 25 bytes (or all if shorter)
	// 	limit := 1000
	// 	if len(msg) < limit {
	// 		limit = len(msg)
	// 	}

	// 	fmt.Print("First bytes: [")
	// 	for i := 0; i < limit; i++ {
	// 		if i > 0 {
	// 			fmt.Print(" ")
	// 		}
	// 		fmt.Print(msg[i])
	// 	}
	// 	fmt.Println("]")
	// }

	w.mutex.Lock()
	defer w.mutex.Unlock()

	for client := range w.clients {
		err := client.WriteMessage(msgType, msg)
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
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("Error: %v", err)
			w.mutex.Lock()
			delete(w.clients, conn)
			w.mutex.Unlock()
			break
		}

		text := string(msg) // convert bytes to string
		fmt.Printf("Received: %s", text)
		// bytes, err := json.Marshal()
		// if err != nil {
		// 	return
		// }

		// if w.HandleNewMessage != nil {
		// 	(*w.HandleNewMessage)(bytes)
		// }
		// else {
		// 	logger.Error().Msg("Do not have any handler for incoming websocket messages")
		// }
	}
}

func ToBytes(msg any) []byte {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return []byte{}
	}
	return bytes
}
