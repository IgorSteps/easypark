package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/IgorSteps/easypark/internal/adapters/rest/handlers"
	"github.com/IgorSteps/easypark/internal/adapters/rest/models"
	"github.com/IgorSteps/easypark/internal/domain/entities"
	mocks "github.com/IgorSteps/easypark/mocks/adapters/rest/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestNotificationCreateHandler_ServeHTTP_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockFacade := &mocks.NotificationFacade{}
	handler := handlers.NewNotificationCreateHandler(testLogger, mockFacade)

	testCreateRequest := models.CreateNotificationRequest{
		ParkingRequestID: uuid.New(),
		ParkingSpaceID:   uuid.New(),
		Location:         "ldddd",
		NotificationType: 0,
	}

	requestBody, err := json.Marshal(testCreateRequest)
	assert.NoError(t, err, "Marshalling request to json must not return error")

	testID := uuid.New()

	// Because we are directly calling the handler in the test without going through a router that parses the URL parameters,
	// we have to manually insert the URL parameters into the request context.
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", testID.String())

	// Create our request context with the formatted chi context.
	reqCtx := context.WithValue(context.Background(), chi.RouteCtxKey, rctx)
	req, _ := http.NewRequestWithContext(reqCtx, "POST", "/drivers/"+testID.String()+"/notifications", bytes.NewBuffer(requestBody))
	rr := httptest.NewRecorder()

	notification := entities.Notification{}
	mockFacade.EXPECT().CreateNotification(
		reqCtx,
		testID,
		testCreateRequest.ParkingRequestID,
		testCreateRequest.ParkingSpaceID,
		testCreateRequest.Location,
		testCreateRequest.NotificationType,
	).Return(notification, nil).Once()

	// --------
	// ACT
	// --------
	handler.ServeHTTP(rr, req)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, http.StatusCreated, rr.Code, "Response codes don't match, should be 201 CREATED")
	mockFacade.AssertExpectations(t)
}
