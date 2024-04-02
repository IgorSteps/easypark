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
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	mocks "github.com/IgorSteps/easypark/mocks/adapters/rest/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestParkingRequestUpdateStatusHandler_ServeHTTP_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockFacade := &mocks.ParkingRequestFacade{}
	handler := handlers.NewParkingRequestStatusHandler(mockFacade, testLogger)

	testUpdareParkingRequestStatusRequest := models.UpdateParkingRequestStatusRequest{
		Status: "approved",
	}

	requestBody, _ := json.Marshal(testUpdareParkingRequestStatusRequest)
	testID := uuid.New()

	// Because we are directly calling the handler in the test without going through a router that parses the URL parameters,
	// we have to manually insert the URL parameters into the request context.
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", testID.String())

	// Create our request context with the formatted chi context.
	reqCtx := context.WithValue(context.Background(), chi.RouteCtxKey, rctx)
	req, _ := http.NewRequestWithContext(reqCtx, "POST", "/parking-requests/"+testID.String()+"/status", bytes.NewBuffer(requestBody))
	rr := httptest.NewRecorder()

	mockFacade.EXPECT().UpdateParkingRequestStatus(req.Context(), testID, "approved").Return(nil).Once()

	// --------
	// ACT
	// --------
	handler.ServeHTTP(rr, req)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, http.StatusOK, rr.Code, "Response codes don't match, should be 200")
	assert.Contains(t, rr.Body.String(), "successfully updated parking request status", "Reponse bodies don't match")
	mockFacade.AssertExpectations(t)
}

func TestParkingRequestUpdateStatusHandler_ServeHTTP_UnhappyPath_InvalidInput(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockFacade := &mocks.ParkingRequestFacade{}
	handler := handlers.NewParkingRequestStatusHandler(mockFacade, testLogger)

	testUpdareParkingRequestStatusRequest := models.UpdateParkingRequestStatusRequest{
		Status: "invalid",
	}

	requestBody, _ := json.Marshal(testUpdareParkingRequestStatusRequest)
	testID := uuid.New()

	// Because we are directly calling the handler in the test without going through a router that parses the URL parameters,
	// we have to manually insert the URL parameters into the request context.
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", testID.String())

	// Create our request context with the formatted chi context.
	reqCtx := context.WithValue(context.Background(), chi.RouteCtxKey, rctx)
	req, _ := http.NewRequestWithContext(reqCtx, "POST", "/parking-requests/"+testID.String()+"/status", bytes.NewBuffer(requestBody))
	rr := httptest.NewRecorder()

	testError := repositories.NewInvalidInputError("boom")
	mockFacade.EXPECT().UpdateParkingRequestStatus(req.Context(), testID, "invalid").Return(testError).Once()

	// --------
	// ACT
	// --------
	handler.ServeHTTP(rr, req)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, http.StatusBadRequest, rr.Code, "Response codes don't match, should be 400")
	assert.Contains(t, rr.Body.String(), "boom\n", "Reponse bodies don't match")
	mockFacade.AssertExpectations(t)
}

func TestParkingRequestUpdateStatusHandler_ServeHTTP_UnhappyPath_NotFound(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockFacade := &mocks.ParkingRequestFacade{}
	handler := handlers.NewParkingRequestStatusHandler(mockFacade, testLogger)

	testUpdareParkingRequestStatusRequest := models.UpdateParkingRequestStatusRequest{
		Status: "invalid",
	}

	requestBody, _ := json.Marshal(testUpdareParkingRequestStatusRequest)
	testID := uuid.New()

	// Because we are directly calling the handler in the test without going through a router that parses the URL parameters,
	// we have to manually insert the URL parameters into the request context.
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", testID.String())

	// Create our request context with the formatted chi context.
	reqCtx := context.WithValue(context.Background(), chi.RouteCtxKey, rctx)
	req, _ := http.NewRequestWithContext(reqCtx, "POST", "/parking-requests/"+testID.String()+"/status", bytes.NewBuffer(requestBody))
	rr := httptest.NewRecorder()

	testError := repositories.NewNotFoundError("boom")
	mockFacade.EXPECT().UpdateParkingRequestStatus(req.Context(), testID, "invalid").Return(testError).Once()

	// --------
	// ACT
	// --------
	handler.ServeHTTP(rr, req)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, http.StatusBadRequest, rr.Code, "Response codes don't match, should be 400")
	assert.Contains(t, rr.Body.String(), "Resource 'boom' not found\n", "Reponse bodies don't match")
	mockFacade.AssertExpectations(t)
}

func TestParkingRequestUpdateStatusHandler_ServeHTTP_UnhappyPath_Internal(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockFacade := &mocks.ParkingRequestFacade{}
	handler := handlers.NewParkingRequestStatusHandler(mockFacade, testLogger)

	testUpdareParkingRequestStatusRequest := models.UpdateParkingRequestStatusRequest{
		Status: "invalid",
	}

	requestBody, _ := json.Marshal(testUpdareParkingRequestStatusRequest)
	testID := uuid.New()

	// Because we are directly calling the handler in the test without going through a router that parses the URL parameters,
	// we have to manually insert the URL parameters into the request context.
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", testID.String())

	// Create our request context with the formatted chi context.
	reqCtx := context.WithValue(context.Background(), chi.RouteCtxKey, rctx)
	req, _ := http.NewRequestWithContext(reqCtx, "POST", "/parking-requests/"+testID.String()+"/status", bytes.NewBuffer(requestBody))
	rr := httptest.NewRecorder()

	testError := repositories.NewInternalError("boom")
	mockFacade.EXPECT().UpdateParkingRequestStatus(req.Context(), testID, "invalid").Return(testError).Once()

	// --------
	// ACT
	// --------
	handler.ServeHTTP(rr, req)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, http.StatusInternalServerError, rr.Code, "Response codes don't match, should be 500")
	assert.Contains(t, rr.Body.String(), "Internal error: boom\n", "Reponse bodies don't match")
	mockFacade.AssertExpectations(t)
}

func TestParkingRequestUpdateStatusHandler_ServeHTTP_UnhappyPath_IDFailedParsing(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockFacade := &mocks.ParkingRequestFacade{}
	handler := handlers.NewParkingRequestStatusHandler(mockFacade, testLogger)

	testUpdareParkingRequestStatusRequest := models.UpdateParkingRequestStatusRequest{
		Status: "invalid",
	}

	requestBody, _ := json.Marshal(testUpdareParkingRequestStatusRequest)
	testID := "incorrect"

	// Because we are directly calling the handler in the test without going through a router that parses the URL parameters,
	// we have to manually insert the URL parameters into the request context.
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", testID)

	// Create our request context with the formatted chi context.
	reqCtx := context.WithValue(context.Background(), chi.RouteCtxKey, rctx)
	req, _ := http.NewRequestWithContext(reqCtx, "POST", "/parking-requests/"+testID+"/status", bytes.NewBuffer(requestBody))
	rr := httptest.NewRecorder()

	// --------
	// ACT
	// --------
	handler.ServeHTTP(rr, req)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, http.StatusBadRequest, rr.Code, "Response codes don't match, should be 400")
	assert.Contains(t, rr.Body.String(), "invalid parking request id", "Reponse bodies don't match")
	mockFacade.AssertExpectations(t)
}
