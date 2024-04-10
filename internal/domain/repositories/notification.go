package repositories

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
)

// NotificationRepository provides an interfaces for CRUD opertions on Notifications.
type NotificationRepository interface {
	// Create creates a notificaioin in our DB.
	Create(ctx context.Context, notification *entities.Notification) error
}
