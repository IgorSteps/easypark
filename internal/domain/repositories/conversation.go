package repositories

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/google/uuid"
)

// ConversationRepository describes an interface for CRUD operations on
type ConversationRepository interface {
	// Create creates a conversation in the database.
	Create(ctx context.Context, conversation *entities.Conversation) (*entities.Conversation, error)
	// GetSingle gets a single conversation from the database using the id.
	GetSingle(ctx context.Context, id uuid.UUID) (entities.Conversation, error)
	// GetMany gets many conversation that match give query.
	GetMany(ctx context.Context, query map[string]interface{}) ([]entities.Conversation, error)
}
