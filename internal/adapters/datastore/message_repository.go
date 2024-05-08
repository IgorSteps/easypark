package datastore

import (
	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type MessagePostgresRepository struct {
	logger *logrus.Logger
	db     Datastore
}

func NewMessagePostgresRepository(l *logrus.Logger, db Datastore) *MessagePostgresRepository {
	return &MessagePostgresRepository{
		logger: l,
		db:     db,
	}
}

func (s *MessagePostgresRepository) Create(message *entities.Message) error {
	result := s.db.Create(message)

	err := result.Error()
	if err != nil {
		s.logger.WithError(err).Error("failed to insert message into the database")
		return repositories.NewInternalError("failed to insert message into the database")
	}

	return nil
}

func (s *MessagePostgresRepository) GetManyForUser(userID uuid.UUID) ([]entities.Message, error) {
	var messages []entities.Message

	result := s.db.Where("receiver_id = ?", userID).FindAll(&messages)
	err := result.Error()
	if err != nil {
		s.logger.WithError(err).WithField("user id", userID).Error("failed to query for many messages in the database")
		return []entities.Message{}, repositories.NewInternalError("failed to query for many messages in the database")
	}

	return messages, nil
}

func (s *MessagePostgresRepository) Delete(messages []entities.Message) error {
	if messages == nil {
		return nil
	}

	// Extract IDs from the messages slice
	var ids []uuid.UUID
	for _, message := range messages {
		ids = append(ids, message.ID)
	}

	result := s.db.Where("id IN ?", ids).Delete(&messages)

	err := result.Error()
	if err != nil {
		s.logger.WithError(err).Error("failed to delete messages")
		return repositories.NewInternalError("failed to delete messages")
	}
	return nil
}
