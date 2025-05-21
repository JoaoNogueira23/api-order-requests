package websocket

import (
	"fmt"
	"net/http"

	"sync"

	"github.com/gorilla/websocket"
)

var Broadcast = make(chan []byte)            // canal de transferência de mensagens
var Clients = make(map[*websocket.Conn]bool) // lista de clientes conectados
var mutex = &sync.Mutex{}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func GetWsHandler() http.HandlerFunc {
	return WsHandler
}

func WriteMessage(conn *websocket.Conn, messageType int, msg []byte) error {
	err := conn.WriteMessage(messageType, msg)
	if err != nil {
		fmt.Println("Error writing message:", err)
		return err
	}
	return nil
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	// upgrade the HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error during connection upgrade:", err)
		return
	}
	defer conn.Close()
	mutex.Lock()
	Clients[conn] = true // add new client to the list
	mutex.Unlock()
	//Listen for incoming messages
	for {
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			mutex.Lock()
			fmt.Println("Error reading message:", err)
			delete(Clients, conn) // remove client from the list
			mutex.Unlock()
			break
		}

		fmt.Printf("Received message: %s\n", msg)

		err = conn.WriteMessage(messageType, msg)
		if err != nil {
			fmt.Println("Error writing message:", err)
			break
		}
	}
}

func StartBrodcast() {
	for {
		msg := <-Broadcast
		mutex.Lock()
		fmt.Printf("Broadcasting message: %s\n", msg)
		for client := range Clients {
			err := WriteMessage(client, websocket.TextMessage, msg)
			if err != nil {
				fmt.Println("Error writing message to client:", err)
				client.Close()
				delete(Clients, client)
			}
			mutex.Unlock()
		}
	}
}
