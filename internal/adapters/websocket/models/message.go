package models

import "github.com/google/uuid"

type ChatMessage struct {
	SenderID   uuid.UUID `json:"senderID"`
	ReceiverID uuid.UUID `json:"receiverID"`
	Content    string    `json:"content"`
}
