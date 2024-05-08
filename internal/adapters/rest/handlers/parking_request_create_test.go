package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/IgorSteps/easypark/internal/adapters/rest/handlers"
	"github.com/IgorSteps/easypark/internal/adapters/rest/models"
	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	mocks "github.com/IgorSteps/easypark/mocks/adapters/rest/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestParkingRequestCreateHandler_ServeHTTP_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockFacade := &mocks.ParkingRequestFacade{}
	handler := handlers.NewParkingRequestCreateHandler(mockFacade, testLogger)

	testCreateParkingRequestRequest := models.CreateParkingRequestRequest{
		DestinationParkingLotID: uuid.New(),
		StartTime:               time.Now(),
		EndTime:                 time.Now().Add(5),
	}

	requestBody, err := json.Marshal(testCreateParkingRequestRequest)
	assert.NoError(t, err, "Marshalling parking request to json must not return error")

	testID := uuid.New()

	// Because we are directly calling the handler in the test without going through a router that parses the URL parameters,
	// we have to manually insert the URL parameters into the request context.
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", testID.String())

	// Create our request context with the formatted chi context.
	reqCtx := context.WithValue(context.Background(), chi.RouteCtxKey, rctx)
	req, _ := http.NewRequestWithContext(reqCtx, "POST", "/drivers/"+testID.String()+"/parking-requests", bytes.NewBuffer(requestBody))
	rr := httptest.NewRecorder()

	createdParkingRequest := &entities.ParkingRequest{
		ID:                      uuid.New(),
		UserID:                  uuid.New(),
		ParkingSpaceID:          nil,
		DestinationParkingLotID: uuid.New(),
		StartTime:               time.Now(),
		EndTime:                 time.Now().Add(5),
		Status:                  entities.RequestStatusPending,
		Cost:                    10,
	}
	mockFacade.EXPECT().CreateParkingRequest(req.Context(), mock.Anything).Return(createdParkingRequest, nil).Once()

	// Create a response from the created parking request to be able to comapre with the response body
	expectedResponse := models.CreateParkingRequestResponse{
		ID:          createdParkingRequest.ID,
		UserID:      createdParkingRequest.UserID,
		DestinationLotID: createdParkingRequest.DestinationParkingLotID,
		StartTime:   createdParkingRequest.StartTime,
		EndTime:     createdParkingRequest.EndTime,
		Status:      createdParkingRequest.Status,
		Cost:        createdParkingRequest.Cost,
	}
	expectedJson, _ := json.Marshal(expectedResponse)

	// --------
	// ACT
	// --------
	handler.ServeHTTP(rr, req)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, http.StatusCreated, rr.Code, "Response codes don't match, should be 201 CREATED")
	assert.Equal(t, string(expectedJson)+"\n", rr.Body.String(), "Responses don't match")
	mockFacade.AssertExpectations(t)
}

func TestParkingRequestCreateHandler_ServeHTTP_UnhappyPath_InternalError(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockFacade := &mocks.ParkingRequestFacade{}
	handler := handlers.NewParkingRequestCreateHandler(mockFacade, testLogger)

	testCreateParkingRequestRequest := models.CreateParkingRequestRequest{
		DestinationParkingLotID: uuid.New(),
		StartTime:               time.Now(),
		EndTime:                 time.Now().Add(5),
	}

	requestBody, err := json.Marshal(testCreateParkingRequestRequest)
	assert.NoError(t, err, "Marshalling parking request to json must not return error")

	testID := uuid.New()

	// Because we are directly calling the handler in the test without going through a router that parses the URL parameters,
	// we have to manually insert the URL parameters into the request context.
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", testID.String())

	// Create our request context with the formatted chi context.
	reqCtx := context.WithValue(context.Background(), chi.RouteCtxKey, rctx)
	req, _ := http.NewRequestWithContext(reqCtx, "POST", "/drivers/"+testID.String()+"/parking-requests", bytes.NewBuffer(requestBody))
	rr := httptest.NewRecorder()

	testError := repositories.NewInternalError("boom")
	createdParkingRequest := &entities.ParkingRequest{}
	mockFacade.EXPECT().CreateParkingRequest(req.Context(), mock.Anything).Return(createdParkingRequest, testError).Once()

	// --------
	// ACT
	// --------
	handler.ServeHTTP(rr, req)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, http.StatusInternalServerError, rr.Code, "Response codes don't match, should be 500")
	assert.Equal(t, testError.Error()+"\n", rr.Body.String(), "Response bodies don't match")
	mockFacade.AssertExpectations(t)
}
