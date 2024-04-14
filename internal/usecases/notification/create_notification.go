package usecases

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// CreateNotification provides business logic to create notifications.
type CreateNotification struct {
	logger           *logrus.Logger
	notificationRepo repositories.NotificationRepository
	spaceRepo        repositories.ParkingSpaceRepository
	alertCreator     repositories.AlertCreator
}

// NewCreateNotification returns a new instance of CreateNotification.
func NewCreateNotification(
	l *logrus.Logger,
	notifRepo repositories.NotificationRepository,
	spaceRepo repositories.ParkingSpaceRepository,
	alertCreator repositories.AlertCreator,
) *CreateNotification {
	return &CreateNotification{
		logger:           l,
		notificationRepo: notifRepo,
		spaceRepo:        spaceRepo,
		alertCreator:     alertCreator,
	}
}

// Execute runs the business logic to create notifications.
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
		// Check for if the parking space status is not available (shouldn't be the case).
		if parkingSpace.Status != entities.StatusAvailable {
			s.logger.WithField("parking space id", parkingSpace.ID).Warn("parking space isn't available but received arrival notification")
		}

		// Check for location mismatch.
		if location != parkingSpace.Name {
			alert, err := s.alertCreator.Execute(ctx, entities.LocationMismatch, "driver arrived at wrong parking space", driverID, spaceID)
			if err != nil {
				s.logger.Error("failed to create location mismatch alert")
				return entities.Notification{}, err
			}

			// Let's debug log the alert.
			s.logger.WithFields(logrus.Fields{
				"id":             alert.ID,
				"type":           alert.Type,
				"msg":            alert.Message,
				"driverID":       alert.UserID,
				"parkingSpaceID": alert.ParkingSpaceID,
			}).Debug("created location mismatch alert")

			s.logger.Info("location mismatch, not updating parking space status to occupied")
			return notification, nil
		}

		// Update parking space status.
		parkingSpace.Status = entities.StatusOccupied
	} else {
		// Check for if the parking space status is not occupied (shouldn't be the case).
		if parkingSpace.Status != entities.StatusOccupied {
			s.logger.WithField("parking space id", parkingSpace.ID).Warn("parking space isn't occupied but received departure notification")
		}

		// Update parking space status.
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
