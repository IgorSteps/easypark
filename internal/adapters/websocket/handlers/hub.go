package handlers

import (
	"encoding/json"

	"github.com/IgorSteps/easypark/internal/adapters/websocket/models"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	logger *logrus.Logger
	// Registered clients, map from userID to client.
	clients map[uuid.UUID]*Client
	// Inbound messages from the clients.
	broadcast chan []byte
	// Register requests from the clients.
	register chan *Client
	// Unregister requests from clients.
	unregister chan *Client
}

func NewHub(l *logrus.Logger) *Hub {
	return &Hub{
		logger:     l,
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[uuid.UUID]*Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client.userID] = client
		case client := <-h.unregister:
			if _, ok := h.clients[client.userID]; ok {
				delete(h.clients, client.userID)
				close(client.send)
			}
		case message := <-h.broadcast:
			var modelMsg models.Message
			if err := json.Unmarshal(message, &modelMsg); err != nil {
				h.logger.WithError(err).Error("failed to unmarshal received message")
				continue // TODO: Find out how to handle error appropriately
			}

			if client, ok := h.clients[modelMsg.ReceiverID]; ok {
				select {
				case client.send <- []byte(modelMsg.Content):
				default:
					close(client.send)
					delete(h.clients, client.userID)
				}
			}
		}
	}
}
