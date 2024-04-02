package usecases_test

import (
	"context"
	"errors"
	"testing"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	usecases "github.com/IgorSteps/easypark/internal/usecases/parkingrequest"
	mocks "github.com/IgorSteps/easypark/mocks/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestUpdateParkingRequestStatus_Execute_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testCtx := context.Background()
	testID := uuid.New()
	testLogger, _ := test.NewNullLogger()
	mockRepo := &mocks.ParkingRequestRepository{}
	usecase := usecases.NewUpdateParkingRequestStatus(testLogger, mockRepo)

	testRequest := &entities.ParkingRequest{
		ID:     testID,
		Status: entities.RequestStatusPending,
	}
	mockRepo.EXPECT().GetParkingRequestByID(testCtx, testID).Return(*testRequest, nil).Once()

	testRequest.Status = entities.RequestStatusApproved
	mockRepo.EXPECT().Save(testCtx, testRequest).Return(nil).Once()

	// ----
	// ACT
	// ----
	err := usecase.Execute(testCtx, testID, entities.RequestStatusApproved)

	// ------
	// ASSERT
	// ------
	assert.Nil(t, err, "Must not error")
	assert.Equal(t, entities.RequestStatusApproved, testRequest.Status, "Status is wrong")
	mockRepo.AssertExpectations(t)
}

func TestUpdateParkingRequestStatus_Execute_UnhappyPath_GetParkingRequestByIDError(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testCtx := context.Background()
	testID := uuid.New()
	testLogger, _ := test.NewNullLogger()
	mockRepo := &mocks.ParkingRequestRepository{}
	usecase := usecases.NewUpdateParkingRequestStatus(testLogger, mockRepo)

	testError := errors.New("boom")
	testRequest := entities.ParkingRequest{}
	mockRepo.EXPECT().GetParkingRequestByID(testCtx, testID).Return(testRequest, testError).Once()

	// ----
	// ACT
	// ----
	err := usecase.Execute(testCtx, testID, entities.RequestStatusApproved)

	// ------
	// ASSERT
	// ------
	assert.EqualError(t, err, testError.Error(), "Must return the error")
	assert.Empty(t, testRequest, "Request must empty")
	mockRepo.AssertExpectations(t)
}

func TestUpdateParkingRequestStatus_Execute_UnhappyPath_SaveError(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testCtx := context.Background()
	testID := uuid.New()
	testLogger, _ := test.NewNullLogger()
	mockRepo := &mocks.ParkingRequestRepository{}
	usecase := usecases.NewUpdateParkingRequestStatus(testLogger, mockRepo)

	testRequest := &entities.ParkingRequest{}
	mockRepo.EXPECT().GetParkingRequestByID(testCtx, testID).Return(*testRequest, nil).Once()

	testError := errors.New("boom")
	testRequest.Status = entities.RequestStatusApproved
	mockRepo.EXPECT().Save(testCtx, testRequest).Return(testError).Once()

	// ----
	// ACT
	// ----
	err := usecase.Execute(testCtx, testID, entities.RequestStatusApproved)

	// ------
	// ASSERT
	// ------
	assert.EqualError(t, err, testError.Error(), "Must return the error")
	mockRepo.AssertExpectations(t)
}
