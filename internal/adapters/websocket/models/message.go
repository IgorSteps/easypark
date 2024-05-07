package models

import (
	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/google/uuid"
)

type Message struct {
	SenderID   uuid.UUID `json:"senderID"`
	ReceiverID uuid.UUID `json:"receiverID"`
	Content    string    `json:"content"`
}

func FromDomain(domainMessage entities.Message) *Message {
	return &Message{
		SenderID:   domainMessage.SenderID,
		ReceiverID: domainMessage.ReceiverID,
		Content:    domainMessage.Content,
	}
}
