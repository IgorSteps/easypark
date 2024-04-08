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
	"github.com/stretchr/testify/mock"
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

func TestParkingLotPostgresRepository_GetAllParkingLots_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockDatastore := &mocks.Datastore{}
	repository := datastore.NewParkingParkingLotPostgresRepository(testLogger, mockDatastore)
	testCtx := context.Background()

	testLots := []entities.ParkingLot{
		{
			ID:            uuid.New(),
			Name:          "vvv",
			Capacity:      10,
			ParkingSpaces: nil,
		},
		{
			ID:            uuid.New(),
			Name:          "vvv",
			Capacity:      10,
			ParkingSpaces: nil,
		},
		{
			ID:            uuid.New(),
			Name:          "vvv",
			Capacity:      10,
			ParkingSpaces: nil,
		},
	}

	var lots []entities.ParkingLot
	mockDatastore.EXPECT().WithContext(testCtx).Return(mockDatastore).Once()
	mockDatastore.EXPECT().Preload("ParkingSpaces").Return(mockDatastore).Once()
	mockDatastore.EXPECT().FindAll(&lots).Return(mockDatastore).Once().Run(func(args mock.Arguments) {
		arg := args.Get(0).(*[]entities.ParkingLot) // Get the first argument passed to FindAll()
		*arg = testLots                             // Set it to the expected park reqs
	})
	mockDatastore.EXPECT().Error().Return(nil).Once()

	// --------
	// ACT
	// --------
	actualLots, err := repository.GetAllParkingLots(testCtx)

	// --------
	// ASSERT
	// --------
	assert.Nil(t, err, "Error must be nil")
	assert.Equal(t, testLots, actualLots, "Parking lots retunred do not equal expected")
	assert.Equal(t, 0, len(hook.Entries), "Logger shouldn't log anything")
	mockDatastore.AssertExpectations(t)
}

func TestParkingLotPostgresRepository_DeleteParkingLot_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockDatastore := &mocks.Datastore{}
	repository := datastore.NewParkingParkingLotPostgresRepository(testLogger, mockDatastore)
	testCtx := context.Background()
	testID := uuid.New()

	mockDatastore.EXPECT().WithContext(testCtx).Return(mockDatastore).Once()
	mockDatastore.EXPECT().Delete(&entities.ParkingLot{}, testID).Return(mockDatastore).Once()
	mockDatastore.EXPECT().Error().Return(nil).Once()

	// --------
	// ACT
	// --------
	err := repository.DeleteParkingLot(testCtx, testID)

	// --------
	// ASSERT
	// --------
	assert.Nil(t, err, "Error must be nil")
	mockDatastore.AssertExpectations(t)
}

func TestParkingLotPostgresRepository_DeleteParkingLot_UnhappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockDatastore := &mocks.Datastore{}
	repository := datastore.NewParkingParkingLotPostgresRepository(testLogger, mockDatastore)
	testCtx := context.Background()
	testID := uuid.New()
	testError := errors.New("boom")

	mockDatastore.EXPECT().WithContext(testCtx).Return(mockDatastore).Once()
	mockDatastore.EXPECT().Delete(&entities.ParkingLot{}, testID).Return(mockDatastore).Once()
	mockDatastore.EXPECT().Error().Return(testError).Once()

	// --------
	// ACT
	// --------
	err := repository.DeleteParkingLot(testCtx, testID)

	// --------
	// ASSERT
	// --------
	assert.IsType(t, &repositories.InternalError{}, err, "Error is of wrong type")
	assert.EqualError(t, err, "Internal error: failed to delete a parking lot", "Error is wrong ")
	mockDatastore.AssertExpectations(t)
}
