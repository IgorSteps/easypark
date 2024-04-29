package repositories

import "github.com/IgorSteps/easypark/internal/domain/entities"

type MessageRepository interface {
	Create(message entities.Message) error
	GetManyForUser(userID string) ([]entities.Message, error)
}
