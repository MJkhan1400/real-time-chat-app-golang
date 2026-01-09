package main

import (
	"log"
	"net/http"

	"github.com/MJkhan1400/real-time-chat-app-golang/internal/server"
)

func main() {
	manager := &server.ClientManager{
		Clients:    make(map[*server.Client]bool),
		Register:   make(chan *server.Client),
		Unregister: make(chan *server.Client),
		Brodcast:   make(chan []byte),
	}

	go manager.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		server.ServerWS(manager, w, r)
	})

	log.Println("Server started on : 12345")
	http.ListenAndServe(":12345", nil)
}
