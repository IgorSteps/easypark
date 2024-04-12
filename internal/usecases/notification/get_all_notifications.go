package usecases

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/sirupsen/logrus"
)

// GetAllNotifications provides business logic to get all notifications.
type GetAllNotifications struct {
	logger           *logrus.Logger
	notificationRepo repositories.NotificationRepository
}

// NewGetAllNotifications returns new instance of GetAllNotifications.
func NewGetAllNotifications(l *logrus.Logger, r repositories.NotificationRepository) *GetAllNotifications {
	return &GetAllNotifications{
		logger:           l,
		notificationRepo: r,
	}
}

// Execute runs the business logic to get all notifications.
func (s *GetAllNotifications) Execute(ctx context.Context) ([]entities.Notification, error) {
	return s.notificationRepo.GetAll(ctx)
}
