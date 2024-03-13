package handlers

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
)

// UserFacade is provides an interface implemented by usecasefacades.UserFacade.
type UserFacade interface {
	// CreateUser is implemented by usecasefacades.UserFacade that wraps user creation usecase.
	CreateUser(ctx context.Context, driver *entities.User) error

	// AuthoriseUser is implemented by usecasefacades.UserFacade that wraps user login usecase.
	AuthoriseUser(ctx context.Context, username, password string) (string, error)
}
