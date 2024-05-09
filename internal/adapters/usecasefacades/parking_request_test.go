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

type ParkingRequestFacadeTestSuite struct {
	creator                *mocks.ParkingRequestCreator
	allGetter              *mocks.ParkingRequestsAllGetter
	driversGetter          *mocks.ParkingRequestDriversGetter
	statusUpdate           *mocks.ParkingRequestStatusUpdater
	spaceAssigner          *mocks.ParkingRequestSpaceAssigner
	automaticSpaceAssigner *mocks.ParkingRequestAutomaticSpaceAssigner
	deassigner             *mocks.ParkingRequestSpaceDeassigner
}

func newParkingRequestFacadeTestSuite() *ParkingRequestFacadeTestSuite {
	return &ParkingRequestFacadeTestSuite{
		creator:                &mocks.ParkingRequestCreator{},
		allGetter:              &mocks.ParkingRequestsAllGetter{},
		driversGetter:          &mocks.ParkingRequestDriversGetter{},
		statusUpdate:           &mocks.ParkingRequestStatusUpdater{},
		spaceAssigner:          &mocks.ParkingRequestSpaceAssigner{},
		automaticSpaceAssigner: &mocks.ParkingRequestAutomaticSpaceAssigner{},
		deassigner:             &mocks.ParkingRequestSpaceDeassigner{},
	}
}

func TestParkingRequestFacade_Create(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	s := newParkingRequestFacadeTestSuite()
	facade := usecasefacades.NewParkingRequestFacade(
		s.creator,
		s.statusUpdate,
		s.spaceAssigner,
		s.allGetter,
		s.driversGetter,
		s.automaticSpaceAssigner,
		s.deassigner,
	)

	parkreq := &entities.ParkingRequest{
		ID: uuid.New(),
	}
	testCtx := context.Background()
	s.creator.EXPECT().Execute(testCtx, parkreq).Return(parkreq, nil).Once()

	// ------
	// ACT
	// ------
	req, err := facade.CreateParkingRequest(testCtx, parkreq)

	// ------
	// ASSERT
	// ------
	assert.NoError(t, err)
	assert.Equal(t, parkreq, req)
}

func TestParkingRequestFacade_UpdateStatus(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	s := newParkingRequestFacadeTestSuite()
	facade := usecasefacades.NewParkingRequestFacade(
		s.creator,
		s.statusUpdate,
		s.spaceAssigner,
		s.allGetter,
		s.driversGetter,
		s.automaticSpaceAssigner,
		s.deassigner,
	)

	id := uuid.New()
	status := "aaaa"
	testCtx := context.Background()
	s.statusUpdate.EXPECT().Execute(testCtx, id, status).Return(nil).Once()

	// ------
	// ACT
	// ------
	err := facade.UpdateParkingRequestStatus(testCtx, id, status)

	// ------
	// ASSERT
	// ------
	assert.NoError(t, err)
}

func TestParkingRequestFacade_AssignSpace(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	s := newParkingRequestFacadeTestSuite()
	facade := usecasefacades.NewParkingRequestFacade(
		s.creator,
		s.statusUpdate,
		s.spaceAssigner,
		s.allGetter,
		s.driversGetter,
		s.automaticSpaceAssigner,
		s.deassigner,
	)

	id := uuid.New()
	anotherId := uuid.New()
	testCtx := context.Background()
	s.spaceAssigner.EXPECT().Execute(testCtx, id, anotherId).Return(nil).Once()

	// ------
	// ACT
	// ------
	err := facade.AssignParkingSpace(testCtx, id, anotherId)

	// ------
	// ASSERT
	// ------
	assert.NoError(t, err)
}

func TestParkingRequestFacade_GetAll(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	s := newParkingRequestFacadeTestSuite()
	facade := usecasefacades.NewParkingRequestFacade(
		s.creator,
		s.statusUpdate,
		s.spaceAssigner,
		s.allGetter,
		s.driversGetter,
		s.automaticSpaceAssigner,
		s.deassigner,
	)
	reqs := []entities.ParkingRequest{
		{ID: uuid.New()},
	}
	testCtx := context.Background()
	s.allGetter.EXPECT().Execute(testCtx).Return(reqs, nil).Once()

	// ------
	// ACT
	// ------
	actualReqs, err := facade.GetAllParkingRequests(testCtx)

	// ------
	// ASSERT
	// ------
	assert.NoError(t, err)
	assert.Equal(t, len(actualReqs), 1)
}

func TestParkingRequestFacade_GetDrivers(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	s := newParkingRequestFacadeTestSuite()
	facade := usecasefacades.NewParkingRequestFacade(
		s.creator,
		s.statusUpdate,
		s.spaceAssigner,
		s.allGetter,
		s.driversGetter,
		s.automaticSpaceAssigner,
		s.deassigner,
	)
	reqs := []entities.ParkingRequest{
		{ID: uuid.New()},
	}
	testID := uuid.New()
	testCtx := context.Background()
	s.driversGetter.EXPECT().Execute(testCtx, testID).Return(reqs, nil).Once()

	// ------
	// ACT
	// ------
	actualReqs, err := facade.GetDriversParkingRequests(testCtx, testID)

	// ------
	// ASSERT
	// ------
	assert.NoError(t, err)
	assert.Equal(t, len(actualReqs), 1)
}
