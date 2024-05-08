package usecases_test

import (
	"context"
	"testing"
	"time"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	usecases "github.com/IgorSteps/easypark/internal/usecases/parkingrequest"
	mocks "github.com/IgorSteps/easypark/mocks/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestAutomaticAssignParkingSpace_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	reqRepo := &mocks.ParkingRequestRepository{}
	spaceRepo := &mocks.ParkingSpaceRepository{}
	usecase := usecases.NewAutomaticAssignParkingSpace(testLogger, reqRepo, spaceRepo)
	testCtx := context.Background()
	reqID := uuid.New()

	testRequest := entities.ParkingRequest{
		ID:                      reqID,
		UserID:                  uuid.New(),
		DestinationParkingLotID: uuid.New(),
		StartTime:               time.Now().Add(5 * time.Minute),
		EndTime:                 time.Now().Add(9 * time.Minute),
		Status:                  entities.RequestStatusPending,
	}

	testParkSpaces := []entities.ParkingSpace{
		{
			ID:           uuid.New(),
			ParkingLotID: uuid.New(),
			Status:       entities.ParkingSpaceStatusAvailable,
		},
	}

	reqRepo.EXPECT().GetSingle(testCtx, reqID).Return(testRequest, nil).Once()
	spaceRepo.EXPECT().FindAvailableSpaces(
		testCtx,
		testRequest.DestinationParkingLotID,
		testRequest.StartTime,
		testRequest.EndTime,
	).Return(testParkSpaces, nil).Once()

	// Update parking request like we do in the usecase.
	testRequest.OnSpaceAssign(testParkSpaces[0].ID)
	reqRepo.EXPECT().Save(testCtx, &testRequest).Return(nil).Once()

	// --------
	// ACT
	// --------
	space, err := usecase.Execute(testCtx, reqID)

	// --------
	// ASSERT
	// --------
	assert.NoError(t, err, "Must not return error")
	assert.Equal(t, testParkSpaces[0], *space)
	reqRepo.AssertExpectations(t)
	spaceRepo.AssertExpectations(t)
}

func TestAutomaticAssignParkingSpace_HappyPath_NoSpaces(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	reqRepo := &mocks.ParkingRequestRepository{}
	spaceRepo := &mocks.ParkingSpaceRepository{}
	usecase := usecases.NewAutomaticAssignParkingSpace(testLogger, reqRepo, spaceRepo)
	testCtx := context.Background()
	reqID := uuid.New()

	testRequest := entities.ParkingRequest{
		ID:                      reqID,
		UserID:                  uuid.New(),
		DestinationParkingLotID: uuid.New(),
		StartTime:               time.Now().Add(5 * time.Minute),
		EndTime:                 time.Now().Add(9 * time.Minute),
		Status:                  entities.RequestStatusPending,
	}

	// empty
	testParkSpaces := []entities.ParkingSpace{}

	reqRepo.EXPECT().GetSingle(testCtx, reqID).Return(testRequest, nil).Once()
	spaceRepo.EXPECT().FindAvailableSpaces(
		testCtx,
		testRequest.DestinationParkingLotID,
		testRequest.StartTime,
		testRequest.EndTime,
	).Return(testParkSpaces, nil).Once()

	// --------
	// ACT
	// --------
	space, err := usecase.Execute(testCtx, reqID)

	// --------
	// ASSERT
	// --------
	assert.Error(t, err, "no available parking spaces at the desired time ", "Must not return error")
	assert.Nil(t, space)
	reqRepo.AssertExpectations(t)
	spaceRepo.AssertExpectations(t)
}
