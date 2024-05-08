package usecases_test

import (
	"context"
	"testing"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	usecases "github.com/IgorSteps/easypark/internal/usecases/parkinglot"
	mocks "github.com/IgorSteps/easypark/mocks/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestGetSingleParkingLot_Execute(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	mockRepo := &mocks.ParkingLotRepository{}
	testLogger, _ := test.NewNullLogger()
	usecase := usecases.NewGetSingleParkingLot(testLogger, mockRepo)

	testID := uuid.New()
	ctx := context.Background()
	lot := &entities.ParkingLot{
		ID: uuid.New(),
	}
	mockRepo.EXPECT().GetSingle(ctx, testID).Return(lot, nil).Once()

	// --------
	// ACT
	// --------
	actualLot, err := usecase.Execute(ctx, testID)

	// --------
	// ASSERT
	// --------
	assert.NoError(t, err)
	assert.Equal(t, lot, actualLot)
}
