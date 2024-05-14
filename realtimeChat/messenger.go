package realtimechat

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrad = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Messages struct {
	Content string `json:"content"`
}

var (
	clientsMu sync.Mutex
	Clients   []*websocket.Conn
)

func (m *Messenger) HTTPServe() http.Handler {
	return http.HandlerFunc(HandleMessages)
}

func (m *Messenger) EndPoint() string {
	return "/message"
}

func HandleMessages(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrad.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not upgrade to WebSocket connection", http.StatusBadRequest)
		return
	}
	defer conn.Close()

	clientsMu.Lock()
	Clients = append(Clients, conn)
	clientsMu.Unlock()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("Error reading message: %s\n", err)
			break
		}

		var receivedMsg Messages
		if err := json.Unmarshal(msg, &receivedMsg); err != nil {
			fmt.Printf("Error unmarshalling JSON: %s\n", err)
			continue
		}
		
		clientsMu.Lock()
		for _, client := range Clients {
			if err := client.WriteMessage(websocket.TextMessage, []byte(receivedMsg.Content)); err != nil {
				fmt.Printf("Error writing message: %s\n", err)
			}
		}
		clientsMu.Unlock()
	}

	conn.Close()
}


