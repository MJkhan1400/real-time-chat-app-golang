package server

import (
	"encoding/json"
	"log"

	"github.com/MJkhan1400/real-time-chat-app-golang/internal/models"
	"github.com/gorilla/websocket"
)

func (c *Client) Read(manager *ClientManager) {
	defer func() {
		manager.Unregister <- c
		c.Socket.Close()
	}()

	for {
		_, msg, err := c.Socket.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			break
		}

		message, _ := json.Marshal(models.Message{
			Sender: c.ID,
			Content: string(msg),
		})

		manager.Brodcast <- message
	}

	func (c *Client) Write() {
		defer c.Socket.Close()

		for msg := range c.Send {
			err := c.Socket.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				return
			}
		}
	}
}
