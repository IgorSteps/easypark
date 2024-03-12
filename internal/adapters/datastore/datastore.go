package datastore

import (
	"context"

	"gorm.io/gorm"
)

// Datastore defines the interface for database operations using GORM.
type Datastore interface {
	Create(ctx context.Context, value interface{}) *gorm.DB
	Where(ctx context.Context, query interface{}, args ...interface{}) *gorm.DB
}
