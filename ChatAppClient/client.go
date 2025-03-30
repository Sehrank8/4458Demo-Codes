package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"os"
)

type Message struct {
	Sender string `json:"sender"`
	Text   string `json:"text"`
}

func main() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8081/ws", nil)
	if err != nil {
		log.Fatal("Error connecting to WebSocket server:", err)
	}
	defer conn.Close()

	fmt.Println("Connected to chat server. Type messages and press Enter to send.")

	go func() {
		for {
			var msg Message
			if err := conn.ReadJSON(&msg); err != nil {
				log.Println("Error reading message:", err)
				return
			}
			fmt.Printf("\n%s: %s\n> ", msg.Sender, msg.Text)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		text := scanner.Text()
		if text == "" {
			continue
		}

		msg := Message{Sender: "", Text: text}
		jsonMsg, _ := json.Marshal(msg)
		if err := conn.WriteMessage(websocket.TextMessage, jsonMsg); err != nil {
			log.Println("Error sending message:", err)
			return
		}
		fmt.Println("Message sent!")
	}
}
