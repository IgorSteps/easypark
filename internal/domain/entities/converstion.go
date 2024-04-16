package entities

import "github.com/google/uuid"

// Conversation represents a conversation between users.
type Conversation struct {
	ID       uuid.UUID
	Messages []Message
}
