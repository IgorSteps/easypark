package websocket

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// Hub manages clients and broadcasts messages to the
// clients.
type Hub struct {
	logger *logrus.Logger
	// Registered clients.
	clients map[uuid.UUID]*Client
	// Register requests from the clients.
	register chan *Client
	// Unregister requests from clients.
	unregister chan *Client
	// Inbound messages from the clients.
	broadcast chan []byte
}

func NewHub(l *logrus.Logger) *Hub {
	return &Hub{
		logger:     l,
		clients:    make(map[uuid.UUID]*Client),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte),
	}
}

func (s *Hub) Run() {
	for {
		select {
		// Register client
		case client := <-s.register:
			s.clients[client.id] = client
		// Unregister client
		case client := <-s.unregister:
			if _, ok := s.clients[client.id]; ok {
				delete(s.clients, client.id)
				close(client.send)
			}
		// Broadcast messages
		case message := <-s.broadcast:
			for _, client := range s.clients {
				select {
				case client.send <- message:
				}
			}
		}
	}
}
