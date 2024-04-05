package datastore_test

import (
	"context"
	"errors"
	"testing"

	"github.com/IgorSteps/easypark/internal/adapters/datastore"
	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	mocks "github.com/IgorSteps/easypark/mocks/adapters/datastore"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestParkingLotPostgresRepository_CreateParkingLot_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockDB := &mocks.Datastore{}
	repo := datastore.NewParkingParkingLotPostgresRepository(testLogger, mockDB)

	testCtx := context.Background()
	parkingLot := &entities.ParkingLot{
		ID: uuid.New(),
	}

	mockDB.EXPECT().WithContext(testCtx).Return(mockDB).Once()
	mockDB.EXPECT().Create(parkingLot).Return(mockDB).Once()
	mockDB.EXPECT().Error().Return(nil).Once()

	// --------
	// ACT
	// --------
	err := repo.CreateParkingLot(testCtx, parkingLot)

	// --------
	// ASSERT
	// --------
	assert.Nil(t, err, "error must be nil")
	mockDB.AssertExpectations(t)
}

func TestParkingLotPostgresRepository_CreateParkingLot_UnhappyPath_InternalError(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockDB := &mocks.Datastore{}
	repo := datastore.NewParkingParkingLotPostgresRepository(testLogger, mockDB)

	testCtx := context.Background()
	parkingLot := &entities.ParkingLot{
		ID: uuid.New(),
	}
	testError := errors.New("boom")

	mockDB.EXPECT().WithContext(testCtx).Return(mockDB).Once()
	mockDB.EXPECT().Create(parkingLot).Return(mockDB).Once()
	mockDB.EXPECT().Error().Return(testError).Once()

	// --------
	// ACT
	// --------
	err := repo.CreateParkingLot(testCtx, parkingLot)

	// --------
	// ASSERT
	// --------
	assert.IsType(t, &repositories.InternalError{}, err, "wrong error type")
	assert.EqualError(t, err, "Internal error: failed to insert parking lot into the database", "errors are not equal")
	mockDB.AssertExpectations(t)
}

func TestParkingLotPostgresRepository_CreateParkingLot_UnhappyPath_UniqueKeyViolation(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockDB := &mocks.Datastore{}
	repo := datastore.NewParkingParkingLotPostgresRepository(testLogger, mockDB)

	testCtx := context.Background()
	parkingLot := &entities.ParkingLot{
		ID:   uuid.New(),
		Name: "sci",
	}
	testError := &pgconn.PgError{
		Code: datastore.ErrCodeUniqueViolation,
	}

	mockDB.EXPECT().WithContext(testCtx).Return(mockDB).Once()
	mockDB.EXPECT().Create(parkingLot).Return(mockDB).Once()
	mockDB.EXPECT().Error().Return(testError).Once()

	// --------
	// ACT
	// --------
	err := repo.CreateParkingLot(testCtx, parkingLot)

	// --------
	// ASSERT
	// --------
	assert.IsType(t, &repositories.ResourceAlreadyExistsError{}, err, "wrong error type")
	assert.EqualError(t, err, "Resource 'sci' already exists", "errors are not equal")
	mockDB.AssertExpectations(t)
}
