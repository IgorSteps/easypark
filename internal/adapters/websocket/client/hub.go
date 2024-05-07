package client

import (
	"context"

	"github.com/IgorSteps/easypark/internal/adapters/websocket/models"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// Hub maintains the set of active clients, registers and unregisters clients, and broadcasts messages to them.
type Hub struct {
	// Registered Clients by their user ids.
	Clients map[uuid.UUID]*Client

	logger *logrus.Logger
	facade MessageFacade

	// Inbound messages from the clients.
	broadcast chan *models.Message
	// Register requests from the clients.
	Register chan *Client
	// Unregister requests from clients.
	unregister chan *Client
}

// NewHub returns a new instance of Hub.
func NewHub(l *logrus.Logger, f MessageFacade) *Hub {
	return &Hub{
		logger:     l,
		facade:     f,
		Clients:    make(map[uuid.UUID]*Client),
		broadcast:  make(chan *models.Message),
		Register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

// Run listens on the Hub channels to handle client connections and message broadcasting.
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client.UserID] = client

			//Dequeue messages if any.
			err := h.dequeueMessages(client)
			if err != nil {
				h.logger.WithError(err).WithField("client id", client.UserID).Error("failed to dequeue messages")
				// Send the error back to the client.
				client.Send <- &models.Message{
					SenderID: client.UserID,
					Content:  "An error ocurred, please try again later.",
				}
			}
			h.logger.Debug("registered client")

		case client := <-h.unregister:
			if _, ok := h.Clients[client.UserID]; ok {
				delete(h.Clients, client.UserID)
				close(client.Send)
			}
			h.logger.Debug("unregistered client")

		case message := <-h.broadcast:
			h.broadcastMessage(message)
		}
	}
}

func (h *Hub) broadcastMessage(msg *models.Message) {
	// Direct message recipient if they are registered with the Hub(they're online).
	if receiverClient, ok := h.Clients[msg.ReceiverID]; ok {
		select {
		case receiverClient.Send <- msg:
			h.logger.WithFields(logrus.Fields{
				"msg":        msg,
				"senderID":   msg.SenderID,
				"receivedID": msg.ReceiverID,
			}).Debug("broadcast message to receiver")
		default: // If send buffer is full.
			// Unregister the client and close the connection.
			close(receiverClient.Send)
			delete(h.Clients, receiverClient.UserID)
			h.logger.Warn("sender buffer is full, unregistered client and closed send channel ")
		}
	} else {
		// Persist message if the recipient is not registered with the Hub(ie. they're offline).
		h.enqueueMessage(msg)
	}

	// Broadcast senders message to themselves.
	if senderClient, ok := h.Clients[msg.SenderID]; ok {
		select {
		case senderClient.Send <- msg:
			h.logger.WithFields(logrus.Fields{
				"msg":        msg,
				"senderID":   msg.SenderID,
				"receivedID": msg.ReceiverID,
			}).Debug("broadcast message to sender")
		default: // If send buffer is full.
			// Unregister the client and close the connection.
			close(senderClient.Send)
			delete(h.Clients, senderClient.UserID)
			h.logger.Warn("sender buffer is full, unregistered client and closed send channel ")
		}
	}
}

func (h *Hub) enqueueMessage(modelMsg *models.Message) {
	_, err := h.facade.EnqueueMessage(context.TODO(), modelMsg.SenderID, modelMsg.ReceiverID, modelMsg.Content)
	if err != nil {
		h.logger.WithError(err).Error("failed to enqueue message")

		// Notify the sender enqueueing failed.
		if sender, ok := h.Clients[modelMsg.SenderID]; ok {
			h.logger.WithField("client id", sender.UserID).Error("enqueueing failed")
			sender.Send <- &models.Message{
				SenderID: modelMsg.SenderID,
				Content:  "An error ocurred, please try again later.",
			}
		} else {
			h.logger.Warn("enqueueing failed, tried to send the error to the sender, but couldn't find them in registered clients")
		}
	}

	h.logger.WithFields(logrus.Fields{
		"recipient id": modelMsg.ReceiverID,
		"message":      modelMsg.Content,
	}).Debug("recipient is offline, enqueued this message")
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
		client.Send <- models.FromDomain(msg)
	}

	return nil
}
