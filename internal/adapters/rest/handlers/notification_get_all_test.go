package handlers_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/IgorSteps/easypark/internal/adapters/rest/handlers"
	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	mocks "github.com/IgorSteps/easypark/mocks/adapters/rest/handlers"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestGetAllNotificationHandler_ServeHTTP_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockFacade := &mocks.NotificationFacade{}
	handler := handlers.NewNotificationGetAllHandler(testLogger, mockFacade)

	textCtx := context.Background()
	req, _ := http.NewRequestWithContext(textCtx, "GET", "/notifications", nil)
	rr := httptest.NewRecorder()

	testNotif := []entities.Notification{
		{
			ID: uuid.New(),
		},
		{
			ID: uuid.New(),
		},
	}

	mockFacade.EXPECT().GetAllNotifications(req.Context()).Return(testNotif, nil).Once()

	// --------
	// ACT
	// --------
	handler.ServeHTTP(rr, req)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, http.StatusOK, rr.Code, "Response codes don't match, should be 200")

	// Unmarshal response body into notification.
	var actualAlerts []entities.Notification
	err := json.Unmarshal(rr.Body.Bytes(), &actualAlerts)
	assert.NoError(t, err, "Must have no error unmarshaling response body")
	assert.Equal(t, testNotif, actualAlerts, "Notification don't match")

	mockFacade.AssertExpectations(t)
}

func TestGetAllNotificationHandler_ServeHTTP_UnhappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockFacade := &mocks.NotificationFacade{}
	handler := handlers.NewNotificationGetAllHandler(testLogger, mockFacade)

	textCtx := context.Background()
	req, _ := http.NewRequestWithContext(textCtx, "GET", "/notifications", nil)
	rr := httptest.NewRecorder()

	testNotif := []entities.Notification{}
	testError := repositories.NewInternalError("boom")
	mockFacade.EXPECT().GetAllNotifications(req.Context()).Return(testNotif, testError).Once()

	// --------
	// ACT
	// --------
	handler.ServeHTTP(rr, req)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, http.StatusInternalServerError, rr.Code, "Response codes don't match, should be 500")

	assert.Equal(t, rr.Body.String(), "Internal error: boom\n", "Errors don't match")

	mockFacade.AssertExpectations(t)
}
