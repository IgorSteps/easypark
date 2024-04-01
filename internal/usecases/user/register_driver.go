package usecases

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/sirupsen/logrus"
)

// RegisterDriver provides business logic to register a driver user.
type RegisterDriver struct {
	Logger         *logrus.Logger
	UserRepository repositories.UserRepository
}

// NewRegisterDriver returns new RegisterDriver usecase.
func NewRegisterDriver(logger *logrus.Logger, repo repositories.UserRepository) *RegisterDriver {
	return &RegisterDriver{
		Logger:         logger,
		UserRepository: repo,
	}
}

// Execute runs the usecase business logic.
func (s *RegisterDriver) Execute(ctx context.Context, user *entities.User) error {
	err := s.validate(ctx, user)
	if err != nil {
		return err
	}

	user.SetOnCreate()
	err = s.UserRepository.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

// Validate checks if the user already exists using their email or username.
func (s *RegisterDriver) validate(ctx context.Context, user *entities.User) error {
	doesExist, err := s.UserRepository.CheckUserExists(ctx, user.Email, user.Username)
	if err != nil {
		return err
	}

	if doesExist {
		s.Logger.WithFields(logrus.Fields{
			"username": user.Username,
			"email":    user.Email,
		}).Warn("user already exists")

		return repositories.NewUserAlreadyExistsError(user.Username, user.Email)
	}

	return nil
}
