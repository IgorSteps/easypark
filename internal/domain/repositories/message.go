package repositories

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
)

// MessageRepository provides an interface for CRUD operations on Messages.
type MessageRepository interface {
	// Create creates a message in our DB.
	Create(ctx context.Context, msg *entities.Message) (*entities.Message, error)
	// GetMany gets many messages.
	GetMany(ctx context.Context, query map[string]interface{}) ([]entities.Message, error)
}
