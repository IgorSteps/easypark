package usecases_test

import (
	"context"
	"testing"

	usecases "github.com/IgorSteps/easypark/internal/usecases/parkinglot"
	mocks "github.com/IgorSteps/easypark/mocks/domain/repositories"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreatePakringLot_Execute_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockRepo := &mocks.ParkingLotRepository{}
	usecase := usecases.NewCreateParkingLot(testLogger, mockRepo)

	testCtx := context.Background()
	testName := "sci"
	testCapacity := 1

	mockRepo.EXPECT().CreateParkingLot(testCtx, mock.Anything).Return(nil).Once()

	// ----
	// ACT
	// ----
	parkingLot, err := usecase.Execute(testCtx, testName, testCapacity)

	// ------
	// ASSERT
	// ------
	assert.Nil(t, err, "Must not have an error")
	assert.Equal(t, testName, parkingLot.Name, "Parking lot name is wrong")
	assert.Equal(t, testCapacity, parkingLot.Capacity, "Parking lot capacity is wrong")
	assert.NotNil(t, parkingLot.ID, "ID must be set")
	assert.NotEmpty(t, parkingLot.ParkingSpaces, "Parking spaces slice must not be empty")
	mockRepo.AssertExpectations(t)
}
