package server

type ClientManager struct {
	Clients    map[*Client]bool
	Register   chan *Client
	Unregister chan *Client
	Brodcast   chan []byte
}

func (manager *ClientManager) Start() {
	for {
		select {
		case client := <-manager.Register:
			manager.Clients[client] = true

		case client := <-manager.Unregister:
			if _, ok := manager.Clients[client]; ok {
				close(client.Send)
				delete(manager.Clients, client)
			}

		case message := <-manager.Brodcast:
			for client := range manager.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(manager.Clients, client)
				}
			}
		}
	}
}
