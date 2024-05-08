package usecases_test

import (
	"context"
	"errors"
	"testing"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	usecases "github.com/IgorSteps/easypark/internal/usecases/parkingspace"
	mocks "github.com/IgorSteps/easypark/mocks/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestGetSingleParkingSpace_Execute_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockRepo := &mocks.ParkingSpaceRepository{}
	usecase := usecases.NewGetSingleParkingSpace(testLogger, mockRepo)

	testCtx := context.Background()
	testID := uuid.New()

	testParkingSpace := entities.ParkingSpace{
		ID:           testID,
		ParkingLotID: uuid.New(),
		Name:         "main lot",
		Status:       entities.ParkingSpaceStatusAvailable,
	}
	mockRepo.EXPECT().GetSingle(testCtx, testID).Return(testParkingSpace, nil).Once()

	// --------
	// ACT
	// --------
	resultSpace, err := usecase.Execute(testCtx, testID)

	// --------
	// ASSERT
	// --------
	assert.Nil(t, err, "Error must be nil")
	assert.Equal(t, testParkingSpace, resultSpace, "Returned space should match expected space")
	mockRepo.AssertExpectations(t)
}

func TestGetSingleParkingSpace_Execute_UnhappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockRepo := &mocks.ParkingSpaceRepository{}
	usecase := usecases.NewGetSingleParkingSpace(testLogger, mockRepo)

	testCtx := context.Background()
	testID := uuid.New()

	testError := errors.New("data retrieval error")
	mockRepo.EXPECT().GetSingle(testCtx, testID).Return(entities.ParkingSpace{}, testError).Once()

	// --------
	// ACT
	// --------
	resultSpace, err := usecase.Execute(testCtx, testID)

	// --------
	// ASSERT
	// --------
	assert.EqualError(t, err, "data retrieval error")
	assert.Empty(t, resultSpace, "Result space should be empty due to error")

	mockRepo.AssertExpectations(t)
}
