package usecases_test

import (
	"context"
	"errors"
	"testing"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	usecases "github.com/IgorSteps/easypark/internal/usecases/parkinglot"
	mocks "github.com/IgorSteps/easypark/mocks/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestGetAllParkingLots_Execute_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testCtx := context.Background()
	testLogger, _ := test.NewNullLogger()
	mockRepo := &mocks.ParkingLotRepository{}
	usecase := usecases.NewGetAllParkingLots(testLogger, mockRepo)

	testLots := []entities.ParkingLot{
		{
			ID:            uuid.New(),
			Name:          "blol",
			Capacity:      10,
			ParkingSpaces: nil,
		},
		{
			ID:            uuid.New(),
			Name:          "blol",
			Capacity:      10,
			ParkingSpaces: nil,
		},
		{
			ID:            uuid.New(),
			Name:          "blol",
			Capacity:      10,
			ParkingSpaces: nil,
		},
	}

	mockRepo.EXPECT().GetAllParkingLots(testCtx).Return(testLots, nil).Once()

	// --------
	// ACT
	// --------
	lots, err := usecase.Execute(testCtx)

	// --------
	// ASSERT
	// --------
	assert.Nil(t, err, "Error must be nil")
	assert.Equal(t, testLots, lots, "Wrong p`rking lots")
	mockRepo.AssertExpectations(t)
}

func TestGetAllParkingLots_Execute_UnhappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testCtx := context.Background()
	testLogger, _ := test.NewNullLogger()
	mockRepo := &mocks.ParkingLotRepository{}
	usecase := usecases.NewGetAllParkingLots(testLogger, mockRepo)

	testErr := errors.New("boom")
	mockRepo.EXPECT().GetAllParkingLots(testCtx).Return(nil, testErr).Once()

	// --------
	// ACT
	// --------
	lots, err := usecase.Execute(testCtx)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, testErr, err, "Errors must be equal")
	assert.Nil(t, lots, "Lots shuld be nil")
	mockRepo.AssertExpectations(t)
}
