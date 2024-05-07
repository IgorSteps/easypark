package usecasefacades_test

import (
	"context"
	"testing"

	"github.com/IgorSteps/easypark/internal/adapters/usecasefacades"
	"github.com/IgorSteps/easypark/internal/domain/entities"
	mocks "github.com/IgorSteps/easypark/mocks/adapters/usecasefacades"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestParkingSpaceFacade_UpdateParkingSpaceStatus(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	mockParkingSpaceStatusUpdater := &mocks.ParkingSpaceStatusUpdater{}
	mockParkingSpaceGetter := &mocks.ParkingSpaceGetter{}
	facade := usecasefacades.NewParkingSpaceFacade(mockParkingSpaceStatusUpdater, mockParkingSpaceGetter)

	ctx := context.Background()
	id := uuid.New()
	status := "aaaaa"
	expectedSpace := entities.ParkingSpace{
		ID: uuid.New(),
	}
	mockParkingSpaceStatusUpdater.EXPECT().Execute(ctx, id, status).Return(expectedSpace, nil).Once()

	// --------
	// ACT
	// --------
	space, err := facade.UpdateParkingSpaceStatus(ctx, id, status)

	// --------
	// ASSERT
	// --------
	assert.NoError(t, err)
	assert.Equal(t, expectedSpace, space)
}

func TestParkingSpaceFacade_GetSingleParkingSpace(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	mockParkingSpaceStatusUpdater := &mocks.ParkingSpaceStatusUpdater{}
	mockParkingSpaceGetter := &mocks.ParkingSpaceGetter{}
	facade := usecasefacades.NewParkingSpaceFacade(mockParkingSpaceStatusUpdater, mockParkingSpaceGetter)

	ctx := context.Background()
	id := uuid.New()
	expectedSpace := entities.ParkingSpace{
		ID: uuid.New(),
	}
	mockParkingSpaceGetter.EXPECT().Execute(ctx, id).Return(expectedSpace, nil).Once()

	// --------
	// ACT
	// --------
	space, err := facade.GetSingleParkingSpace(ctx, id)

	// --------
	// ASSERT
	// --------
	assert.NoError(t, err)
	assert.Equal(t, expectedSpace, space)
}
