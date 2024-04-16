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
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	mocks "github.com/IgorSteps/easypark/mocks/adapters/rest/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestParkingSpaceUpdateStatusHandler_ServeHTTP_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockFacade := &mocks.ParkingSpaceFacade{}
	handler := handlers.NewParkingSpaceStatusHandler(mockFacade, testLogger)

	testModel := models.UpdateParkingSpaceStatus{
		Status: "available",
	}

	requestBody, _ := json.Marshal(testModel)
	testID := uuid.New()

	// Because we are directly calling the handler in the test without going through a router that parses the URL parameters,
	// we have to manually insert the URL parameters into the request context.
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", testID.String())

	// Create our request context with the formatted chi context.
	reqCtx := context.WithValue(context.Background(), chi.RouteCtxKey, rctx)
	req, _ := http.NewRequestWithContext(reqCtx, "PATCH", "/parking-space/"+testID.String()+"/status", bytes.NewBuffer(requestBody))
	rr := httptest.NewRecorder()

	testSpace := entities.ParkingSpace{
		ID:     uuid.New(),
		Status: entities.ParkingSpaceStatusAvailable,
	}
	mockFacade.EXPECT().UpdateParkingSpaceStatus(req.Context(), testID, "available").Return(testSpace, nil).Once()

	// --------
	// ACT
	// --------
	handler.ServeHTTP(rr, req)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, http.StatusOK, rr.Code, "Response codes don't match, should be 200")

	// Unmarshal response body into ParkingSpace entity.
	var targetEntity entities.ParkingSpace
	err := json.Unmarshal(rr.Body.Bytes(), &targetEntity)
	assert.NoError(t, err, "Must have no error unmarshaling response body")
	assert.Equal(t, testSpace.Status, targetEntity.Status, "Parking space statuses don't match")

	mockFacade.AssertExpectations(t)
}

func TestParkingSpaceUpdateStatusHandler_ServeHTTP_UnhappyPath_InvalidInputError(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockFacade := &mocks.ParkingSpaceFacade{}
	handler := handlers.NewParkingSpaceStatusHandler(mockFacade, testLogger)

	testModel := models.UpdateParkingSpaceStatus{
		Status: "available",
	}

	requestBody, _ := json.Marshal(testModel)
	testID := uuid.New()

	// Because we are directly calling the handler in the test without going through a router that parses the URL parameters,
	// we have to manually insert the URL parameters into the request context.
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", testID.String())

	// Create our request context with the formatted chi context.
	reqCtx := context.WithValue(context.Background(), chi.RouteCtxKey, rctx)
	req, _ := http.NewRequestWithContext(reqCtx, "PATCH", "/parking-space/"+testID.String()+"/status", bytes.NewBuffer(requestBody))
	rr := httptest.NewRecorder()

	testSpace := entities.ParkingSpace{}
	testError := repositories.NewInvalidInputError("boom")
	mockFacade.EXPECT().UpdateParkingSpaceStatus(req.Context(), testID, "available").Return(testSpace, testError).Once()

	// --------
	// ACT
	// --------
	handler.ServeHTTP(rr, req)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, http.StatusBadRequest, rr.Code, "Response codes don't match, should be 400")
	assert.Equal(t, "boom\n", rr.Body.String(), "Response body is wrong")
	mockFacade.AssertExpectations(t)
}

func TestParkingSpaceUpdateStatusHandler_ServeHTTP_UnhappyPath_NotFoundError(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockFacade := &mocks.ParkingSpaceFacade{}
	handler := handlers.NewParkingSpaceStatusHandler(mockFacade, testLogger)

	testModel := models.UpdateParkingSpaceStatus{
		Status: "available",
	}

	requestBody, _ := json.Marshal(testModel)
	testID := uuid.New()

	// Because we are directly calling the handler in the test without going through a router that parses the URL parameters,
	// we have to manually insert the URL parameters into the request context.
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", testID.String())

	// Create our request context with the formatted chi context.
	reqCtx := context.WithValue(context.Background(), chi.RouteCtxKey, rctx)
	req, _ := http.NewRequestWithContext(reqCtx, "PATCH", "/parking-space/"+testID.String()+"/status", bytes.NewBuffer(requestBody))
	rr := httptest.NewRecorder()

	testSpace := entities.ParkingSpace{}
	testError := repositories.NewNotFoundError("boom")
	mockFacade.EXPECT().UpdateParkingSpaceStatus(req.Context(), testID, "available").Return(testSpace, testError).Once()

	// --------
	// ACT
	// --------
	handler.ServeHTTP(rr, req)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, http.StatusBadRequest, rr.Code, "Response codes don't match, should be 400")
	assert.Equal(t, "Resource 'boom' not found\n", rr.Body.String(), "Response body is wrong")
	mockFacade.AssertExpectations(t)
}
