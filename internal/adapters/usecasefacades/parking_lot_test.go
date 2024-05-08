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

type ParkingLotFacadeTestSuite struct {
	creator      *mocks.ParkingLotCreator
	getter       *mocks.ParkingLotGetter
	deleter      *mocks.ParkingLotDeleter
	singleGetter *mocks.ParkingLotSingleGetter
}

func newParkingLotFacadeTestSuite() *ParkingLotFacadeTestSuite {
	return &ParkingLotFacadeTestSuite{
		creator:      &mocks.ParkingLotCreator{},
		getter:       &mocks.ParkingLotGetter{},
		deleter:      &mocks.ParkingLotDeleter{},
		singleGetter: &mocks.ParkingLotSingleGetter{},
	}
}

func TestParkingLotFacade_Create(t *testing.T) {
	// ---------
	// ASSEMBLE
	// ---------
	s := newParkingLotFacadeTestSuite()
	facade := usecasefacades.NewParkingLotFacade(
		s.creator,
		s.getter,
		s.deleter,
		s.singleGetter,
	)

	testCtx := context.Background()
	lot := entities.ParkingLot{
		ID: uuid.New(),
	}
	s.creator.EXPECT().Execute(testCtx, "b", 1).Return(lot, nil).Once()

	// ---------
	// ACT
	// ---------
	actualLot, err := facade.CreateParkingLot(testCtx, "b", 1)

	// ---------
	// ASSERT
	// ---------
	assert.NoError(t, err)
	assert.Equal(t, lot, actualLot)
}

func TestParkingLotFacade_GetAll(t *testing.T) {
	// ---------
	// ASSEMBLE
	// ---------
	s := newParkingLotFacadeTestSuite()
	facade := usecasefacades.NewParkingLotFacade(
		s.creator,
		s.getter,
		s.deleter,
		s.singleGetter,
	)

	testCtx := context.Background()
	lots := []entities.ParkingLot{
		{ID: uuid.New()},
	}
	s.getter.EXPECT().Execute(testCtx).Return(lots, nil).Once()

	// ---------
	// ACT
	// ---------
	actualLot, err := facade.GetAllParkingLots(testCtx)

	// ---------
	// ASSERT
	// ---------
	assert.NoError(t, err)
	assert.Equal(t, len(actualLot), 1)
}

func TestParkingLotFacade_Delete(t *testing.T) {
	// ---------
	// ASSEMBLE
	// ---------
	s := newParkingLotFacadeTestSuite()
	facade := usecasefacades.NewParkingLotFacade(
		s.creator,
		s.getter,
		s.deleter,
		s.singleGetter,
	)

	testCtx := context.Background()
	testID := uuid.New()

	s.deleter.EXPECT().Execute(testCtx, testID).Return(nil).Once()

	// ---------
	// ACT
	// ---------
	err := facade.DeleteParkingLot(testCtx, testID)

	// ---------
	// ASSERT
	// ---------
	assert.NoError(t, err)
}
