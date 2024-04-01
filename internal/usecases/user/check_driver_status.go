package usecases

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// CheckDriverStatus provides business logic to check driver's status.
type CheckDriverStatus struct {
	logger *logrus.Logger
	repo   repositories.UserRepository
}

// NewCheckDriverStatus provides new instance of CheckDriverStatus.
func NewCheckDriverStatus(l *logrus.Logger, r repositories.UserRepository) *CheckDriverStatus {
	return &CheckDriverStatus{
		logger: l,
		repo:   r,
	}
}

// Execute runs the business logic to check if the driver is banned.
func (s *CheckDriverStatus) Execute(ctx context.Context, id uuid.UUID) (bool, error) {
	var driver entities.User

	err := s.repo.GetDriverByID(ctx, id, &driver)
	if err != nil {
		return true, err
	}

	if driver.Status == entities.StatusBanned {
		return true, nil
	}

	return false, nil
}
