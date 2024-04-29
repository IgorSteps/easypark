package models

import "github.com/google/uuid"

type Message struct {
	ReceiverID uuid.UUID `json:"receiverID"`
	Content    string    `json:"content"`
}
