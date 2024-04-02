package routes_test

import (
	"net/http"
	"testing"

	"github.com/IgorSteps/easypark/internal/adapters/rest/routes"
	"github.com/IgorSteps/easypark/internal/domain/entities"
	mocks "github.com/IgorSteps/easypark/mocks/adapters/rest/routes"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRoutes_NewRouter_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	mockHandlerFactory := &mocks.HandlerFactory{}
	mockMiddleware := &mocks.Middleware{}
	logger := logrus.New()
	// Test handler to return from the factory.
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Define a simple pass-through middleware function for testing
	passThroughMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r) // Call the next handler in the chain
		})
	}

	mockHandlerFactory.EXPECT().DriverCreate().Return(testHandler).Once()
	mockHandlerFactory.EXPECT().UserAuthorise().Return(testHandler).Once()
	mockHandlerFactory.EXPECT().GetAllDrivers().Return(testHandler).Once()
	mockHandlerFactory.EXPECT().DriverBan().Return(testHandler).Once()
	mockHandlerFactory.EXPECT().ParkingRequestCreate().Return(testHandler).Once()
	mockHandlerFactory.EXPECT().ParkingRequestStatusUpdate().Return(testHandler).Once()

	// This middlware will get executed for very route invocation.
	mockMiddleware.EXPECT().Authorise(mock.AnythingOfType("http.HandlerFunc")).Return(testHandler).Times(5)
	mockMiddleware.EXPECT().CheckStatus(mock.AnythingOfType("http.HandlerFunc")).Return(testHandler).Twice()
	mockMiddleware.EXPECT().RequireRole(entities.RoleDriver).Return(passThroughMiddleware).Once()
	mockMiddleware.EXPECT().RequireRole(entities.RoleAdmin).Return(passThroughMiddleware).Once()

	// --------
	// ACT
	// --------
	r := routes.NewRouter(mockHandlerFactory, mockMiddleware, logger)

	// --------
	// ASSERT
	// --------
	assert.NotNil(t, r, "Router must not be nil")
	mockHandlerFactory.AssertExpectations(t)
	mockMiddleware.AssertExpectations(t)
}
