package datastore

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/sirupsen/logrus"
)

// NotificationPostgresRepository implements NotificationRepository interface to provide database operation on notifications.
type NotificationPostgresRepository struct {
	Logger *logrus.Logger
	DB     Datastore
}

// NewNotificationPostgresRepository returns a new instance of NotificationPostgresRepository.
func NewNotificationPostgresRepository(l *logrus.Logger, db Datastore) *NotificationPostgresRepository {
	return &NotificationPostgresRepository{
		Logger: l,
		DB:     db,
	}
}

// Create creates a new record in the database.
func (s *NotificationPostgresRepository) Create(ctx context.Context, notification *entities.Notification) error {
	result := s.DB.WithContext(ctx).Create(notification)

	err := result.Error()
	if err != nil {
		s.Logger.WithError(err).Error("failed to create a notification in the database")
		return repositories.NewInternalError("failed to create a notification in the database")
	}

	return nil
}

// GetAll get all notifications from our database.
func (s *NotificationPostgresRepository) GetAll(ctx context.Context) ([]entities.Notification, error) {
	var notifications []entities.Notification
	result := s.DB.WithContext(ctx).FindAll(&notifications)

	err := result.Error()
	if err != nil {
		s.Logger.WithError(err).Error("failed to get all notifications in the database")
		return nil, repositories.NewInternalError("failed to get all notifications in the database")
	}

	return notifications, nil
}
