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
	mockReq := &mocks.ParkingRequestRepository{}
	usecase := usecases.NewUpdateParkingSpaceStatus(testLogger, mockRepo, mockReq)

	testCtx := context.Background()
	testID := uuid.New()
	testStatus := "available"

	testParkingSpace := entities.ParkingSpace{
		ID:           uuid.New(),
		ParkingLotID: uuid.New(),
		Name:         "main lot",
		Status:       entities.ParkingSpaceStatusBlocked,
	}
	mockRepo.EXPECT().GetSingle(testCtx, testID).Return(testParkingSpace, nil).Once()

	testParkingSpace.Status = entities.ParkingSpaceStatusAvailable
	mockRepo.EXPECT().Save(testCtx, &testParkingSpace).Return(nil).Once()

	// --------
	// ACT
	// --------
	space, err := usecase.Execute(testCtx, testID, testStatus)

	// --------
	// ASSERT
	// --------
	assert.Nil(t, err, "Error must be nil")
	assert.Equal(t, entities.ParkingSpaceStatusAvailable, space.Status, "Space status must be available")
	mockRepo.AssertExpectations(t)
}

func TestUpdateParkingSpaceStatus_Execute_HappyPath_Block(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockRepo := &mocks.ParkingSpaceRepository{}
	mockReq := &mocks.ParkingRequestRepository{}
	usecase := usecases.NewUpdateParkingSpaceStatus(testLogger, mockRepo, mockReq)

	testCtx := context.Background()
	testID := uuid.New()
	testStatus := "blocked"
	parkSpaceID := uuid.New()
	testParkingRequests := []entities.ParkingRequest{
		{
			ID:             uuid.New(),
			ParkingSpaceID: &parkSpaceID,
		},
		{
			ID:             uuid.New(),
			ParkingSpaceID: &parkSpaceID,
		},
	}
	testParkingSpace := entities.ParkingSpace{
		ID:              parkSpaceID,
		ParkingLotID:    uuid.New(),
		Name:            "main lot",
		Status:          entities.ParkingSpaceStatusBlocked,
		ParkingRequests: testParkingRequests,
	}
	mockRepo.EXPECT().GetSingle(testCtx, testID).Return(testParkingSpace, nil).Once()

	for _, req := range testParkingRequests {
		req.OnSpaceDeassign()
		mockReq.EXPECT().Save(testCtx, &req).Return(nil).Once()
	}

	testParkingSpace.Status = entities.ParkingSpaceStatusBlocked
	testParkingSpace.ParkingRequests = nil
	mockRepo.EXPECT().Save(testCtx, &testParkingSpace).Return(nil).Once()

	// --------
	// ACT
	// --------
	space, err := usecase.Execute(testCtx, testID, testStatus)

	// --------
	// ASSERT
	// --------
	assert.Nil(t, err, "Error must be nil")
	assert.Equal(t, entities.ParkingSpaceStatusBlocked, space.Status, "Space status must be blocked")
	assert.Nil(t, space.ParkingRequests)
	mockRepo.AssertExpectations(t)
	mockReq.AssertExpectations(t)
}

func TestUpdateParkingSpaceStatus_Execute_UnhappyPath_FailedStatusParsing(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockRepo := &mocks.ParkingSpaceRepository{}
	mockReq := &mocks.ParkingRequestRepository{}
	usecase := usecases.NewUpdateParkingSpaceStatus(testLogger, mockRepo, mockReq)
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
	mockReq := &mocks.ParkingRequestRepository{}
	usecase := usecases.NewUpdateParkingSpaceStatus(testLogger, mockRepo, mockReq)
	testCtx := context.Background()
	testID := uuid.New()
	testStatus := "available"

	testError := errors.New("boom")
	mockRepo.EXPECT().GetSingle(testCtx, testID).Return(entities.ParkingSpace{}, testError).Once()

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
