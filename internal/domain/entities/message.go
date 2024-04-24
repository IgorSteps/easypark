package entities

import "github.com/google/uuid"

// Message represents a message.
type Message struct {
	Content     []byte
	SenderID    uuid.UUID
	RecipientID uuid.UUID
}
