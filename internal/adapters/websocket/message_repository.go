package websocket

import (
	"errors"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type MessageWebSocketRepository struct {
	logger *logrus.Logger
	hub    *Hub
}

// NewWebSocketMessageRepository returns a new instance of WebSocketMessageRepository.
func NewWebSocketMessageRepository(logger *logrus.Logger, hub *Hub) *MessageWebSocketRepository {
	return &MessageWebSocketRepository{
		logger: logger,
		hub:    hub,
	}
}

// SendToDriver sends a message to a specific user via WebSocket.
func (repo *MessageWebSocketRepository) Send(msg entities.Message, userID uuid.UUID) error {
	client, ok := repo.hub.clients[userID]
	if !ok {
		repo.logger.Error("Driver not connected or does not exist")
		return errors.New("driver not connected or does not exist")
	}
	// What should I do now?
	message := msg.Content // Assuming msg.Content is []byte
	select {
	case client.send <- message:
	default:
		repo.logger.Error("Failed to send message to user: ", userID)
		return errors.New("failed to send message")
	}

	return nil
}
