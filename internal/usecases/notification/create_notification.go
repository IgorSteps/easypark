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
	spaceRepo        repositories.ParkingSpaceRepository
}

func NewCreateNotification(l *logrus.Logger, notifRepo repositories.NotificationRepository, spaceRepo repositories.ParkingSpaceRepository) *CreateNotification {
	return &CreateNotification{
		logger:           l,
		notificationRepo: notifRepo,
		spaceRepo:        spaceRepo,
	}
}

func (s *CreateNotification) Execute(ctx context.Context, driverID uuid.UUID, spaceID uuid.UUID, location string, notificationType int) (entities.Notification, error) {
	domainNotificationType, err := parseNotificationType(notificationType)
	if err != nil {
		s.logger.WithError(err).Error("invalid notification type")
		return entities.Notification{}, err
	}

	parkingSpace, err := s.spaceRepo.GetParkingSpaceByID(ctx, spaceID)
	if err != nil {
		return entities.Notification{}, err
	}

	notification := entities.Notification{}
	notification.OnCreate(driverID, spaceID, location, domainNotificationType)

	err = s.notificationRepo.Create(ctx, &notification)
	if err != nil {
		return entities.Notification{}, err
	}

	// Check notification type and update parking space status accordingly.
	if domainNotificationType == entities.ArrivalNotification {
		// Arrival.
		if parkingSpace.Status != entities.StatusAvailable {
			// Let's warn if the driver notifies about arrival at the space, which status is not available.
			s.logger.WithField("parking space id", parkingSpace.ID).Warn("parking space isn't available but received arrival notification, continuing...")
		}

		parkingSpace.Status = entities.StatusOccupied
	} else {
		// Departure.
		if parkingSpace.Status != entities.StatusOccupied {
			// Let's warn if the driver notifies about departure from a space, which status is not occupied.
			s.logger.WithField("parking space id", parkingSpace.ID).Warn("parking space isn't occupied but received departure notification, continuing...")
		}

		parkingSpace.Status = entities.StatusAvailable
	}

	// Save updated parking space.
	err = s.spaceRepo.Save(ctx, &parkingSpace)
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
