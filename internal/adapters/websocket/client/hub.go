package client

import (
	"context"
	"encoding/json"

	"github.com/IgorSteps/easypark/internal/adapters/websocket/models"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// Hub maintains the set of active clients and broadcasts messages to the them.
type Hub struct {
	logger *logrus.Logger
	facade MessageFacade

	// Registered Clients by their user ids.
	Clients map[uuid.UUID]*Client
	// Inbound messages from the clients.
	Broadcast chan []byte
	// Register requests from the clients.
	Register chan *Client
	// Unregister requests from clients.
	Unregister chan *Client
}

func NewHub(l *logrus.Logger, f MessageFacade) *Hub {
	return &Hub{
		logger:     l,
		facade:     f,
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[uuid.UUID]*Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client.UserID] = client

			// Dequeue messages if any.
			err := h.dequeueMessages(client)
			if err != nil {
				// Send the error back to the client.
				client.Send <- []byte(err.Error())
			}

		case client := <-h.Unregister:
			if _, ok := h.Clients[client.UserID]; ok {
				delete(h.Clients, client.UserID)
				close(client.Send)
			} else {
				h.logger.WithField("user id", client.UserID).Warn("tried to unregister non registered client")
			}

		case message := <-h.Broadcast:
			var modelMsg models.Message
			if err := json.Unmarshal(message, &modelMsg); err != nil {
				h.logger.WithField("raw message", string(message)).WithError(err).Error("failed to unmarshal received message")
				continue // TODO: Find out how to handle this error appropriately.
			}

			// Direct message if recipient is registered with the Hub(they're online).
			if client, ok := h.Clients[modelMsg.ReceiverID]; ok {
				select {
				case client.Send <- []byte(modelMsg.Content):
				default:
					close(client.Send)
					delete(h.Clients, client.UserID)
				}
			} else {
				// Persist message if the user is not registered with the Hub(ie. they're offline).
				_, err := h.facade.EnqueueMessage(context.TODO(), modelMsg.SenderID, modelMsg.ReceiverID, modelMsg.Content)
				if err != nil {
					h.logger.WithError(err).Error("failed to enqueue message")

					// Notify the sender enqueueing failed.
					if sender, ok := h.Clients[modelMsg.SenderID]; ok {
						sender.Send <- []byte("failed to enqueue message")
					} else {
						h.logger.Warn("sender is not found in registered clients")
						continue
					}
				}

				h.logger.WithFields(logrus.Fields{
					"recipient id": modelMsg.ReceiverID,
					"message":      modelMsg.Content,
				}).Debug("recipient is offline, enqueued this message")
			}
		}
	}
}

func (h *Hub) dequeueMessages(client *Client) error {
	messages, err := h.facade.DequeueMessages(context.TODO(), client.UserID)
	if err != nil {
		h.logger.WithError(err).Error("failed to dequeue messages")
		return err
	}
	h.logger.WithFields(logrus.Fields{
		"userID":   client.UserID,
		"messages": messages,
	}).Debug("dequeued messages for user")

	// Send each message to the client's send channel.
	for _, msg := range messages {
		client.Send <- []byte(msg.Content)
	}

	return nil
}
