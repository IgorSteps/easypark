package repositories

import (
	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/google/uuid"
)

type MessageRepository interface {
	Send(msg entities.Message, userID uuid.UUID) error
}
