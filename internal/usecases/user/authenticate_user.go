package usecases

import (
	"context"

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
func (s *AuthenticateUser) Execute(ctx context.Context, username, password string) (string, error) {
	user, err := s.userRepository.GetDriverByUsername(ctx, username)
	if err != nil {
		return "", err
	}

	// Not how it should be done in real world.
	if user.Password != password {
		s.logger.WithField("username", username).Warn("provided invalid credentials")
		return "", repositories.NewInvalidInputError("invalid password")
	}

	token, err := s.tokenService.GenerateToken(user)
	if err != nil {
		s.logger.WithError(err).Error("failed to generate auth token")
		return "", repositories.NewInternalError("failed to generate auth token")
	}

	return token, nil
}
