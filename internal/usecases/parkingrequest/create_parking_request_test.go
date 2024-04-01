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
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestCreateParkingRequest_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testContext := context.Background()
	testParkingRequest := &entities.ParkingRequest{
		Destination: "comp",
		StartTime:   time.Now(),
		EndTime:     time.Now().Add(5),
	}
	testLogger, _ := test.NewNullLogger()
	mockRepo := &mocks.ParkingRequestRepository{}
	usecase := usecases.NewCreateParkingRequest(testLogger, mockRepo)

	mockRepo.EXPECT().CreateParkingRequest(testContext, testParkingRequest).Return(nil).Once()

	// --------
	// ACT
	// --------
	req, err := usecase.Execute(testContext, testParkingRequest)

	// --------
	// ASSERT
	// --------
	assert.Nil(t, err, "Error must be nil")
	assert.Equal(t, entities.RequestStatusPending, req.Status, "Request status must be pending")
	assert.NotNil(t, req.ID, "ID must be set")
	mockRepo.AssertExpectations(t)
}

func TestCreateParkingRequest_UnhappyPath_Invalid(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testContext := context.Background()
	testParkingRequest := &entities.ParkingRequest{
		Destination: "comp",
		StartTime:   time.Now().Add(50000), // make start time after end time to fail validation.
		EndTime:     time.Now(),
	}
	testLogger, _ := test.NewNullLogger()
	mockRepo := &mocks.ParkingRequestRepository{}
	usecase := usecases.NewCreateParkingRequest(testLogger, mockRepo)

	// --------
	// ACT
	// --------
	req, err := usecase.Execute(testContext, testParkingRequest)

	// --------
	// ASSERT
	// --------
	assert.Nil(t, req, "Returned request must be nil")
	assert.IsType(t, &repositories.InvalidInputError{}, err, "Error is of the wrong type")
	assert.EqualError(t, err, "start time cannot be after the end time", "Error message is wrong")
	mockRepo.AssertExpectations(t)
}

func TestCreateParkingRequest_UnhappyPath_RepositoryError(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testContext := context.Background()
	testParkingRequest := &entities.ParkingRequest{
		Destination: "comp",
		StartTime:   time.Now(),
		EndTime:     time.Now().Add(50),
	}
	testLogger, _ := test.NewNullLogger()
	mockRepo := &mocks.ParkingRequestRepository{}
	usecase := usecases.NewCreateParkingRequest(testLogger, mockRepo)

	testError := errors.New("boom")
	mockRepo.EXPECT().CreateParkingRequest(testContext, testParkingRequest).Return(testError).Once()

	// --------
	// ACT
	// --------
	req, err := usecase.Execute(testContext, testParkingRequest)

	// --------
	// ASSERT
	// --------
	assert.Nil(t, req, "Returned request must be nil")
	assert.EqualError(t, err, "boom", "Error message is wrong")

	mockRepo.AssertExpectations(t)
}
