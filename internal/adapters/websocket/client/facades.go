package client

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/google/uuid"
)

// MessageFacade provides an interface implemented by usecasefacades.MessageFacade.
type MessageFacade interface {
	// EnqueueMessage is implemented by usecasefacades.MessageFacade that wraps enqueue message usecase.
	EnqueueMessage(ctx context.Context, senderID, receiverID uuid.UUID, content string) (entities.Message, error)
	// DequeueMessages is implemented by usecasefacades.MessageFacade that wraps dequeue messages usecase.
	DequeueMessages(ctx context.Context, userID uuid.UUID) ([]entities.Message, error)
}
