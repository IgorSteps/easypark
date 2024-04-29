package client

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/google/uuid"
)

type MessageFacade interface {
	EnqueueMessage(ctx context.Context, senderID, receiverID uuid.UUID, content string) (entities.Message, error)
	DequeueMessages(ctx context.Context, userID uuid.UUID) ([]entities.Message, error)
}
