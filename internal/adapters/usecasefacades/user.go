package usecasefacades

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
)

// UserCreator provides an interface implemented by the RegisterUser usecase.
type UserCreator interface {
	Execute(ctx context.Context, user *entities.User) error
}

// UserFacade uses facade pattern to wrap user' usecases in DB transactions or other things.
type UserFacade struct {
	userCreator UserCreator
}

// NewUserFacade creates new instance of UserFacade.
func NewUserFacade(uCreator UserCreator) *UserFacade {
	return &UserFacade{
		userCreator: uCreator,
	}
}

// CreateUser wraps the RegisterUser usecase.
func (s *UserFacade) CreateUser(ctx context.Context, user *entities.User) error {
	return s.userCreator.Execute(ctx, user)
}
