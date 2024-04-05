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

func TestParkingRequestUpdateSpaceHandler_ServeHTTP_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockFacade := &mocks.ParkingRequestFacade{}
	handler := handlers.NewParkingRequestSpaceHandler(mockFacade, testLogger)

	testReq := models.ParkingRequestSpaceUpdateRequest{
		ParkingSpaceID: uuid.New(),
	}

	requestBody, _ := json.Marshal(testReq)
	requestID := uuid.New()

	// Because we are directly calling the handler in the test without going through a router that parses the URL parameters,
	// we have to manually insert the URL parameters into the request context.
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", requestID.String())

	// Create our request context with the formatted chi context.
	reqCtx := context.WithValue(context.Background(), chi.RouteCtxKey, rctx)
	req, _ := http.NewRequestWithContext(reqCtx, "POST", "/parking-requests/"+requestID.String()+"/space", bytes.NewBuffer(requestBody))
	rr := httptest.NewRecorder()

	mockFacade.EXPECT().AssignParkingSpace(req.Context(), requestID, testReq.ParkingSpaceID).Return(nil).Once()

	// --------
	// ACT
	// --------
	handler.ServeHTTP(rr, req)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, http.StatusOK, rr.Code, "Response codes don't match, should be 200")
	assert.Contains(t, rr.Body.String(), "successfully assigned a space to a parking request", "Reponse bodies don't match")
	mockFacade.AssertExpectations(t)
}

func TestParkingRequestUpdateSpaceHandler_ServeHTTP_UnhappyPath_InvalidInputError(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockFacade := &mocks.ParkingRequestFacade{}
	handler := handlers.NewParkingRequestSpaceHandler(mockFacade, testLogger)

	testReq := models.ParkingRequestSpaceUpdateRequest{
		ParkingSpaceID: uuid.New(),
	}

	requestBody, _ := json.Marshal(testReq)
	requestID := uuid.New()

	// Because we are directly calling the handler in the test without going through a router that parses the URL parameters,
	// we have to manually insert the URL parameters into the request context.
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", requestID.String())

	// Create our request context with the formatted chi context.
	reqCtx := context.WithValue(context.Background(), chi.RouteCtxKey, rctx)
	req, _ := http.NewRequestWithContext(reqCtx, "POST", "/parking-requests/"+requestID.String()+"/space", bytes.NewBuffer(requestBody))
	rr := httptest.NewRecorder()

	testErr := repositories.NewInvalidInputError("boom")
	mockFacade.EXPECT().AssignParkingSpace(req.Context(), requestID, testReq.ParkingSpaceID).Return(testErr).Once()

	// --------
	// ACT
	// --------
	handler.ServeHTTP(rr, req)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, http.StatusBadRequest, rr.Code, "Response codes don't match, should be 400")
	assert.Contains(t, rr.Body.String(), testErr.Error(), "Reponse bodies don't match")
	mockFacade.AssertExpectations(t)
}

func TestParkingRequestUpdateSpaceHandler_ServeHTTP_UnhappyPath_NotFoundError(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockFacade := &mocks.ParkingRequestFacade{}
	handler := handlers.NewParkingRequestSpaceHandler(mockFacade, testLogger)

	testReq := models.ParkingRequestSpaceUpdateRequest{
		ParkingSpaceID: uuid.New(),
	}

	requestBody, _ := json.Marshal(testReq)
	requestID := uuid.New()

	// Because we are directly calling the handler in the test without going through a router that parses the URL parameters,
	// we have to manually insert the URL parameters into the request context.
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", requestID.String())

	// Create our request context with the formatted chi context.
	reqCtx := context.WithValue(context.Background(), chi.RouteCtxKey, rctx)
	req, _ := http.NewRequestWithContext(reqCtx, "POST", "/parking-requests/"+requestID.String()+"/space", bytes.NewBuffer(requestBody))
	rr := httptest.NewRecorder()

	testErr := repositories.NewNotFoundError("boom")
	mockFacade.EXPECT().AssignParkingSpace(req.Context(), requestID, testReq.ParkingSpaceID).Return(testErr).Once()

	// --------
	// ACT
	// --------
	handler.ServeHTTP(rr, req)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, http.StatusBadRequest, rr.Code, "Response codes don't match, should be 400")
	assert.Contains(t, rr.Body.String(), testErr.Error(), "Reponse bodies don't match")
	mockFacade.AssertExpectations(t)
}

func TestParkingRequestUpdateSpaceHandler_ServeHTTP_UnhappyPath_InternalError(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockFacade := &mocks.ParkingRequestFacade{}
	handler := handlers.NewParkingRequestSpaceHandler(mockFacade, testLogger)

	testReq := models.ParkingRequestSpaceUpdateRequest{
		ParkingSpaceID: uuid.New(),
	}

	requestBody, _ := json.Marshal(testReq)
	requestID := uuid.New()

	// Because we are directly calling the handler in the test without going through a router that parses the URL parameters,
	// we have to manually insert the URL parameters into the request context.
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", requestID.String())

	// Create our request context with the formatted chi context.
	reqCtx := context.WithValue(context.Background(), chi.RouteCtxKey, rctx)
	req, _ := http.NewRequestWithContext(reqCtx, "POST", "/parking-requests/"+requestID.String()+"/space", bytes.NewBuffer(requestBody))
	rr := httptest.NewRecorder()

	testErr := repositories.NewInternalError("boom")
	mockFacade.EXPECT().AssignParkingSpace(req.Context(), requestID, testReq.ParkingSpaceID).Return(testErr).Once()

	// --------
	// ACT
	// --------
	handler.ServeHTTP(rr, req)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, http.StatusInternalServerError, rr.Code, "Response codes don't match, should be 500")
	assert.Contains(t, rr.Body.String(), testErr.Error(), "Reponse bodies don't match")
	mockFacade.AssertExpectations(t)
}
