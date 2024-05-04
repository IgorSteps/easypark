package usecases

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/sirupsen/logrus"
)

// AuthenticateUser provides business logic to authenticate users.
type AuthenticateUser struct {
	logger         *logrus.Logger
	userRepository repositories.UserRepository
	tokenService   repositories.TokenRepository
}

// NewAuthenticateUser returns new instance of AuthenticateUser.
func NewAuthenticateUser(l *logrus.Logger, repo repositories.UserRepository, tService repositories.TokenRepository) *AuthenticateUser {
	return &AuthenticateUser{
		logger:         l,
		userRepository: repo,
		tokenService:   tService,
	}
}

// Execute runs business logic to authenticate users. Returns auth token.
func (s *AuthenticateUser) Execute(ctx context.Context, username, password string) (*entities.User, string, error) {
	user, err := s.userRepository.GetDriverByUsername(ctx, username)
	if err != nil {
		return nil, "", err
	}

	// TODO: Use hashing.
	if user.Password != password {
		s.logger.WithField("username", username).Warn("provided invalid credentials")
		return nil, "", repositories.NewInvalidInputError("invalid password")
	}

	token, err := s.tokenService.GenerateToken(user)
	if err != nil {
		s.logger.WithError(err).Error("failed to generate auth token")
		return nil, "", repositories.NewInternalError("failed to generate auth token")
	}

	return user, token, nil
}
