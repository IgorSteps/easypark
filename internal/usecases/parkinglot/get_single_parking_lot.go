package usecases

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// GetSingleParkingLot provides business logic to get single parking lot.
type GetSingleParkingLot struct {
	Logger *logrus.Logger
	Repo   repositories.ParkingLotRepository
}

// NewGetSingleParkingLot returns a new instance of GetSingleParkingLot.
func NewGetSingleParkingLot(l *logrus.Logger, r repositories.ParkingLotRepository) *GetSingleParkingLot {
	return &GetSingleParkingLot{
		Logger: l,
		Repo:   r,
	}
}

// Execute runs the business logic to get single parking lot.
func (s *GetSingleParkingLot) Execute(ctx context.Context, lotID uuid.UUID) (*entities.ParkingLot, error) {
	lot, err := s.Repo.GetSingle(ctx, lotID)
	if err != nil {
		return nil, err
	}

	lot.OnGet()

	return lot, nil
}
