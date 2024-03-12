package db

import (
	"context"

	"github.com/IgorSteps/easypark/internal/adapters/datastore"
	"gorm.io/gorm"
)

type GormWrapper struct {
	DB *gorm.DB
}

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

func (s *GormWrapper) First(value interface{}) datastore.Datastore {
	return &GormWrapper{DB: s.DB.First(value)}
}

func (s *GormWrapper) Error() error {
	return s.DB.Error
}
