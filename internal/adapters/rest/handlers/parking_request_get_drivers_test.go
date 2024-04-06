package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/IgorSteps/easypark/internal/adapters/rest/handlers"
	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	mocks "github.com/IgorSteps/easypark/mocks/adapters/rest/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestDriversParkingRequestsGetHandler_ServeHTTP_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockFacade := &mocks.ParkingRequestFacade{}

	// Don't set times here because they break tests.
	testParkRequests := []entities.ParkingRequest{
		{
			ID:                      uuid.New(),
			UserID:                  uuid.New(),
			DestinationParkingLotID: uuid.New(),
			Status:                  entities.RequestStatusApproved,
		},
		{
			ID:                      uuid.New(),
			UserID:                  uuid.New(),
			DestinationParkingLotID: uuid.New(),
			Status:                  entities.RequestStatusApproved,
		},
	}

	handler := handlers.NewDriversParkingRequestsGetHandler(testLogger, mockFacade)
	testID := uuid.New()

	// Because we are directly calling the handler in the test without going through a router that parses the URL parameters,
	// we have to manually insert the URL parameters into the request context.
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", testID.String())

	// Create our request context with the formatted chi context.
	reqCtx := context.WithValue(context.Background(), chi.RouteCtxKey, rctx)
	req, _ := http.NewRequestWithContext(reqCtx, "GET", "/drivers/"+testID.String()+"/parking-requests", bytes.NewBuffer(nil))
	rr := httptest.NewRecorder()

	mockFacade.EXPECT().GetDriversParkingRequests(reqCtx, testID).Return(testParkRequests, nil).Once()

	// --------
	// ACT
	// --------
	handler.ServeHTTP(rr, req)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, http.StatusOK, rr.Code, "Response codes don't match, should be 200 OK")

	// Unmarshal response body into slice of ParkingRequests.
	var actualParkReqs []entities.ParkingRequest
	err := json.Unmarshal(rr.Body.Bytes(), &actualParkReqs)
	assert.NoError(t, err, "Must have no error unmarshaling response body")
	assert.Equal(t, testParkRequests, actualParkReqs, "Parking requests don't match")

	mockFacade.AssertExpectations(t)
}

func TestDriversParkingRequestsGetHandler_ServeHTTP_UnhappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockFacade := &mocks.ParkingRequestFacade{}

	// Don't set times here because they break tests.
	testParkRequests := []entities.ParkingRequest{}

	handler := handlers.NewDriversParkingRequestsGetHandler(testLogger, mockFacade)
	testID := uuid.New()

	// Because we are directly calling the handler in the test without going through a router that parses the URL parameters,
	// we have to manually insert the URL parameters into the request context.
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", testID.String())

	// Create our request context with the formatted chi context.
	reqCtx := context.WithValue(context.Background(), chi.RouteCtxKey, rctx)
	req, _ := http.NewRequestWithContext(reqCtx, "GET", "/drivers/"+testID.String()+"/parking-requests", bytes.NewBuffer(nil))
	rr := httptest.NewRecorder()

	mockFacade.EXPECT().GetDriversParkingRequests(reqCtx, testID).Return(testParkRequests, repositories.NewInternalError("boom")).Once()

	// --------
	// ACT
	// --------
	handler.ServeHTTP(rr, req)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, http.StatusInternalServerError, rr.Code, "Response codes don't match, should be 500")
	assert.Equal(t, "Internal error: boom\n", rr.Body.String(), "Response body is wrong")
	mockFacade.AssertExpectations(t)
}
