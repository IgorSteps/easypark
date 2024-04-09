package usecasefacades

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/google/uuid"
)

// NotificationCreator provides an interface implemented by CreateNotification usecase.
type NotificationCreator interface {
	Execute(ctx context.Context, driverID, spaceID uuid.UUID, location string, notificationType int) (entities.Notification, error)
}

// NotificationFacade uses facade pattern to wrap notification usecases to allow for managing other things such as DB transactions if needed.
type NotificationFacade struct {
	notificationCreator NotificationCreator
}

// NewNotificationFacade returns new instance of NotificationFacade.
func NewNotificationFacade(creator NotificationCreator) *NotificationFacade {
	return &NotificationFacade{
		notificationCreator: creator,
	}
}

// CreateNotification wraps the CreateNotification usecase.
func (s *NotificationFacade) CreateNotification(ctx context.Context, driverID, spaceID uuid.UUID, location string, notificationType int) (entities.Notification, error) {
	return s.notificationCreator.Execute(ctx, driverID, spaceID, location, notificationType)
}
