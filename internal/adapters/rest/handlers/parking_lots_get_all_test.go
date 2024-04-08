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
	"github.com/google/uuid"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestParkingLotGetAllHandler_ServeHTTP_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockFacade := &mocks.ParkingLotFacade{}

	testLots := []entities.ParkingLot{
		{
			ID:       uuid.New(),
			Name:     "bob",
			Capacity: 10,
		},
		{
			ID:       uuid.New(),
			Name:     "dod",
			Capacity: 10,
		},
	}

	testCtx := context.Background()
	handler := handlers.NewParkingLotGetAllHandler(testLogger, mockFacade)
	mockFacade.EXPECT().GetAllParkingLots(testCtx).Return(testLots, nil).Once()

	req, _ := http.NewRequest("GET", "/parking-lots", bytes.NewBuffer(nil))
	rr := httptest.NewRecorder()

	// --------
	// ACT
	// --------
	handler.ServeHTTP(rr, req)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, http.StatusOK, rr.Code, "Response codes don't match, should be 200 OK")

	// Unmarshal response body into slice of ParkingRequests.
	var actualLots []entities.ParkingLot
	err := json.Unmarshal(rr.Body.Bytes(), &actualLots)
	assert.NoError(t, err, "Must have no error unmarshaling response body")
	assert.Equal(t, testLots, actualLots, "Parking lots don't match")

	mockFacade.AssertExpectations(t)
}

func TestParkingLotGetAllHandler_ServeHTTP_UnhappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockFacade := &mocks.ParkingLotFacade{}

	testCtx := context.Background()
	handler := handlers.NewParkingLotGetAllHandler(testLogger, mockFacade)
	testError := repositories.NewInternalError("boom")
	mockFacade.EXPECT().GetAllParkingLots(testCtx).Return(nil, testError).Once()

	req, _ := http.NewRequest("GET", "/parking-lots", bytes.NewBuffer(nil))
	rr := httptest.NewRecorder()

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
