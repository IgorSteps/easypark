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

func TestGetAllParkingRequests_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testCtx := context.Background()
	testLogger, _ := test.NewNullLogger()
	mockRepo := &mocks.ParkingRequestRepository{}
	usecase := usecases.NewGetAllParkingRequests(testLogger, mockRepo)

	// Don't set times here because they break tests.
	testParkRequests := []entities.ParkingRequest{
		{
			ID:                      uuid.New(),
			UserID:                  uuid.New(),
			DestinationParkingLotID: uuid.New(),
			Status:                  entities.RequestStatusApproved,
		},
		{
			ID:                      uuid.New(),
			UserID:                  uuid.New(),
			DestinationParkingLotID: uuid.New(),
			Status:                  entities.RequestStatusApproved,
		},
	}

	mockRepo.EXPECT().GetMany(testCtx, map[string]interface{}(nil)).Return(testParkRequests, nil).Once()

	// ------
	// ACT
	// ------
	reqs, err := usecase.Execute(testCtx)

	// ------
	// ASSERT
	// ------
	assert.Nil(t, err, "Error must be nil")
	assert.Equal(t, testParkRequests, reqs, "Requests don't match")
}

func TestGetAllParkingRequests_UnhappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testCtx := context.Background()
	testLogger, _ := test.NewNullLogger()
	mockRepo := &mocks.ParkingRequestRepository{}
	usecase := usecases.NewGetAllParkingRequests(testLogger, mockRepo)

	testErr := errors.New("boom")
	mockRepo.EXPECT().GetMany(testCtx, map[string]interface{}(nil)).Return([]entities.ParkingRequest{}, testErr).Once()

	// ------
	// ACT
	// ------
	reqs, err := usecase.Execute(testCtx)

	// ------
	// ASSERT
	// ------
	assert.EqualError(t, err, "boom", "Error is wrong")
	assert.Empty(t, reqs, "Requests slice must be empty")
}
