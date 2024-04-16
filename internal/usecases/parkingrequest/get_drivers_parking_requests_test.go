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

func TestGetDriversParkingRequests_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testCtx := context.Background()
	testLogger, _ := test.NewNullLogger()
	mockRepo := &mocks.ParkingRequestRepository{}
	usecase := usecases.NewGetDriversParkingRequests(testLogger, mockRepo)

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
	testID := uuid.New()
	query := map[string]interface{}{
		"user_id": testID.String(),
	}
	mockRepo.EXPECT().GetMany(testCtx, query).Return(testParkRequests, nil).Once()

	// ------
	// ACT
	// ------
	reqs, err := usecase.Execute(testCtx, testID)

	// ------
	// ASSERT
	// ------
	assert.Nil(t, err, "Error must be nil")
	assert.Equal(t, testParkRequests, reqs, "Requests don't match")
}

func TestGetDriversParkingRequests_UnhappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testCtx := context.Background()
	testLogger, _ := test.NewNullLogger()
	mockRepo := &mocks.ParkingRequestRepository{}
	usecase := usecases.NewGetDriversParkingRequests(testLogger, mockRepo)
	testID := uuid.New()
	testErr := errors.New("boom")
	query := map[string]interface{}{
		"user_id": testID.String(),
	}
	mockRepo.EXPECT().GetMany(testCtx, query).Return(nil, testErr).Once()
	// ------
	// ACT
	// ------
	reqs, err := usecase.Execute(testCtx, testID)

	// ------
	// ASSERT
	// ------
	assert.EqualError(t, err, "boom", "Error is wrong")
	assert.Empty(t, reqs, "Requests slice must be empty")
}
