package datastore

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/sirupsen/logrus"
)

// Error code to signify that a unique contraint has been violated.
const ErrCodeUniqueViolation = "23505"

// ParkingLotPostgresRepository implements PakringLotRepository interface to provide database operation on parking lots.
type ParkingLotPostgresRepository struct {
	Logger *logrus.Logger
	DB     Datastore
}

// NewParkingParkingLotPostgresRepository returns a new instance of ParkingLotPostgresRepository.
func NewParkingParkingLotPostgresRepository(l *logrus.Logger, db Datastore) *ParkingLotPostgresRepository {
	return &ParkingLotPostgresRepository{
		Logger: l,
		DB:     db,
	}
}

// CreateParkingLot creates a parking lot in the database.
func (s *ParkingLotPostgresRepository) CreateParkingLot(ctx context.Context, parkingLot *entities.ParkingLot) error {
	result := s.DB.WithContext(ctx).Create(parkingLot)
	err := result.Error()
	if err != nil {
		if pgError, ok := err.(*pgconn.PgError); ok {
			if pgError.Code == ErrCodeUniqueViolation {
				s.Logger.WithField("name", parkingLot.Name).WithError(err).Warn("parking lot with this name already exists")
				return repositories.NewResourceAlreadyExistsError(parkingLot.Name)
			}
		}

		s.Logger.WithError(err).Error("failed to insert parking lot into the database")
		return repositories.NewInternalError("failed to insert parking lot into the database")
	}

	return nil
}
