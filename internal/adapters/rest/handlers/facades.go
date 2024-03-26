package handlers

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/google/uuid"
)

// UserFacade is provides an interface implemented by usecasefacades.UserFacade.
type UserFacade interface {
	// CreateDriver is implemented by usecasefacades.UserFacade that wraps driver user creation usecase.
	CreateDriver(ctx context.Context, driver *entities.User) error

	// AuthoriseUser is implemented by usecasefacades.UserFacade that wraps user login usecase.
	AuthoriseUser(ctx context.Context, username, password string) (string, error)

	// GetAllDriverUsers is implemented by the usecasefacades.Userfacade that wraps getting all driver users usecase.
	GetAllDriverUsers(ctx context.Context) ([]entities.User, error)

	// BanDriver is implemented by the usecasefacades.Userfacade that wraps banning a driver usecase.
	BanDriver(ctx context.Context, id uuid.UUID) error
}
