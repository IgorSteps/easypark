package repositories

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
)

// NotificationRepository provides an interfaces for CRUD operations on Notifications.
type NotificationRepository interface {
	// Create creates a notification in our DB.
	Create(ctx context.Context, notification *entities.Notification) error

	// GetAll gets all notifications for our DB.
	GetAll(ctx context.Context) ([]entities.Notification, error)
}
