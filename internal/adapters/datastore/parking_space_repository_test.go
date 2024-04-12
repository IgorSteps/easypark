package datastore_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/IgorSteps/easypark/internal/adapters/datastore"
	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	mocks "github.com/IgorSteps/easypark/mocks/adapters/datastore"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestParkingSpaceRepository_GetParkingSpaceByID_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockDB := &mocks.Datastore{}
	repo := datastore.NewParkingSpacePostgresRepository(testLogger, mockDB)

	pakingSpace := &entities.ParkingSpace{}
	testID := uuid.New()

	mockDB.EXPECT().WithContext(mock.Anything).Return(mockDB).Once()
	mockDB.EXPECT().Preload("ParkingRequests").Return(mockDB).Once()
	mockDB.EXPECT().First(pakingSpace, "id = ?", testID).Return(mockDB).Once()
	mockDB.EXPECT().Error().Return(nil).Once()

	// ----
	// ACT
	// ----
	_, err := repo.GetParkingSpaceByID(context.Background(), testID)

	// ------
	// ASSERT
	// ------
	assert.Nil(t, err, "error must be nil")
	mockDB.AssertExpectations(t)
}

func TestParkingSpaceRepository_GetParkingSpaceByID_UnhappyPath_NotFound(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockDB := &mocks.Datastore{}
	repo := datastore.NewParkingSpacePostgresRepository(testLogger, mockDB)

	pakingSpace := &entities.ParkingSpace{}
	testID := uuid.New()

	mockDB.EXPECT().WithContext(mock.Anything).Return(mockDB).Once()
	mockDB.EXPECT().Preload("ParkingRequests").Return(mockDB).Once()
	mockDB.EXPECT().First(pakingSpace, "id = ?", testID).Return(mockDB).Once()
	mockDB.EXPECT().Error().Return(gorm.ErrRecordNotFound).Once()

	// ----
	// ACT
	// ----
	_, err := repo.GetParkingSpaceByID(context.Background(), testID)

	// ------
	// ASSERT
	// ------
	assert.EqualError(t, err, fmt.Sprintf("Resource '%s' not found", testID), "Errors must match")
	assert.IsType(t, &repositories.NotFoundError{}, err, "Wrong type of error")
	mockDB.AssertExpectations(t)
}

func TestParkingSpaceRepository_GetParkingSpaceByID_UnhappyPath_Internal(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockDB := &mocks.Datastore{}
	repo := datastore.NewParkingSpacePostgresRepository(testLogger, mockDB)

	pakingSpace := &entities.ParkingSpace{}
	testID := uuid.New()
	testError := errors.New("boom")

	mockDB.EXPECT().WithContext(mock.Anything).Return(mockDB).Once()
	mockDB.EXPECT().Preload("ParkingRequests").Return(mockDB).Once()
	mockDB.EXPECT().First(pakingSpace, "id = ?", testID).Return(mockDB).Once()
	mockDB.EXPECT().Error().Return(testError).Once()

	// ----
	// ACT
	// ----
	_, err := repo.GetParkingSpaceByID(context.Background(), testID)

	// ------
	// ASSERT
	// ------
	assert.EqualError(t, err, "Internal error: failed to query for parking space in the database", "Errors must match")
	assert.IsType(t, &repositories.InternalError{}, err, "Wrong type of error")
	mockDB.AssertExpectations(t)
}

func TestParkingSpaceRepository_Save_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockDB := &mocks.Datastore{}
	repo := datastore.NewParkingSpacePostgresRepository(testLogger, mockDB)

	pakingSpace := &entities.ParkingSpace{}

	mockDB.EXPECT().WithContext(mock.Anything).Return(mockDB).Once()
	mockDB.EXPECT().Save(pakingSpace).Return(mockDB).Once()
	mockDB.EXPECT().Error().Return(nil).Once()

	// ----
	// ACT
	// ----
	err := repo.Save(context.Background(), pakingSpace)

	// ------
	// ASSERT
	// ------
	assert.Nil(t, err, "error must be nil")
	mockDB.AssertExpectations(t)
}

func TestParkingSpaceRepository_Save_UnhappyPath_Internal(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockDB := &mocks.Datastore{}
	repo := datastore.NewParkingSpacePostgresRepository(testLogger, mockDB)

	pakingSpace := &entities.ParkingSpace{}

	mockDB.EXPECT().WithContext(mock.Anything).Return(mockDB).Once()
	mockDB.EXPECT().Save(pakingSpace).Return(mockDB).Once()
	mockDB.EXPECT().Error().Return(errors.New("boom")).Once()

	// ----
	// ACT
	// ----
	err := repo.Save(context.Background(), pakingSpace)

	// ------
	// ASSERT
	// ------
	assert.IsType(t, &repositories.InternalError{}, err, "Wrong error type")
	assert.EqualError(t, err, "Internal error: failed to save updated parking space in the database", "Error message is wrong")
	mockDB.AssertExpectations(t)
}
