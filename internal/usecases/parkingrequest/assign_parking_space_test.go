package usecases_test

import (
	"context"
	"testing"
	"time"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	usecases "github.com/IgorSteps/easypark/internal/usecases/parkingrequest"
	mocks "github.com/IgorSteps/easypark/mocks/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestAssignParkingSpace_HappyPath_Available(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockRequestRepository := &mocks.ParkingRequestRepository{}
	mockSpaceRepository := &mocks.ParkingSpaceRepository{}
	usecase := usecases.NewAssignParkingSpace(testLogger, mockRequestRepository, mockSpaceRepository)

	testCtx := context.Background()
	testRequestID := uuid.New()
	testParkingSpaceID := uuid.New()
	testParkingLotID := uuid.New()
	testRequest := entities.ParkingRequest{
		ID:                      testRequestID,
		UserID:                  uuid.New(),
		DestinationParkingLotID: testParkingLotID,
		StartTime:               time.Now().Add(5 * time.Minute),
		EndTime:                 time.Now().Add(9 * time.Minute),
		Status:                  entities.RequestStatusPending,
	}
	testParkSpace := entities.ParkingSpace{
		ID:              testParkingSpaceID,
		ParkingLotID:    testParkingLotID,
		Status:          entities.StatusAvailable,
		ParkingRequests: NewTestParkRequests(),
	}

	// Setup mocks.
	mockRequestRepository.EXPECT().GetParkingRequestByID(testCtx, testRequestID).Return(testRequest, nil).Once()
	mockSpaceRepository.EXPECT().GetParkingSpaceByID(testCtx, testParkingSpaceID).Return(testParkSpace, nil).Once()

	// Update parking request like we do in the usecase.
	testRequest.OnSpaceAssign(testParkingSpaceID)
	mockRequestRepository.EXPECT().Save(testCtx, &testRequest).Return(nil).Once()

	// ------
	// ACT
	// ------
	err := usecase.Execute(testCtx, testRequestID, testParkingSpaceID)

	// ------
	// ASSERT
	// ------
	assert.NoError(t, err, "Must not return error")

	mockRequestRepository.AssertExpectations(t)
	mockSpaceRepository.AssertExpectations(t)
}

func TestAssignParkingSpace_HappyPath_Occupied(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockRequestRepository := &mocks.ParkingRequestRepository{}
	mockSpaceRepository := &mocks.ParkingSpaceRepository{}
	usecase := usecases.NewAssignParkingSpace(testLogger, mockRequestRepository, mockSpaceRepository)

	testCtx := context.Background()
	testRequestID := uuid.New()
	testParkingSpaceID := uuid.New()
	testParkingLotID := uuid.New()
	testRequest := entities.ParkingRequest{
		ID:                      testRequestID,
		UserID:                  uuid.New(),
		DestinationParkingLotID: testParkingLotID,
		StartTime:               time.Now().Add(5 * time.Minute),
		EndTime:                 time.Now().Add(9 * time.Minute),
		Status:                  entities.RequestStatusPending,
	}
	testParkSpace := entities.ParkingSpace{
		ID:              testParkingSpaceID,
		ParkingLotID:    testParkingLotID,
		Status:          entities.StatusOccupied,
		ParkingRequests: NewTestParkRequests(),
	}

	// Setup mocks.
	mockRequestRepository.EXPECT().GetParkingRequestByID(testCtx, testRequestID).Return(testRequest, nil).Once()
	mockSpaceRepository.EXPECT().GetParkingSpaceByID(testCtx, testParkingSpaceID).Return(testParkSpace, nil).Once()

	// Update parking request like we do in the usecase.
	testRequest.OnSpaceAssign(testParkingSpaceID)
	mockRequestRepository.EXPECT().Save(testCtx, &testRequest).Return(nil).Once()

	// ------
	// ACT
	// ------
	err := usecase.Execute(testCtx, testRequestID, testParkingSpaceID)

	// ------
	// ASSERT
	// ------
	assert.NoError(t, err, "Must not return error")

	mockRequestRepository.AssertExpectations(t)
	mockSpaceRepository.AssertExpectations(t)
}

// This tests unhappy path when a parking request is outdated.
func TestAssignParkingSpace_UnhappyPath_OutdatedParkingRequest(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockRequestRepository := &mocks.ParkingRequestRepository{}
	mockSpaceRepository := &mocks.ParkingSpaceRepository{}
	usecase := usecases.NewAssignParkingSpace(testLogger, mockRequestRepository, mockSpaceRepository)

	testCtx := context.Background()
	testRequestID := uuid.New()
	testParkingSpaceID := uuid.New()
	testParkingLotID := uuid.New()
	testRequest := entities.ParkingRequest{
		ID:                      testRequestID,
		UserID:                  uuid.New(),
		DestinationParkingLotID: testParkingLotID,
		StartTime:               time.Now(), // Set time to now, which will be less than the current time when the Execute() runs
		EndTime:                 time.Now().Add(9 * time.Minute),
		Status:                  entities.RequestStatusPending,
	}
	testParkSpace := entities.ParkingSpace{
		ID:              testParkingSpaceID,
		ParkingLotID:    testParkingLotID,
		Status:          entities.StatusAvailable,
		ParkingRequests: NewTestParkRequests(),
	}

	// Setup mocks.
	mockRequestRepository.EXPECT().GetParkingRequestByID(testCtx, testRequestID).Return(testRequest, nil).Once()
	mockSpaceRepository.EXPECT().GetParkingSpaceByID(testCtx, testParkingSpaceID).Return(testParkSpace, nil).Once()

	// ------
	// ACT
	// ------
	err := usecase.Execute(testCtx, testRequestID, testParkingSpaceID)

	// ------
	// ASSERT
	// ------
	assert.EqualError(t, err, "not allowed to assign a parking space to a parking request with the desired start time in the past", "Error is wrong")
	assert.IsType(t, &repositories.InvalidInputError{}, err, "Returned error is of the wrong type")
	mockRequestRepository.AssertExpectations(t)
	mockSpaceRepository.AssertExpectations(t)
}

func TestAssignParkingSpace_UnhappyPath_Overlap(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockRequestRepository := &mocks.ParkingRequestRepository{}
	mockSpaceRepository := &mocks.ParkingSpaceRepository{}
	usecase := usecases.NewAssignParkingSpace(testLogger, mockRequestRepository, mockSpaceRepository)

	testCtx := context.Background()
	testRequestID := uuid.New()
	testParkingSpaceID := uuid.New()
	testParkingLotID := uuid.New()
	testRequest := entities.ParkingRequest{
		ID:                      testRequestID,
		UserID:                  uuid.New(),
		DestinationParkingLotID: testParkingLotID,
		// set time to overlap with one of the existing  park requests
		StartTime: time.Now().Add(10 * time.Minute),
		EndTime:   time.Now().Add(30 * time.Minute),
		Status:    entities.RequestStatusPending,
	}
	testParkSpace := entities.ParkingSpace{
		ID:              testParkingSpaceID,
		ParkingLotID:    testParkingLotID,
		Status:          entities.StatusAvailable,
		ParkingRequests: NewTestParkRequests(),
	}

	// Setup mocks.
	mockRequestRepository.EXPECT().GetParkingRequestByID(testCtx, testRequestID).Return(testRequest, nil).Once()
	mockSpaceRepository.EXPECT().GetParkingSpaceByID(testCtx, testParkingSpaceID).Return(testParkSpace, nil).Once()

	// ------
	// ACT
	// ------
	err := usecase.Execute(testCtx, testRequestID, testParkingSpaceID)

	// ------
	// ASSERT
	// ------
	assert.EqualError(t, err, "there is an overlap with existing parking requests time slots", "Error is wrong")
	assert.IsType(t, &repositories.InvalidInputError{}, err, "Returned error is of the wrong type")
	mockRequestRepository.AssertExpectations(t)
	mockSpaceRepository.AssertExpectations(t)
}

// This tests unhappy path when a parking request is already rejected.
func TestAssignParkingSpace_UnhappyPath_RejectedParkingRequest(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockRequestRepository := &mocks.ParkingRequestRepository{}
	mockSpaceRepository := &mocks.ParkingSpaceRepository{}
	usecase := usecases.NewAssignParkingSpace(testLogger, mockRequestRepository, mockSpaceRepository)

	testCtx := context.Background()
	testRequestID := uuid.New()
	testParkingSpaceID := uuid.New()
	testParkingLotID := uuid.New()
	testRequest := entities.ParkingRequest{
		ID:                      testRequestID,
		UserID:                  uuid.New(),
		DestinationParkingLotID: testParkingLotID,
		StartTime:               time.Now().Add(5 * time.Minute),
		EndTime:                 time.Now().Add(9 * time.Minute),
		Status:                  entities.RequestStatusRejected, // set to rejected
	}
	testParkSpace := entities.ParkingSpace{
		ID:              testParkingSpaceID,
		ParkingLotID:    testParkingLotID,
		Status:          entities.StatusAvailable,
		ParkingRequests: NewTestParkRequests(),
	}

	// Setup mocks.
	mockRequestRepository.EXPECT().GetParkingRequestByID(testCtx, testRequestID).Return(testRequest, nil).Once()
	mockSpaceRepository.EXPECT().GetParkingSpaceByID(testCtx, testParkingSpaceID).Return(testParkSpace, nil).Once()

	// ------
	// ACT
	// ------
	err := usecase.Execute(testCtx, testRequestID, testParkingSpaceID)

	// ------
	// ASSERT
	// ------
	assert.EqualError(t, err, "not allowed to assign parking space to a 'rejected' parking request", "Error is wrong")
	assert.IsType(t, &repositories.InvalidInputError{}, err, "Returned error is of the wrong type")
	mockRequestRepository.AssertExpectations(t)
	mockSpaceRepository.AssertExpectations(t)
}

func TestAssignParkingSpace_UnhappyPath_WrongParkingSpace(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockRequestRepository := &mocks.ParkingRequestRepository{}
	mockSpaceRepository := &mocks.ParkingSpaceRepository{}
	usecase := usecases.NewAssignParkingSpace(testLogger, mockRequestRepository, mockSpaceRepository)

	testCtx := context.Background()
	testRequestID := uuid.New()
	testParkingSpaceID := uuid.New()
	testParkingLotID := uuid.New()
	testRequest := entities.ParkingRequest{
		ID:                      testRequestID,
		UserID:                  uuid.New(),
		DestinationParkingLotID: testParkingLotID,
		StartTime:               time.Now().Add(5 * time.Minute),
		EndTime:                 time.Now().Add(9 * time.Minute),
		Status:                  entities.RequestStatusPending,
	}
	testParkSpace := entities.ParkingSpace{
		ID:           testParkingSpaceID,
		ParkingLotID: uuid.New(), // different ID to the one in the testRequest
		Status:       entities.StatusAvailable,
	}

	// Setup mocks
	mockRequestRepository.EXPECT().GetParkingRequestByID(testCtx, testRequestID).Return(testRequest, nil).Once()
	mockSpaceRepository.EXPECT().GetParkingSpaceByID(testCtx, testParkingSpaceID).Return(testParkSpace, nil).Once()

	// ------
	// ACT
	// ------
	err := usecase.Execute(testCtx, testRequestID, testParkingSpaceID)

	// ------
	// ASSERT
	// ------
	assert.IsType(t, &repositories.InvalidInputError{}, err, "Error is of wrong type")
	assert.EqualError(t, err, "parking space is not in the desired parking lot", "Must return error")
	assert.Equal(t, 1, len(hook.Entries), "Logger should've loged 1 time")
	assert.Equal(t, logrus.WarnLevel, hook.LastEntry().Level, "Level must be warn")
	assert.Equal(t, "parking space is not in the desired parking lot", hook.LastEntry().Message, "Log message is wrong")
	assert.Equal(t, testRequest.DestinationParkingLotID, hook.LastEntry().Data["desired"], "Wrong field in the logger")
	assert.Equal(t, testParkSpace.ParkingLotID, hook.LastEntry().Data["actual"], "Wrong field in the logger")
	mockRequestRepository.AssertExpectations(t)
	mockSpaceRepository.AssertExpectations(t)
}

func TestAssignParkingSpace_UnhappyPath_ParkingSpaceBlocked(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockRequestRepository := &mocks.ParkingRequestRepository{}
	mockSpaceRepository := &mocks.ParkingSpaceRepository{}
	usecase := usecases.NewAssignParkingSpace(testLogger, mockRequestRepository, mockSpaceRepository)

	testCtx := context.Background()
	testRequestID := uuid.New()
	testParkingSpaceID := uuid.New()
	testParkingLotID := uuid.New()
	testRequest := entities.ParkingRequest{
		ID:                      testRequestID,
		UserID:                  uuid.New(),
		DestinationParkingLotID: testParkingLotID,
		StartTime:               time.Now().Add(5 * time.Minute),
		EndTime:                 time.Now().Add(9 * time.Minute),
		Status:                  entities.RequestStatusPending,
	}
	testParkSpace := entities.ParkingSpace{
		ID:           testParkingSpaceID,
		ParkingLotID: testParkingLotID,
		Status:       entities.StatusBlocked, // blocked status
	}

	// Setup mocks
	mockRequestRepository.EXPECT().GetParkingRequestByID(testCtx, testRequestID).Return(testRequest, nil).Once()
	mockSpaceRepository.EXPECT().GetParkingSpaceByID(testCtx, testParkingSpaceID).Return(testParkSpace, nil).Once()

	// ------
	// ACT
	// ------

	err := usecase.Execute(testCtx, testRequestID, testParkingSpaceID)

	// ------
	// ASSERT
	// ------
	assert.IsType(t, &repositories.InvalidInputError{}, err, "Error is of wrong type")
	assert.EqualError(t, err, "not allowed to assign blocked parking space", "Must return error")

	mockRequestRepository.AssertExpectations(t)
	mockSpaceRepository.AssertExpectations(t)

}

func TestAssignParkingSpace_UnhappyPath_ParkingSpaceReserved(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockRequestRepository := &mocks.ParkingRequestRepository{}
	mockSpaceRepository := &mocks.ParkingSpaceRepository{}
	usecase := usecases.NewAssignParkingSpace(testLogger, mockRequestRepository, mockSpaceRepository)

	testCtx := context.Background()
	testRequestID := uuid.New()
	testParkingSpaceID := uuid.New()
	testParkingLotID := uuid.New()
	testRequest := entities.ParkingRequest{
		ID:                      testRequestID,
		UserID:                  uuid.New(),
		DestinationParkingLotID: testParkingLotID,
		StartTime:               time.Now().Add(5 * time.Minute),
		EndTime:                 time.Now().Add(9 * time.Minute),
		Status:                  entities.RequestStatusPending,
	}
	testParkSpace := entities.ParkingSpace{
		ID:           testParkingSpaceID,
		ParkingLotID: testParkingLotID,
		Status:       entities.StatusReserved, // reserved status
	}

	// Setup mocks
	mockRequestRepository.EXPECT().GetParkingRequestByID(testCtx, testRequestID).Return(testRequest, nil).Once()
	mockSpaceRepository.EXPECT().GetParkingSpaceByID(testCtx, testParkingSpaceID).Return(testParkSpace, nil).Once()

	// ------
	// ACT
	// ------

	err := usecase.Execute(testCtx, testRequestID, testParkingSpaceID)

	// ------
	// ASSERT
	// ------
	assert.IsType(t, &repositories.InvalidInputError{}, err, "Error is of wrong type")
	assert.EqualError(t, err, "not allowed to assign reserved parking space", "Must return error")

	mockRequestRepository.AssertExpectations(t)
	mockSpaceRepository.AssertExpectations(t)

}

// NewTestParkRequests returns a slice of ParkingRequest with different time slots.
func NewTestParkRequests() []entities.ParkingRequest {
	now := time.Now()
	return []entities.ParkingRequest{
		{
			ID:                      uuid.New(),
			UserID:                  uuid.New(),
			DestinationParkingLotID: uuid.New(),
			StartTime:               now.Add(10 * time.Minute),
			EndTime:                 now.Add(30 * time.Minute), // A request from now for the next 30 minutes
			Status:                  entities.RequestStatusApproved,
		},
		{
			ID:                      uuid.New(),
			UserID:                  uuid.New(),
			DestinationParkingLotID: uuid.New(),
			StartTime:               now.Add(1 * time.Hour), // Starts 1 hour from now
			EndTime:                 now.Add(2 * time.Hour), // Ends 2 hours from now
			Status:                  entities.RequestStatusApproved,
		},
		{
			ID:                      uuid.New(),
			UserID:                  uuid.New(),
			DestinationParkingLotID: uuid.New(),
			StartTime:               now.Add(2*time.Hour + 30*time.Minute), // Starts 2.5 hours from now
			EndTime:                 now.Add(3 * time.Hour),                // Ends 3 hours from now
			Status:                  entities.RequestStatusApproved,
		},
		{
			ID:                      uuid.New(),
			UserID:                  uuid.New(),
			DestinationParkingLotID: uuid.New(),
			StartTime:               now.Add(-1 * time.Hour),    // Started 1 hour ago
			EndTime:                 now.Add(-30 * time.Minute), // Ended 30 minutes ago
			Status:                  entities.RequestStatusApproved,
		},
	}
}
