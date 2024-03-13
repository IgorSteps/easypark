package routes_test

import (
	"net/http"
	"testing"

	"github.com/IgorSteps/easypark/internal/adapters/rest/routes"
	mocks "github.com/IgorSteps/easypark/mocks/adapters/rest/routes"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestRoutes_NewRouter_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	mockHandlerFactory := &mocks.HandlerFactory{}
	logger := logrus.New()
	// Test handler to return from the factory.
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	mockHandlerFactory.EXPECT().UserCreate().Return(testHandler).Once()
	mockHandlerFactory.EXPECT().UserAuthorise().Return(testHandler).Once()

	// --------
	// ACT
	// --------
	r := routes.NewRouter(mockHandlerFactory, logger)

	// --------
	// ASSERT
	// --------
	assert.NotNil(t, r, "Router must not be nil")
	mockHandlerFactory.AssertExpectations(t)
}
