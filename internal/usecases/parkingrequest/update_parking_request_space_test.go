package usecases_test

import (
	"context"
	"errors"
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
	"github.com/stretchr/testify/mock"
)

func TestUpdateParkingRequestSpace_Execute_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockRequestRepository := &mocks.ParkingRequestRepository{}
	mockSpaceRepository := &mocks.ParkingSpaceRepository{}
	usecase := usecases.NewUpdateParkingRequestSpace(testLogger, mockRequestRepository, mockSpaceRepository)

	testCtx := context.Background()
	testRequestID := uuid.New()
	testParkingSpaceID := uuid.New()
	testParkingLotID := uuid.New()
	testRequest := entities.ParkingRequest{
		ID:                      testRequestID,
		UserID:                  uuid.New(),
		DestinationParkingLotID: testParkingLotID,
		StartTime:               time.Now(),
		EndTime:                 time.Now().Add(500),
		Status:                  entities.RequestStatusPending,
	}
	testParkSpace := entities.ParkingSpace{
		ID:           testParkingSpaceID,
		ParkingLotID: testParkingLotID,
		Status:       entities.StatusAvailable,
	}

	// Setup mocks
	mockRequestRepository.EXPECT().GetParkingRequestByID(testCtx, testRequestID).Return(testRequest, nil).Once()
	mockSpaceRepository.EXPECT().GetParkingSpaceByID(testCtx, testParkingSpaceID).Return(testParkSpace, nil).Once()

	// Update both entitites.
	testParkSpace.OnAssign(testRequest.StartTime, testRequest.EndTime, testRequest.UserID)
	testRequest.ParkingSpaceID = &testParkSpace.ID

	mockSpaceRepository.EXPECT().Save(testCtx, &testParkSpace).Return(nil).Once()
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

func TestUpdateParkingRequestSpace_Execute_UnhappyPath_WrongParkingSpace(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockRequestRepository := &mocks.ParkingRequestRepository{}
	mockSpaceRepository := &mocks.ParkingSpaceRepository{}
	usecase := usecases.NewUpdateParkingRequestSpace(testLogger, mockRequestRepository, mockSpaceRepository)

	testCtx := context.Background()
	testRequestID := uuid.New()
	testParkingSpaceID := uuid.New()
	testParkingLotID := uuid.New()
	testRequest := entities.ParkingRequest{
		ID:                      testRequestID,
		UserID:                  uuid.New(),
		DestinationParkingLotID: testParkingLotID,
		StartTime:               time.Now(),
		EndTime:                 time.Now().Add(500),
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

func TestUpdateParkingRequestSpace_Execute_UnhappyPath_ParkingSpaceUnavailable(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockRequestRepository := &mocks.ParkingRequestRepository{}
	mockSpaceRepository := &mocks.ParkingSpaceRepository{}
	usecase := usecases.NewUpdateParkingRequestSpace(testLogger, mockRequestRepository, mockSpaceRepository)

	testCtx := context.Background()
	testRequestID := uuid.New()
	testParkingSpaceID := uuid.New()
	testParkingLotID := uuid.New()
	testRequest := entities.ParkingRequest{
		ID:                      testRequestID,
		UserID:                  uuid.New(),
		DestinationParkingLotID: testParkingLotID,
		StartTime:               time.Now(),
		EndTime:                 time.Now().Add(500),
		Status:                  entities.RequestStatusPending,
	}
	testParkSpace := entities.ParkingSpace{
		ID:           testParkingSpaceID,
		ParkingLotID: testParkingLotID,
		Status:       entities.StatusBlocked, // Set status to blocked.
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
	assert.EqualError(t, err, "parking space isn't available", "Must return error")
	assert.Equal(t, 1, len(hook.Entries), "Logger should've loged 1 time")
	assert.Equal(t, logrus.WarnLevel, hook.LastEntry().Level, "Level must be warn")
	assert.Equal(t, "parking space isn't available", hook.LastEntry().Message, "Log message is wrong")
	assert.Equal(t, testParkSpace.Status, hook.LastEntry().Data["status"], "Wrong field in the logger")

	mockRequestRepository.AssertExpectations(t)
	mockSpaceRepository.AssertExpectations(t)
}

func TestUpdateParkingRequestSpace_Execute_UnhappyPath_RepoErrors(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockRequestRepository := &mocks.ParkingRequestRepository{}
	mockSpaceRepository := &mocks.ParkingSpaceRepository{}
	usecase := usecases.NewUpdateParkingRequestSpace(testLogger, mockRequestRepository, mockSpaceRepository)

	testRequestID := uuid.New()
	testParkingLotID := uuid.New()
	testSpaceID := uuid.New()
	testError := errors.New("boom")

	testRequest := entities.ParkingRequest{
		ID:                      testRequestID,
		UserID:                  uuid.New(),
		DestinationParkingLotID: testParkingLotID,
		StartTime:               time.Now(),
		EndTime:                 time.Now().Add(500),
		Status:                  entities.RequestStatusPending,
	}
	testParkSpace := entities.ParkingSpace{
		ID:           testSpaceID,
		ParkingLotID: testParkingLotID,
		Status:       entities.StatusAvailable,
	}

	tests := []struct {
		name          string
		setupMocks    func()
		expectedError error
	}{
		{
			name: "Error retrieving parking space",
			setupMocks: func() {
				mockSpaceRepository.EXPECT().GetParkingSpaceByID(mock.Anything, testSpaceID).Return(entities.ParkingSpace{}, testError).Once()
			},
			expectedError: testError,
		},
		{
			name: "Error retrieving parking request",
			setupMocks: func() {
				mockSpaceRepository.EXPECT().GetParkingSpaceByID(mock.Anything, testSpaceID).Return(entities.ParkingSpace{}, nil).Once()
				mockRequestRepository.EXPECT().GetParkingRequestByID(mock.Anything, testRequestID).Return(entities.ParkingRequest{}, testError).Once()
			},
			expectedError: testError,
		},
		{
			name: "Error saving updated parking space",
			setupMocks: func() {
				mockSpaceRepository.EXPECT().GetParkingSpaceByID(mock.Anything, testSpaceID).Return(testParkSpace, nil).Once()
				mockRequestRepository.EXPECT().GetParkingRequestByID(mock.Anything, testRequestID).Return(testRequest, nil).Once()
				// Update parking space.
				testParkSpace.OnAssign(testRequest.StartTime, testRequest.EndTime, testRequest.UserID)
				mockSpaceRepository.EXPECT().Save(mock.Anything, &testParkSpace).Return(testError).Once()
			},
			expectedError: testError,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMocks()

			// -----
			// ACT
			// -----
			err := usecase.Execute(context.Background(), testRequestID, testSpaceID)

			// ------
			// ASSERT
			// ------
			assert.Equal(t, tc.expectedError, err)
			mockRequestRepository.AssertExpectations(t)
			mockSpaceRepository.AssertExpectations(t)
		})
	}
}
