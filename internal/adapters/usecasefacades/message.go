package usecasefacades

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/google/uuid"
)

// MessageEnqueuer provides an interface implemented by QueueMessage usecase.
type MessageEnqueuer interface {
	Execute(ctx context.Context, senderID, receiverID uuid.UUID, content string) (entities.Message, error)
}

// MessageDequeuer provides an interface implemented by DequeueMessage usecase.
type MessageDequeuer interface {
	Execute(ctx context.Context, userID uuid.UUID) ([]entities.Message, error)
}

// MessageFacade uses facade pattern to wrap message usecases to allow for managing other things such as DB transactions if needed.
type MessageFacade struct {
	msgEnqueuer MessageEnqueuer
	msgDequeuer MessageDequeuer
}

// NewMessageFacade returns a new instance of MessageFacade.
func NewMessageFacade(msgEq MessageEnqueuer, msgDq MessageDequeuer) *MessageFacade {
	return &MessageFacade{
		msgEnqueuer: msgEq,
		msgDequeuer: msgDq,
	}
}

// EnqueueMessage wraps the QueueMessage usecase.
func (s *MessageFacade) EnqueueMessage(ctx context.Context, senderID, receiverID uuid.UUID, content string) (entities.Message, error) {
	return s.msgEnqueuer.Execute(ctx, senderID, receiverID, content)
}

// DequeueMessage wraps the DequeueMessage usecase.
func (s *MessageFacade) DequeueMessages(ctx context.Context, userID uuid.UUID) ([]entities.Message, error) {
	return s.msgDequeuer.Execute(ctx, userID)
}
