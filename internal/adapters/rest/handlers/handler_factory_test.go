package handlers_test

import (
	"testing"

	"github.com/IgorSteps/easypark/internal/adapters/rest/handlers"
	mocks "github.com/IgorSteps/easypark/mocks/adapters/rest/handlers"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestHandlers_UserCreate_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger := logrus.New()
	mockUserFacade := &mocks.UserFacade{}
	mockParkingRequestFacade := &mocks.ParkingRequestFacade{}
	mockParkingLotFacade := &mocks.ParkingLotFacade{}
	mockFacade := handlers.NewFacade(mockUserFacade, mockParkingRequestFacade, mockParkingLotFacade)

	handlerFactory := handlers.NewHandlerFactory(testLogger, mockFacade)

	// ---
	// ACT
	// ---
	handler := handlerFactory.DriverCreate()

	// ------
	// ASSERT
	// ------
	assert.NotNil(t, handler, "Handler must not be nil")
}

func TestHandlers_UserLogin_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger := logrus.New()
	mockUserFacade := &mocks.UserFacade{}
	mockParkingRequestFacade := &mocks.ParkingRequestFacade{}
	mockParkingLotFacade := &mocks.ParkingLotFacade{}
	mockFacade := handlers.NewFacade(mockUserFacade, mockParkingRequestFacade, mockParkingLotFacade)

	handlerFactory := handlers.NewHandlerFactory(testLogger, mockFacade)

	// ---
	// ACT
	// ---
	handler := handlerFactory.UserAuthorise()

	// ------
	// ASSERT
	// ------
	assert.NotNil(t, handler, "Handler must not be nil")
}
