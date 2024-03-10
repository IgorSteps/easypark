package usecases

import (
	"context"
	"errors"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/sirupsen/logrus"
)

type RegisterUser struct {
	Logger   *logrus.Logger
	UserRepo repositories.UserRepository
}

// NewRegisterUser returns new RegisterUser usecase.
func NewRegisterUser(logger *logrus.Logger, repo repositories.UserRepository) *RegisterUser {
	return &RegisterUser{
		Logger:   logger,
		UserRepo: repo,
	}
}

// Execute runs the usecase business logic.
func (s *RegisterUser) Execute(ctx context.Context, user entities.User) error {
	err := s.validate(ctx, user)
	if err != nil {
		return err
	}

	_, err = s.UserRepo.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

// Validate checks if the user already exists using their email.
func (s *RegisterUser) validate(ctx context.Context, user entities.User) error {
	doesExist, err := s.UserRepo.CheckUserExistsByEmail(ctx, user.Email)
	if err != nil {
		s.Logger.WithFields(logrus.Fields{
			"email": user.Email,
			"error": err.Error(),
		}).Error("failed to get user by email")

		return err
	}

	if doesExist {
		s.Logger.WithFields(logrus.Fields{
			"email": user.Email,
		}).Warn("user already exists")

		return errors.New("user already exists")
	}

	return nil
}
