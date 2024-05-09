package usecases_test

import (
	"context"
	"testing"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	usecases "github.com/IgorSteps/easypark/internal/usecases/parkingrequest"
	mocks "github.com/IgorSteps/easypark/mocks/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestDeassignParkingSpace_Execute(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	mockRepo := &mocks.ParkingRequestRepository{}
	testLogger, _ := test.NewNullLogger()
	usecase := usecases.NewDeassignParkingSpace(testLogger, mockRepo)

	testCtx := context.Background()
	requestID := uuid.New()

	spaceID := uuid.New()
	testRequest := entities.ParkingRequest{
		ID:             uuid.New(),
		Status:         entities.RequestStatusApproved,
		ParkingSpaceID: &spaceID,
	}
	mockRepo.EXPECT().GetSingle(testCtx, requestID).Return(testRequest, nil).Once()

	testRequest.OnSpaceDeassign()
	mockRepo.EXPECT().Save(testCtx, &testRequest).Return(nil).Once()

	// --------
	// ACT
	// --------
	err := usecase.Execute(testCtx, requestID)

	// --------
	// ASSERT
	// --------
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
