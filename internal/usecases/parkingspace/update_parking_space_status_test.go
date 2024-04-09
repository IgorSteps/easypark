package usecases_test

import (
	"context"
	"errors"
	"testing"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	usecases "github.com/IgorSteps/easypark/internal/usecases/parkingspace"
	mocks "github.com/IgorSteps/easypark/mocks/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestUpdateParkingSpaceStatus_Execute_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockRepo := &mocks.ParkingSpaceRepository{}
	usecase := usecases.NewUpdateParkingSpaceStatus(testLogger, mockRepo)

	testCtx := context.Background()
	testID := uuid.New()
	testStatus := "available"

	testParkingSpace := entities.ParkingSpace{
		ID:           uuid.New(),
		ParkingLotID: uuid.New(),
		Name:         "main lot",
		Status:       entities.StatusBlocked,
	}
	mockRepo.EXPECT().GetParkingSpaceByID(testCtx, testID).Return(testParkingSpace, nil).Once()

	testParkingSpace.Status = entities.StatusAvailable
	mockRepo.EXPECT().Save(testCtx, &testParkingSpace).Return(nil).Once()

	// --------
	// ACT
	// --------
	space, err := usecase.Execute(testCtx, testID, testStatus)

	// --------
	// ASSERT
	// --------
	assert.Nil(t, err, "Error must be nil")
	assert.Equal(t, entities.StatusAvailable, space.Status, "Space status must be available")
	mockRepo.AssertExpectations(t)
}

func TestUpdateParkingSpaceStatus_Execute_UnhappyPath_FailedStatusParsing(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockRepo := &mocks.ParkingSpaceRepository{}
	usecase := usecases.NewUpdateParkingSpaceStatus(testLogger, mockRepo)

	testCtx := context.Background()
	testID := uuid.New()
	testStatus := "invalid"

	// --------
	// ACT
	// --------
	space, err := usecase.Execute(testCtx, testID, testStatus)

	// --------
	// ASSERT
	// --------
	assert.IsType(t, &repositories.InvalidInputError{}, err, "Error is of wrong type")
	assert.EqualError(t, err, "failed to parse given status")
	assert.Empty(t, space, "Space must be empty")

	assert.Equal(t, 1, len(hook.Entries), "Logger must only log once")
	assert.Equal(t, "failed to parse given status", hook.LastEntry().Message, "Log message is wrong")
	assert.Equal(t, testStatus, hook.LastEntry().Data["status"], "Status field is wrong")
	assert.Equal(t, err, hook.LastEntry().Data["error"], "Error field is wrong")
	mockRepo.AssertExpectations(t)
}

func TestUpdateParkingSpaceStatus_Execute_UnhappyPath_GetParkingSpaceError(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockRepo := &mocks.ParkingSpaceRepository{}
	usecase := usecases.NewUpdateParkingSpaceStatus(testLogger, mockRepo)

	testCtx := context.Background()
	testID := uuid.New()
	testStatus := "available"

	testError := errors.New("boom")
	mockRepo.EXPECT().GetParkingSpaceByID(testCtx, testID).Return(entities.ParkingSpace{}, testError).Once()

	// --------
	// ACT
	// --------
	space, err := usecase.Execute(testCtx, testID, testStatus)

	// --------
	// ASSERT
	// --------
	assert.EqualError(t, err, "boom")
	assert.Empty(t, space, "Space must be empty")

	mockRepo.AssertExpectations(t)
}
