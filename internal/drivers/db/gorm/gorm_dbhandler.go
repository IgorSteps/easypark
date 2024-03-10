package gorm

import (
	"context"

	"gorm.io/gorm"
)

// GormDbHandler is a wrapper around GORM.
type GormDBHandler struct {
	DB *gorm.DB
}

func (g *GormDBHandler) Create(ctx context.Context, value interface{}) *gorm.DB {
	return g.DB.WithContext(ctx).Create(value)
}
