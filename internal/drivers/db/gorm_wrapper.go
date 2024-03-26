package db

import (
	"context"

	"github.com/IgorSteps/easypark/internal/adapters/datastore"
	"gorm.io/gorm"
)

// GormWrapper wraps a gorm.DB connection to implement the datastore.Datastore interface.
type GormWrapper struct {
	DB *gorm.DB
}

// NewGormWrapper creates a new GormWrapper instance.
func NewGormWrapper(db *gorm.DB) *GormWrapper {
	return &GormWrapper{
		DB: db,
	}
}

// WithContext sets the context for the current operation chain.
func (g *GormWrapper) WithContext(ctx context.Context) datastore.Datastore {
	return &GormWrapper{DB: g.DB.WithContext(ctx)}
}

func (s *GormWrapper) Create(value interface{}) datastore.Datastore {
	return &GormWrapper{DB: s.DB.Create(value)}
}

func (s *GormWrapper) Where(query interface{}, args ...interface{}) datastore.Datastore {
	return &GormWrapper{DB: s.DB.Where(query, args...)}
}

// First gets the first record matching the query.
func (s *GormWrapper) First(value interface{}, args ...interface{}) datastore.Datastore {
	return &GormWrapper{DB: s.DB.First(value, args...)}
}

// FindAll gets all records.
func (s *GormWrapper) FindAll(value interface{}) datastore.Datastore {
	return &GormWrapper{DB: s.DB.Find(value)}
}

// Save saves the updated record.
func (s *GormWrapper) Save(value interface{}) datastore.Datastore {
	return &GormWrapper{DB: s.DB.Save(value)}
}

// Error returns any errors encountered.
func (s *GormWrapper) Error() error {
	return s.DB.Error
}
