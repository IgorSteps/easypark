package entities

import "github.com/google/uuid"

// Message represents a message.
type Message struct {
	ID             string
	Content        string
	ConversationID uuid.UUID
	SenderID       uuid.UUID
	RecipientID    uuid.UUID
	Timestamp      int
}
