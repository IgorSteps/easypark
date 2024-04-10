package models

import "github.com/google/uuid"

// CreateNotificationRequest represents an incoming HTTP request to create notificaiton.
type CreateNotificationRequest struct {
	ParkingSpaceID   uuid.UUID
	Location         string
	NotificationType int
}
