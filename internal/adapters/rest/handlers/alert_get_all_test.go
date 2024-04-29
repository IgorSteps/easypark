package handlers_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/IgorSteps/easypark/internal/adapters/rest/handlers"
	"github.com/IgorSteps/easypark/internal/domain/entities"
	mocks "github.com/IgorSteps/easypark/mocks/adapters/rest/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestGetAllAlertHandler_ServeHTTP_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockFacade := &mocks.AlertFacade{}
	handler := handlers.NewAlertGetAllHandler(testLogger, mockFacade)

	testID := uuid.New()

	// Because we are directly calling the handler in the test without going through a router that parses the URL parameters,
	// we have to manually insert the URL parameters into the request context.
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", testID.String())

	// Create our request context with the formatted chi context.
	reqCtx := context.WithValue(context.Background(), chi.RouteCtxKey, rctx)
	req, _ := http.NewRequestWithContext(reqCtx, "GET", "/alerts/"+testID.String(), nil)
	rr := httptest.NewRecorder()

	testAlerts := []entities.Alert{
		{
			ID: uuid.New(),
		},
	}

	mockFacade.EXPECT().GetAllAlerts(req.Context()).Return(testAlerts, nil).Once()

	// --------
	// ACT
	// --------
	handler.ServeHTTP(rr, req)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, http.StatusOK, rr.Code, "Response codes don't match, should be 200")

	// Unmarshal response body into Alert.
	var actualAlert []entities.Alert
	err := json.Unmarshal(rr.Body.Bytes(), &actualAlert)
	assert.NoError(t, err, "Must have no error unmarshaling response body")
	assert.Equal(t, testAlerts, actualAlert, "Alerts don't match")

	mockFacade.AssertExpectations(t)
}
