package usecases

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type CreateNotification struct {
	logger           *logrus.Logger
	notificationRepo repositories.NotificationRepository
}

func NewCreateNotification(l *logrus.Logger, r repositories.NotificationRepository) *CreateNotification {
	return &CreateNotification{
		logger:           l,
		notificationRepo: r,
	}
}

func (s *CreateNotification) Execute(ctx context.Context, driverID uuid.UUID, spaceID uuid.UUID, location string, notificationType int) (entities.Notification, error) {
	domainNotificationType, err := parseNotificationType(notificationType)
	if err != nil {
		s.logger.WithError(err).Error("invalid notification type")
		return entities.Notification{}, err
	}

	notification := entities.Notification{}
	notification.OnCreate(driverID, spaceID, location, domainNotificationType)

	err = s.notificationRepo.Create(ctx, &notification)
	if err != nil {
		return entities.Notification{}, err
	}

	return notification, nil
}

// parseNotificationType converts an integer to a NotificationType.
func parseNotificationType(value int) (entities.NotificationType, error) {
	switch value {
	case int(entities.ArrivalNotification):
		return entities.ArrivalNotification, nil
	case int(entities.DepartureNotification):
		return entities.DepartureNotification, nil
	default:
		return 0, repositories.NewInvalidInputError("invalid notification type")
	}
}
