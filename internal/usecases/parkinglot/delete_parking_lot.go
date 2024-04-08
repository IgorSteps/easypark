package usecases

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// DeteleParkingLot provides business logic to delete a parking lot.
type DeteleParkingLot struct {
	logger *logrus.Logger
	repo   repositories.ParkingLotRepository
}

// NewDeleteParkingLot returns a new instance of DeteleParkingLot.
func NewDeleteParkingLot(l *logrus.Logger, r repositories.ParkingLotRepository) *DeteleParkingLot {
	return &DeteleParkingLot{
		logger: l,
		repo:   r,
	}
}

// Execute runs the business logic.
func (s *DeteleParkingLot) Execute(ctx context.Context, id uuid.UUID) error {
	err := s.repo.DeleteParkingLot(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
