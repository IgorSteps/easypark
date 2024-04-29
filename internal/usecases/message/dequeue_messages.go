package usecases

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// DequeueMessages provides business logic to dequeue messages(get and delete them from the datastore).
type DequeueMessages struct {
	logger      *logrus.Logger
	messageRepo repositories.MessageRepository
}

// NewDequeueMessages returns a new instance of DequeueMessages.
func NewDequeueMessages(l *logrus.Logger, mr repositories.MessageRepository) *DequeueMessages {
	return &DequeueMessages{
		logger:      l,
		messageRepo: mr,
	}
}

// Execute runs the business logic.
func (s *DequeueMessages) Execute(ctx context.Context, userID uuid.UUID) ([]entities.Message, error) {
	messages, err := s.messageRepo.GetManyForUser(userID)
	if err != nil {
		return []entities.Message{}, err
	}

	err = s.messageRepo.Delete(messages)
	if err != nil {
		return []entities.Message{}, err
	}

	return messages, nil
}
