package usecases

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/sirupsen/logrus"
)

// GetDrivers provides business logic to get all driver users.
//
// NOTE: Can and should be extended to include a filter for a real world scenario.
type GetDrivers struct {
	logger *logrus.Logger
	repo   repositories.UserRepository
}

// NewGetDrivers returns a new instance of GetDrivers.
func NewGetDrivers(l *logrus.Logger, r repositories.UserRepository) *GetDrivers {
	return &GetDrivers{
		logger: l,
		repo:   r,
	}
}

// Execute runs the business logic to get all driver users.
func (s *GetDrivers) Execute(ctx context.Context) ([]entities.User, error) {
	users, err := s.repo.GetAllDriverUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}
