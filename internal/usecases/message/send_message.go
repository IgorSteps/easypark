package usecases

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/sirupsen/logrus"
)

// CreateMessage provides business logic to create messages.
type CreateMessage struct {
	logger      *logrus.Logger
	messageRepo repositories.MessageRepository
	userRepo    repositories.UserRepository
}

// NewCreateMessage returns a new instance of CreateMessage.
func NewCreateMessage(l *logrus.Logger, messageRepo repositories.MessageRepository, userRepo repositories.UserRepository) *CreateMessage {
	return &CreateMessage{
		logger:      l,
		messageRepo: messageRepo,
		userRepo:    userRepo,
	}
}

// Execute runs the business logic.
func (s *CreateMessage) Execute(ctx context.Context, msg entities.Message) error {
	var sender entities.User
	err := s.userRepo.GetSingle(ctx, msg.SenderID, &sender)
	if err != nil {
		return err
	}

	var admin entities.User
	err = s.userRepo.GetAdmin(ctx, &admin)
	if err != nil {
		return err
	}

	// Check if the sender is the Admin.
	if sender.Role == entities.RoleAdmin {
		return s.messageRepo.Send(msg, msg.RecipientID)
	} else {
		// Driver sending to the admin
		return s.messageRepo.Send(msg, admin.ID)
	}
}
