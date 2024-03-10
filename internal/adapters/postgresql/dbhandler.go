package postgresql

import (
	"context"

	"gorm.io/gorm"
)

// DBHandler defines the interface for database operations used by the repositories.
type DBHandler interface {
	Create(ctx context.Context, value interface{}) *gorm.DB
}
