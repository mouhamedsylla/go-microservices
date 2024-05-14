package realtimechat

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var clients []websocket.Conn

type Messag struct {
	Content string `json:"content"`
}

func (m *Messenger) HTTPServe() http.Handler {
	return http.HandlerFunc(m.Messenger)
}

func (m *Messenger) EndPoint() string {
	return "/discussion"
}

func (m *Messenger) Messenger(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "welcome to the discussion microservices...")
	//conn, _ := upgrader.Upgrade(w, r, nil)
	conn, _ := upgrader.Upgrade(w, r, nil)

	clients = append(clients, *conn)
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}

		var receivedMsg Message
		if err := json.Unmarshal(msg, &receivedMsg); err != nil {
			fmt.Println("Error umarshalling JSON: ", err)
			continue
		}

		fmt.Printf("Message sended by %s: %s", conn.RemoteAddr(), receivedMsg.Content)

		for _, client := range clients {
			if err := client.WriteMessage(websocket.TextMessage, []byte(receivedMsg.Content)); err != nil {
				fmt.Println("Error writting message:", err)
			}
		}
	}

}
