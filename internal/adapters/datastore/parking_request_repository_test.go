package datastore_test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/IgorSteps/easypark/internal/adapters/datastore"
	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	mocks "github.com/IgorSteps/easypark/mocks/adapters/datastore"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestParkingRequestPostgresRepository_CreateParkingRequest_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockDatastore := &mocks.Datastore{}
	repository := datastore.NewParkingRequestPostgresRepository(mockDatastore, testLogger)
	testCtx := context.Background()
	testParkingRequest := &entities.ParkingRequest{
		ID:             uuid.New(),
		UserID:         uuid.New(),
		ParkingSpaceID: nil,
		Destination:    "goom",
		StartTime:      time.Now(),
		EndTime:        time.Now().Add(5),
		Status:         entities.RequestStatusPending,
	}

	mockDatastore.EXPECT().WithContext(testCtx).Return(mockDatastore).Once()
	mockDatastore.EXPECT().Create(testParkingRequest).Return(mockDatastore).Once()
	mockDatastore.EXPECT().Error().Return(nil).Once()

	// --------
	// ACT
	// --------
	err := repository.CreateParkingRequest(testCtx, testParkingRequest)

	// --------
	// ASSERT
	// --------
	assert.Nil(t, err, "Error must be nil")
	assert.Equal(t, 0, len(hook.Entries), "Logger shouldn't log anything")
	mockDatastore.AssertExpectations(t)
}

func TestParkingRequestPostgresRepository_CreateParkingRequest_UnhappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockDatastore := &mocks.Datastore{}
	repository := datastore.NewParkingRequestPostgresRepository(mockDatastore, testLogger)
	testCtx := context.Background()
	testParkingRequest := &entities.ParkingRequest{
		ID:             uuid.New(),
		UserID:         uuid.New(),
		ParkingSpaceID: nil,
		Destination:    "goom",
		StartTime:      time.Now(),
		EndTime:        time.Now().Add(5),
		Status:         entities.RequestStatusPending,
	}
	testError := errors.New("boom")
	mockDatastore.EXPECT().WithContext(testCtx).Return(mockDatastore).Once()
	mockDatastore.EXPECT().Create(testParkingRequest).Return(mockDatastore).Once()
	mockDatastore.EXPECT().Error().Return(testError).Once()

	// --------
	// ACT
	// --------
	err := repository.CreateParkingRequest(testCtx, testParkingRequest)

	// --------
	// ASSERT
	// --------
	assert.IsType(t, &repositories.InternalError{}, err, "Returned error is of wrong type")
	assert.EqualError(t, err, "Internal error: failed to insert parking request into the database", "Error message is wrong")
	// Assert logger
	assert.Equal(t, 1, len(hook.Entries), "Logger should've output 1 error log")
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	assert.Equal(t, "failed to insert parking request into the database", hook.LastEntry().Message, "Errors message is not equal")
	assert.Equal(t, testError, hook.LastEntry().Data["error"], "Error in the logger is wrong")
	mockDatastore.AssertExpectations(t)
}

func TestParkingRequestPostgresRepository_GetAllParkingRequests_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockDatastore := &mocks.Datastore{}
	repository := datastore.NewParkingRequestPostgresRepository(mockDatastore, testLogger)
	testCtx := context.Background()
	testParkingRequests := []entities.ParkingRequest{
		{
			ID:             uuid.New(),
			UserID:         uuid.New(),
			ParkingSpaceID: nil,
			Destination:    "foo",
			StartTime:      time.Now(),
			EndTime:        time.Now().Add(5),
			Status:         entities.RequestStatusPending,
		},
		{
			ID:             uuid.New(),
			UserID:         uuid.New(),
			ParkingSpaceID: nil,
			Destination:    "boo",
			StartTime:      time.Now(),
			EndTime:        time.Now().Add(5),
			Status:         entities.RequestStatusPending,
		},
	}

	var requests []entities.ParkingRequest
	mockDatastore.EXPECT().WithContext(testCtx).Return(mockDatastore).Once()
	mockDatastore.EXPECT().FindAll(&requests).Return(mockDatastore).Once().Run(func(args mock.Arguments) {
		arg := args.Get(0).(*[]entities.ParkingRequest) // Get the first argument passed to FindAll()
		*arg = testParkingRequests                      // Set it to the expected park reqs
	})
	mockDatastore.EXPECT().Error().Return(nil).Once()

	// --------
	// ACT
	// --------
	actualParkingRequests, err := repository.GetAllParkingRequests(testCtx)

	// --------
	// ASSERT
	// --------
	assert.Nil(t, err, "Error must be nil")
	assert.Equal(t, testParkingRequests, actualParkingRequests, "Parking requests retunred do not equal expected")
	assert.Equal(t, 0, len(hook.Entries), "Logger shouldn't log anything")
	mockDatastore.AssertExpectations(t)
}

func TestParkingRequestPostgresRepository_GetAllParkingRequests_UnhappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockDatastore := &mocks.Datastore{}
	repository := datastore.NewParkingRequestPostgresRepository(mockDatastore, testLogger)
	testCtx := context.Background()
	testError := errors.New("boom")

	var requests []entities.ParkingRequest
	mockDatastore.EXPECT().WithContext(testCtx).Return(mockDatastore).Once()
	mockDatastore.EXPECT().FindAll(&requests).Return(mockDatastore).Once()
	mockDatastore.EXPECT().Error().Return(testError).Once()

	// --------
	// ACT
	// --------
	actualParkingRequests, err := repository.GetAllParkingRequests(testCtx)

	// --------
	// ASSERT
	// --------
	assert.Empty(t, actualParkingRequests, "Parking requests retunred must be empty")
	assert.IsType(t, &repositories.InternalError{}, err, "Error type is wrong")
	assert.EqualError(t, err, "Internal error: failed to query for all parking requests in the database")

	// Assert logger
	assert.Equal(t, 1, len(hook.Entries), "Logger should log the error")
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level, "Log level must be Error")
	assert.Equal(t, "failed to query for all parking requests in the database", hook.LastEntry().Message, "Log message is wrong")
	assert.Equal(t, testError, hook.LastEntry().Data["error"], "Error in the logger is wrong")

	mockDatastore.AssertExpectations(t)
}

func TestParkingRequestPostgresRepository_GetAllParkingRequestsForUser_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockDatastore := &mocks.Datastore{}
	repository := datastore.NewParkingRequestPostgresRepository(mockDatastore, testLogger)
	testCtx := context.Background()
	testUserID := uuid.New()
	testParkingRequests := []entities.ParkingRequest{
		{
			ID:             uuid.New(),
			UserID:         testUserID,
			ParkingSpaceID: nil,
			Destination:    "foo",
			StartTime:      time.Now(),
			EndTime:        time.Now().Add(5),
			Status:         entities.RequestStatusPending,
		},
		{
			ID:             uuid.New(),
			UserID:         testUserID,
			ParkingSpaceID: nil,
			Destination:    "boo",
			StartTime:      time.Now(),
			EndTime:        time.Now().Add(5),
			Status:         entities.RequestStatusPending,
		},
	}

	var requests []entities.ParkingRequest
	mockDatastore.EXPECT().WithContext(testCtx).Return(mockDatastore).Once()
	mockDatastore.EXPECT().Where("user_id = ?", testUserID).Return(mockDatastore).Once()
	mockDatastore.EXPECT().FindAll(&requests).Return(mockDatastore).Once().Run(func(args mock.Arguments) {
		arg := args.Get(0).(*[]entities.ParkingRequest) // Get the first argument passed to FindAll()
		*arg = testParkingRequests                      // Set it to the expected park reqs
	})
	mockDatastore.EXPECT().Error().Return(nil).Once()

	// --------
	// ACT
	// --------
	actualParkingRequests, err := repository.GetAllParkingRequestsForUser(testCtx, testUserID)

	// --------
	// ASSERT
	// --------
	assert.Nil(t, err, "Error must be nil")
	assert.Equal(t, testParkingRequests, actualParkingRequests, "Parking requests retunred do not equal expected")
	assert.Equal(t, 0, len(hook.Entries), "Logger shouldn't log anything")
	mockDatastore.AssertExpectations(t)
}

func TestParkingRequestPostgresRepository_GetAllParkingRequestsForUser_UnhappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockDatastore := &mocks.Datastore{}
	repository := datastore.NewParkingRequestPostgresRepository(mockDatastore, testLogger)
	testCtx := context.Background()
	testUserID := uuid.New()
	testError := errors.New("boom")

	var requests []entities.ParkingRequest
	mockDatastore.EXPECT().WithContext(testCtx).Return(mockDatastore).Once()
	mockDatastore.EXPECT().Where("user_id = ?", testUserID).Return(mockDatastore).Once()
	mockDatastore.EXPECT().FindAll(&requests).Return(mockDatastore).Once()
	mockDatastore.EXPECT().Error().Return(testError).Once()

	// --------
	// ACT
	// --------
	actualParkingRequests, err := repository.GetAllParkingRequestsForUser(testCtx, testUserID)

	// --------
	// ASSERT
	// --------
	assert.NotNil(t, err, "Error must not be nil")
	assert.IsType(t, &repositories.InternalError{}, err, "Error returned is of wrong type")
	assert.EqualError(t, err, "Internal error: failed to query for all parking requests in the database for particular user")
	assert.Empty(t, actualParkingRequests, "Parking requests retunred must be empty")
	mockDatastore.AssertExpectations(t)

	// Assert logger
	assert.Equal(t, 1, len(hook.Entries), "Logger should log the error")
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level, "Log level must be Error")
	assert.Equal(t, "failed to query for all parking requests in the database for particular user", hook.LastEntry().Message, "Log message is wrong")
	assert.Equal(t, testError, hook.LastEntry().Data["error"], "Error in the logger is wrong")
}

func TestParkingRequestPostgresRepository_GetParkingRequestByID_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockDatastore := &mocks.Datastore{}
	repository := datastore.NewParkingRequestPostgresRepository(mockDatastore, testLogger)
	testCtx := context.Background()
	testParkingRequestID := uuid.New()
	testParkingRequest := entities.ParkingRequest{
		ID:             testParkingRequestID,
		UserID:         uuid.New(),
		ParkingSpaceID: nil,
		Destination:    "foo",
		StartTime:      time.Now(),
		EndTime:        time.Now().Add(5),
		Status:         entities.RequestStatusPending,
	}

	var request entities.ParkingRequest
	mockDatastore.EXPECT().WithContext(testCtx).Return(mockDatastore).Once()
	mockDatastore.EXPECT().First(&request, "id = ?", testParkingRequestID).Return(mockDatastore).Once().Run(func(args mock.Arguments) {
		arg := args.Get(0).(*entities.ParkingRequest) // Get the first argument passed to FindAll()
		*arg = testParkingRequest                     // Set it to the expected park reqs
	})

	mockDatastore.EXPECT().Error().Return(nil).Once()

	// --------
	// ACT
	// --------
	actualParkingRequest, err := repository.GetParkingRequestByID(testCtx, testParkingRequestID)

	// --------
	// ASSERT
	// --------
	assert.Nil(t, err, "Error must be nil")
	assert.Equal(t, testParkingRequest, actualParkingRequest, "Parking request retunred does not equal expected")
	assert.Equal(t, 0, len(hook.Entries), "Logger shouldn't log anything")
	mockDatastore.AssertExpectations(t)
}

func TestParkingRequestPostgresRepository_GetParkingRequestByID_UnhappyPath_NotFound(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockDatastore := &mocks.Datastore{}
	repository := datastore.NewParkingRequestPostgresRepository(mockDatastore, testLogger)
	testCtx := context.Background()
	testParkingRequestID := uuid.New()

	var request entities.ParkingRequest
	mockDatastore.EXPECT().WithContext(testCtx).Return(mockDatastore).Once()
	mockDatastore.EXPECT().First(&request, "id = ?", testParkingRequestID).Return(mockDatastore).Once()
	mockDatastore.EXPECT().Error().Return(gorm.ErrRecordNotFound).Once()

	// --------
	// ACT
	// --------
	actualParkingRequest, err := repository.GetParkingRequestByID(testCtx, testParkingRequestID)

	// --------
	// ASSERT
	// --------
	assert.NotNil(t, err, "Error must not be nil")
	assert.IsType(t, &repositories.NotFoundError{}, err, "Error returned is of wrong type")
	assert.EqualError(t, err, fmt.Sprintf("Resource '%s' not found", testParkingRequestID), "Errors are not equal")
	assert.Empty(t, actualParkingRequest, "Parking request returned must be empty")
	mockDatastore.AssertExpectations(t)

	// Assert logger
	assert.Equal(t, 1, len(hook.Entries), "Logger should log the error")
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level, "Log level must be Error")
	assert.Equal(t, "failed to find parking request with given id in the database", hook.LastEntry().Message, "Log message is wrong")
	assert.Equal(t, testParkingRequestID, hook.LastEntry().Data["parking request id"], "ID field in the logger is wrong")
}

func TestParkingRequestPostgresRepository_GetParkingRequestByID_UnhappyPath_InternalError(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockDatastore := &mocks.Datastore{}
	repository := datastore.NewParkingRequestPostgresRepository(mockDatastore, testLogger)
	testCtx := context.Background()
	testParkingRequestID := uuid.New()
	testError := errors.New("boom")

	var request entities.ParkingRequest
	mockDatastore.EXPECT().WithContext(testCtx).Return(mockDatastore).Once()
	mockDatastore.EXPECT().First(&request, "id = ?", testParkingRequestID).Return(mockDatastore).Once()
	mockDatastore.EXPECT().Error().Return(testError).Once()

	// --------
	// ACT
	// --------
	actualParkingRequest, err := repository.GetParkingRequestByID(testCtx, testParkingRequestID)

	// --------
	// ASSERT
	// --------
	assert.NotNil(t, err, "Error must not be nil")
	assert.IsType(t, &repositories.InternalError{}, err, "Error returned is of wrong type")
	assert.EqualError(t, err, "Internal error: failed to query for parking request in the database", "Errors are not equal")
	assert.Empty(t, actualParkingRequest, "Parking request returned must be empty")
	mockDatastore.AssertExpectations(t)

	// Assert logger
	assert.Equal(t, 1, len(hook.Entries), "Logger should log the error")
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level, "Log level must be Error")
	assert.Equal(t, "failed to query for parking request in the database", hook.LastEntry().Message, "Log message is wrong")
	assert.Equal(t, testError, hook.LastEntry().Data["error"], "Error field in the logger is wrong")
}
