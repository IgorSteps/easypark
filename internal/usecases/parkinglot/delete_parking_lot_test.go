package usecases_test

import (
	"context"
	"errors"
	"testing"

	usecases "github.com/IgorSteps/easypark/internal/usecases/parkinglot"
	mocks "github.com/IgorSteps/easypark/mocks/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestDeteleParkingLot_Execute_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockRepo := &mocks.ParkingLotRepository{}
	usecase := usecases.NewDeleteParkingLot(testLogger, mockRepo)

	testCtx := context.Background()
	testID := uuid.New()
	mockRepo.EXPECT().DeleteParkingLot(testCtx, testID).Return(nil).Once()

	// --------
	// ACT
	// --------
	err := usecase.Execute(testCtx, testID)

	// --------
	// ASSERT
	// --------
	assert.Nil(t, err, "Error must be nil")
	mockRepo.AssertExpectations(t)
}

func TestDeteleParkingLot_Execute_UnhappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockRepo := &mocks.ParkingLotRepository{}
	usecase := usecases.NewDeleteParkingLot(testLogger, mockRepo)

	testCtx := context.Background()
	testID := uuid.New()
	testError := errors.New("boom")
	mockRepo.EXPECT().DeleteParkingLot(testCtx, testID).Return(testError).Once()

	// --------
	// ACT
	// --------
	err := usecase.Execute(testCtx, testID)

	// --------
	// ASSERT
	// --------
	assert.EqualError(t, err, "boom", "Error is wrong")
	mockRepo.AssertExpectations(t)
}
