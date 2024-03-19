package usecasefacades

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
)

// DriverCreator provides an interface implemented by the RegisterDriver usecase.
type DriverCreator interface {
	Execute(ctx context.Context, user *entities.User) error
}

// UserAuthenticator provides an interface implemented by the AuthenticateUser usecase.
type UserAuthenticator interface {
	Execute(ctx context.Context, username, password string) (string, error)
}

// DriversGetter provides an interfaces implemented by the GetDrivers usecase.
type DriversGetter interface {
	Execute(ctx context.Context) ([]entities.User, error)
}

// UserFacade uses facade pattern to wrap user' usecases to allow for managing other things such as DB transactions if needed.
type UserFacade struct {
	driverCreator     DriverCreator
	userAuthenticator UserAuthenticator
	driversGetter     DriversGetter
}

// NewUserFacade creates new instance of UserFacade.
func NewUserFacade(uCreator DriverCreator, uAuthenticator UserAuthenticator, uGetter DriversGetter) *UserFacade {
	return &UserFacade{
		driverCreator:     uCreator,
		userAuthenticator: uAuthenticator,
		driversGetter:     uGetter,
	}
}

// CreateDriver wraps the RegisterDriver usecase.
func (s *UserFacade) CreateDriver(ctx context.Context, user *entities.User) error {
	return s.driverCreator.Execute(ctx, user)
}

// AuthoriseUser wraps the AuthenticateUser usecase.
func (s *UserFacade) AuthoriseUser(ctx context.Context, username, password string) (string, error) {
	return s.userAuthenticator.Execute(ctx, username, password)
}

// GetAllDriverUsers wraps the GetDrivers usecase.
func (s *UserFacade) GetAllDriverUsers(ctx context.Context) ([]entities.User, error) {
	return s.driversGetter.Execute(ctx)
}
