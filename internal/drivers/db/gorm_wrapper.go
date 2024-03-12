package db

import (
	"context"

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

func (s *GormWrapper) Create(ctx context.Context, value interface{}) *gorm.DB {
	return s.DB.WithContext(ctx).Create(value)
}

func (s *GormWrapper) Where(ctx context.Context, query interface{}, args ...interface{}) *gorm.DB {
	return s.DB.WithContext(ctx).Where(query, args...)
}
