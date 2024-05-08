package entities

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID         uuid.UUID
	SenderID   uuid.UUID
	ReceiverID uuid.UUID
	Content    string
	Delivered  bool
	Timestamp  time.Time
}
