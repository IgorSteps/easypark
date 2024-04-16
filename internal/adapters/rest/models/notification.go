package models

import "github.com/google/uuid"

// CreateNotificationRequest represents an incoming HTTP request to create notification.
type CreateNotificationRequest struct {
	ParkingRequestID uuid.UUID
	ParkingSpaceID   uuid.UUID
	Location         string
	NotificationType int
}
