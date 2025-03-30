package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

type Client struct {
	conn     *websocket.Conn
	username string
}

var (
	clients   = make(map[*Client]bool)
	broadcast = make(chan Message)
	mutex     sync.Mutex
	userCount = 0
)

type Message struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading connection:", err)
		return
	}
	defer conn.Close()

	mutex.Lock()
	userCount++
	username := fmt.Sprintf("User%d", userCount)
	client := &Client{conn: conn, username: username}
	clients[client] = true
	mutex.Unlock()

	log.Printf("%s connected\n", username)

	for {
		var msg Message
		if err := conn.ReadJSON(&msg); err != nil {
			log.Println("Error reading message:", err)
			mutex.Lock()
			delete(clients, client)
			mutex.Unlock()
			break
		}
		msg.Sender = client.username
		log.Printf("Received from %s: %s\n", msg.Sender, msg.Text)
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		log.Printf("Broadcasting: %s: %s\n", msg.Sender, msg.Text)
		mutex.Lock()
		for client := range clients {
			if err := client.conn.WriteJSON(msg); err != nil {
				log.Println("Error sending message:", err)
				client.conn.Close()
				delete(clients, client)
			}
		}
		mutex.Unlock()
	}
}

func main() {
	http.HandleFunc("/ws", handleConnections)
	go handleMessages()

	fmt.Println("WebSocket server started on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
