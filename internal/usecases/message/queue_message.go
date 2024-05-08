package usecases

import (
	"context"
	"time"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// QueueMessage provides business logic to save messages.
type QueueMessage struct {
	logger      *logrus.Logger
	userRepo    repositories.UserRepository
	messageRepo repositories.MessageRepository
}

// NewQueueMessage returns a new instance of QueueMessage.
func NewQueueMessage(lgr *logrus.Logger, userRepo repositories.UserRepository, messageRepo repositories.MessageRepository) *QueueMessage {
	return &QueueMessage{
		logger:      lgr,
		userRepo:    userRepo,
		messageRepo: messageRepo,
	}
}

// Execute runs the business logic.
func (s *QueueMessage) Execute(ctx context.Context, senderID, receiverID uuid.UUID, content string) (entities.Message, error) {
	var sender entities.User
	err := s.userRepo.GetDriverByID(ctx, senderID, &sender)
	if err != nil {
		return entities.Message{}, err
	}

	var receiver entities.User
	err = s.userRepo.GetDriverByID(ctx, receiverID, &receiver)
	if err != nil {
		return entities.Message{}, err
	}

	if sender.Role == "driver" && receiver.Role != "admin" {
		s.logger.Error("drivers can only send messages to the admin")
		return entities.Message{}, repositories.NewInvalidInputError("drivers can only send messages to the admin")
	}

	message := entities.Message{
		ID:         uuid.New(),
		SenderID:   senderID,
		ReceiverID: receiverID,
		Content:    content,
		Delivered:  false,
		Timestamp:  time.Now(),
	}

	err = s.messageRepo.Create(&message)
	if err != nil {
		return entities.Message{}, err
	}

	return message, nil
}
