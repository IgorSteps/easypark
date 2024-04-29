package repositories

import (
	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/google/uuid"
)

type MessageRepository interface {
	Create(message *entities.Message) error
	GetManyForUser(userID uuid.UUID) ([]entities.Message, error)
	Delete(messages []entities.Message) error
}
