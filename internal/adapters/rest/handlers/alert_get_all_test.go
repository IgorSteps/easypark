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

	textCtx := context.Background()
	req, _ := http.NewRequestWithContext(textCtx, "GET", "/alerts", nil)
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
	var actualAlerts []entities.Alert
	err := json.Unmarshal(rr.Body.Bytes(), &actualAlerts)
	assert.NoError(t, err, "Must have no error unmarshaling response body")
	assert.Equal(t, testAlerts, actualAlerts, "Alerts don't match")

	mockFacade.AssertExpectations(t)
}
