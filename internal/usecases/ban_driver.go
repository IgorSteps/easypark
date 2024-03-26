package usecases

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// BanDriver provides business logic to ban a driver.
type BanDriver struct {
	logger *logrus.Logger
	repo   repositories.UserRepository
}

// NewBanDriver returns a new instance of BanDriver.
func NewBanDriver(l *logrus.Logger, r repositories.UserRepository) *BanDriver {
	return &BanDriver{
		logger: l,
		repo:   r,
	}
}

// Execute runs the business logic to ban a driver.
func (s *BanDriver) Execute(ctx context.Context, id uuid.UUID) error {
	var userToUpdate entities.User
	err := s.repo.GetDriverByID(ctx, id, &userToUpdate)
	if err != nil {
		return err
	}

	userToUpdate.Ban()

	err = s.repo.Save(ctx, &userToUpdate)
	if err != nil {
		return err
	}

	return nil
}
