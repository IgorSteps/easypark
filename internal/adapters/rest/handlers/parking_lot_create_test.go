package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/IgorSteps/easypark/internal/adapters/rest/handlers"
	"github.com/IgorSteps/easypark/internal/adapters/rest/models"
	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	mocks "github.com/IgorSteps/easypark/mocks/adapters/rest/handlers"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestParkingLotCreateHandler_ServeHTTP_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockFacade := &mocks.ParkingLotFacade{}
	handler := handlers.NewParkingLotCreateHandler(testLogger, mockFacade)

	testReq := &models.CreateParkingLotRequest{
		Name:     "sci",
		Capacity: 1,
	}

	requestBody, _ := json.Marshal(testReq)
	req, _ := http.NewRequest("POST", "/parking-requests", bytes.NewBuffer(requestBody))
	rr := httptest.NewRecorder()

	testLot := entities.ParkingLot{
		ID:       uuid.New(),
		Name:     testReq.Name,
		Capacity: testReq.Capacity,
	}
	mockFacade.EXPECT().CreateParkingLot(req.Context(), testReq.Name, testReq.Capacity).Return(testLot, nil).Once()

	// ---
	// ACT
	// ---
	handler.ServeHTTP(rr, req)

	// ------
	// ASSERT
	// ------
	assert.Equal(t, http.StatusCreated, rr.Code, "Response codes don't match, should be 201 CREATED")

	var targetModel models.CreateParkingLotResponse
	err := json.Unmarshal(rr.Body.Bytes(), &targetModel)
	assert.NoError(t, err, "Must not return err")

	assert.Equal(t, testReq.Name, targetModel.Name, "Reponse bodies names don't match")
	assert.Equal(t, testReq.Capacity, targetModel.Capacity, "Reponse bodies capacities don't match")
	mockFacade.AssertExpectations(t)
}

func TestParkingLotCreateHandler_ServeHTTP_UnhappyPath_ResourceAlreadyExistsError(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockFacade := &mocks.ParkingLotFacade{}
	handler := handlers.NewParkingLotCreateHandler(testLogger, mockFacade)

	testReq := &models.CreateParkingLotRequest{
		Name:     "sci",
		Capacity: 1,
	}

	requestBody, _ := json.Marshal(testReq)
	req, _ := http.NewRequest("POST", "/parking-requests", bytes.NewBuffer(requestBody))
	rr := httptest.NewRecorder()

	testLot := entities.ParkingLot{}
	testErr := repositories.NewResourceAlreadyExistsError("boom")
	mockFacade.EXPECT().CreateParkingLot(req.Context(), testReq.Name, testReq.Capacity).Return(testLot, testErr).Once()

	// ---
	// ACT
	// ---
	handler.ServeHTTP(rr, req)

	// ------
	// ASSERT
	// ------
	assert.Equal(t, http.StatusBadRequest, rr.Code, "Response codes don't match, should be 400")
	assert.Equal(t, "Resource 'boom' already exists\n", rr.Body.String(), "Response body is wrong")
	mockFacade.AssertExpectations(t)
}

func TestParkingLotCreateHandler_ServeHTTP_UnhappyPath_InternalError(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockFacade := &mocks.ParkingLotFacade{}
	handler := handlers.NewParkingLotCreateHandler(testLogger, mockFacade)

	testReq := &models.CreateParkingLotRequest{
		Name:     "sci",
		Capacity: 1,
	}

	requestBody, _ := json.Marshal(testReq)
	req, _ := http.NewRequest("POST", "/parking-requests", bytes.NewBuffer(requestBody))
	rr := httptest.NewRecorder()

	testLot := entities.ParkingLot{}
	testErr := repositories.NewInternalError("boom")
	mockFacade.EXPECT().CreateParkingLot(req.Context(), testReq.Name, testReq.Capacity).Return(testLot, testErr).Once()

	// ---
	// ACT
	// ---
	handler.ServeHTTP(rr, req)

	// ------
	// ASSERT
	// ------
	assert.Equal(t, http.StatusInternalServerError, rr.Code, "Response codes don't match, should be 400")
	assert.Equal(t, "Internal error: boom\n", rr.Body.String(), "Response body is wrong")
	mockFacade.AssertExpectations(t)
}
