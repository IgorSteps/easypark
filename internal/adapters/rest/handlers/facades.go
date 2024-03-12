package handlers

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
)

// UserFacade is provides an interface implemented by usecasefacades.UserFacade.
type UserFacade interface {
	// CreateUser is implemented by secasefacades.UserFacade that wraps user creation usecase.
	CreateUser(ctx context.Context, driver *entities.User) error
}
