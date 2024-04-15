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
	requestRepo      repositories.ParkingRequestRepository
	alertCreator     repositories.AlertCreator
}

// NewCreateNotification returns a new instance of CreateNotification.
func NewCreateNotification(
	l *logrus.Logger,
	notifRepo repositories.NotificationRepository,
	spaceRepo repositories.ParkingSpaceRepository,
	requestRepo repositories.ParkingRequestRepository,
	alertCreator repositories.AlertCreator,
) *CreateNotification {
	return &CreateNotification{
		logger:           l,
		notificationRepo: notifRepo,
		spaceRepo:        spaceRepo,
		requestRepo:      requestRepo,
		alertCreator:     alertCreator,
	}
}

// Execute runs the business logic to create notifications.
func (s *CreateNotification) Execute(
	ctx context.Context,
	driverID,
	parkingRequestID,
	spaceID uuid.UUID,
	location string,
	notificationType int,
) (entities.Notification, error) {
	// Parse notification type.
	domainNotificationType, err := parseNotificationType(notificationType)
	if err != nil {
		s.logger.WithError(err).Error("invalid notification type")
		return entities.Notification{}, err
	}

	// Get parking space.
	parkingSpace, err := s.spaceRepo.GetSingle(ctx, spaceID)
	if err != nil {
		return entities.Notification{}, err
	}

	// Get parking request.
	parkingRequest, err := s.requestRepo.GetSingle(ctx, parkingRequestID)
	if err != nil {
		return entities.Notification{}, err
	}

	// Check for mismatch between driver IDs.
	if parkingRequest.UserID != driverID {
		s.logger.Error("userID in the parking request doesn't match userID of whoever created the notification")
		return entities.Notification{}, repositories.NewInvalidInputError("userID in the parking request doesn't match userID of whoever created the notification")
	}

	// Check for mismatch between parking space IDs.
	if *parkingRequest.ParkingSpaceID != spaceID {
		s.logger.Error("parkingSpaceID in the parking request doesn't match parkingSpaceID in the notification")
		return entities.Notification{}, repositories.NewInvalidInputError("parkingSpaceID in the parking request doesn't match parkingSpaceID in the notification")
	}

	// Create notification.
	notification := entities.Notification{}
	notification.OnCreate(driverID, spaceID, location, domainNotificationType)
	err = s.notificationRepo.Create(ctx, &notification)
	if err != nil {
		return entities.Notification{}, err
	}

	// Check notification type and update parking space and request accordingly.
	if domainNotificationType == entities.ArrivalNotification {
		// Check if the parking space and request statuses are incorrect.
		if parkingSpace.Status != entities.ParkingSpaceStatusAvailable {
			s.logger.WithField("parking space id", parkingSpace.ID).Error("parking space that isn't available received arrival notification")
			return entities.Notification{}, repositories.NewInvalidInputError("parking space that isn't available received arrival notification")
		}
		if parkingRequest.Status != entities.RequestStatusApproved {
			s.logger.WithField("parking request id", parkingRequest.ID).Error("parking request that isn't approved received arrival notification")
			return entities.Notification{}, repositories.NewInvalidInputError("parking request that isn't approved received arrival notification")
		}

		// Check for location mismatch.
		if location != parkingSpace.Name {
			alert, err := s.alertCreator.Execute(ctx, entities.LocationMismatch, "driver arrived at wrong parking space", driverID, spaceID)
			if err != nil {
				s.logger.WithError(err).Error("failed to create location mismatch alert")
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

			s.logger.Info("location mismatch on arrival")
			return notification, nil
		}

		// Update parking space and request.
		parkingSpace.OnArrival()
		parkingRequest.OnArrivalNotification()
	} else {
		// Check if the parking space and request statuses are incorrect.
		if parkingSpace.Status != entities.ParkingSpaceStatusOccupied {
			s.logger.WithField("parking space id", parkingSpace.ID).Error("parking space that isn't occupied cannot receive departure notification")
			return entities.Notification{}, repositories.NewInvalidInputError("parking space that isn't occupied cannot receive departure notification")
		}
		if parkingRequest.Status != entities.RequestStatusActive {
			s.logger.WithField("parking request id", parkingRequest.ID).Error("parking request that isn't active cannot receive departure notification")
			return entities.Notification{}, repositories.NewInvalidInputError("parking request that isn't active cannot receive departure notification")
		}

		// Update parking space and request.
		parkingSpace.OnDeparture()
		parkingRequest.OnDepartureNotification()
	}

	// Save updated parking request.
	err = s.requestRepo.Save(ctx, &parkingRequest)
	if err != nil {
		return entities.Notification{}, err
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
