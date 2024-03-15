package usecases

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// RegisterUser provides business logic to register a user.
type RegisterUser struct {
	Logger         *logrus.Logger
	UserRepository repositories.UserRepository
}

// NewRegisterUser returns new RegisterUser usecase.
func NewRegisterUser(logger *logrus.Logger, repo repositories.UserRepository) *RegisterUser {
	return &RegisterUser{
		Logger:         logger,
		UserRepository: repo,
	}
}

// Execute runs the usecase business logic.
func (s *RegisterUser) Execute(ctx context.Context, user *entities.User) error {
	err := s.validate(ctx, user)
	if err != nil {
		return err
	}

	// TODO: Move to user entity
	user.ID = uuid.New()
	user.Role = entities.RoleDriver

	err = s.UserRepository.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

// Validate checks if the user already exists using their email or username.
func (s *RegisterUser) validate(ctx context.Context, user *entities.User) error {
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
