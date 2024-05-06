package usecases_test

import (
	"context"
	"testing"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	usecases "github.com/IgorSteps/easypark/internal/usecases/parkingspace"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockParkingSpaceRepository struct {
	mock.Mock
}

func (m *MockParkingSpaceRepository) GetMany(ctx context.Context, query map[string]interface{}) ([]entities.ParkingSpace, error) {
	args := m.Called(ctx, query)
	return args.Get(0).([]entities.ParkingSpace), args.Error(1)
}

func (m *MockParkingSpaceRepository) Save(ctx context.Context, space *entities.ParkingSpace) error {
	args := m.Called(ctx, space)
	return args.Error(0)
}

func (m *MockParkingSpaceRepository) GetSingle(ctx context.Context, parkingSpaceID uuid.UUID) (entities.ParkingSpace, error) {
	args := m.Called(ctx, parkingSpaceID)
	return args.Get(0).(entities.ParkingSpace), args.Error(1)
}

func TestNewGetSingleParkingSpace(t *testing.T) {
	logger := logrus.New()
	repo := new(MockParkingSpaceRepository)

	usecase := usecases.NewGetSingleParkingSpace(logger, repo)
	assert.NotNil(t, usecase, "Expected GetSingleParkingSpace not to be nil")
}

func TestExecute_GetSingleParkingSpace(t *testing.T) {
	ctx := context.Background()
	parkingSpaceID := uuid.New()
	expectedParkingSpace := entities.ParkingSpace{
		ID:           parkingSpaceID,
		ParkingLotID: uuid.New(),
		Name:         "Location 123",
		Status:       entities.ParkingSpaceStatusAvailable,
	}

	logger := logrus.New()
	repo := new(MockParkingSpaceRepository)
	repo.On("GetSingle", ctx, parkingSpaceID).Return(expectedParkingSpace, nil) // Setting up the expectation

	usecase := usecases.NewGetSingleParkingSpace(logger, repo)

	space, err := usecase.Execute(ctx, parkingSpaceID)

	repo.AssertExpectations(t) // Verify that GetSingle was called as expected
	assert.NoError(t, err, "Unexpected error during execution")
	assert.Equal(t, expectedParkingSpace, space, "Returned parking space does not match expected")
}
