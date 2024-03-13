package usecases

import (
	"context"
	"time"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/sirupsen/logrus"
)

// TokenService provides an interface to generate auth tokens for users.
type TokenService interface {
	// GenereateToken creates an auth token for a given user.
	GenerateToken(user *entities.User, expiresAt int64) (string, error)
}

// AuthenticateUser provides business logic to authenticate users.
type AuthenticateUser struct {
	logger         *logrus.Logger
	userRepository repositories.UserRepository
	tokenService   TokenService
}

// NewAuthenticateUser returns new instance of AuthenticateUser.
func NewAuthenticateUser(l *logrus.Logger, repo repositories.UserRepository, tService TokenService) *AuthenticateUser {
	return &AuthenticateUser{
		logger:         l,
		userRepository: repo,
		tokenService:   tService,
	}
}

// Execute runs business logic to authenticate users. Returns auth token.
func (s *AuthenticateUser) Execute(ctx context.Context, username, password string) (string, error) {
	user, err := s.userRepository.FindByUsername(ctx, username)
	if err != nil {
		return "", err
	}

	// Not how it should be done in real world.
	if user.Password != password {
		s.logger.WithField("username", username).Warn("provided invalid credentials")
		return "", repositories.NewInvalidCredentialsError()
	}

	expiresAt := time.Now().Add(time.Hour * 1).Unix() // set for 1 hour // move into config?
	token, err := s.tokenService.GenerateToken(user, expiresAt)
	if err != nil {
		s.logger.WithError(err).Error("failed to generate auth token")
		return "", repositories.NewInternalError("failed to generate auth token")
	}

	return token, nil
}
