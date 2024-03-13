package usecasefacades

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
)

// UserCreator provides an interface implemented by the RegisterUser usecase.
type UserCreator interface {
	Execute(ctx context.Context, user *entities.User) error
}

// UserAuthenticator provides an interface implemented by the AuthenticateUser usecase.
type UserAuthenticator interface {
	Execute(ctx context.Context, username, password string) (string, error)
}

// UserFacade uses facade pattern to wrap user' usecases while also managing other things such as DB transactions if neeeded.
type UserFacade struct {
	userCreator       UserCreator
	userAuthenticator UserAuthenticator
}

// NewUserFacade creates new instance of UserFacade.
func NewUserFacade(uCreator UserCreator, uAuthenticator UserAuthenticator) *UserFacade {
	return &UserFacade{
		userCreator:       uCreator,
		userAuthenticator: uAuthenticator,
	}
}

// CreateUser wraps the RegisterUser usecase.
func (s *UserFacade) CreateUser(ctx context.Context, user *entities.User) error {
	return s.userCreator.Execute(ctx, user)
}

// AuthoriseUser wraps the AuthenticateUser usecase.
func (s *UserFacade) AuthoriseUser(ctx context.Context, username, password string) (string, error) {
	return s.userAuthenticator.Execute(ctx, username, password)
}
