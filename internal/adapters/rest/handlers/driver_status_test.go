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
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestDriverStatusHandler_ServeHTTP_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockFacade := &mocks.UserFacade{}
	handler := handlers.NewDriverStatusHandler(mockFacade, testLogger)

	testStatusRequest := models.UpdateStatusRequest{
		Status: "ban",
	}

	requestBody, _ := json.Marshal(testStatusRequest)
	testID := uuid.New()

	// Because we are directly calling the handler in the test without going through a router that parses the URL parameters,
	// we have to manually insert the URL parameters into the request context.
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", testID.String())

	// Create our request context with the formatted chi context.
	reqCtx := context.WithValue(context.Background(), chi.RouteCtxKey, rctx)
	req, _ := http.NewRequestWithContext(reqCtx, "POST", "/drivers/"+testID.String()+"/status", bytes.NewBuffer(requestBody))
	rr := httptest.NewRecorder()

	mockFacade.EXPECT().BanDriver(req.Context(), testID).Return(nil).Once()

	// --------
	// ACT
	// --------
	handler.ServeHTTP(rr, req)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, http.StatusOK, rr.Code, "Response codes don't match, should be 200")
	assert.Contains(t, rr.Body.String(), "successfully updated user status", "Reponse bodies don't match")
	mockFacade.AssertExpectations(t)
}

func TestDriverStatusHandler_ServeHTTP_UnhappyPath_UserNotFoundError(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockFacade := &mocks.UserFacade{}
	handler := handlers.NewDriverStatusHandler(mockFacade, testLogger)

	testStatusRequest := models.UpdateStatusRequest{
		Status: "ban",
	}

	requestBody, _ := json.Marshal(testStatusRequest)
	testID := uuid.New()
	testErr := repositories.NewNotFoundError("no")

	// Because we are directly calling the handler in the test without going through a router that parses the URL parameters,
	// we have to manually insert the URL parameters into the request context.
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", testID.String())

	// Create our request context with the formatted chi context.
	reqCtx := context.WithValue(context.Background(), chi.RouteCtxKey, rctx)
	req, _ := http.NewRequestWithContext(reqCtx, "POST", "/drivers/"+testID.String()+"/status", bytes.NewBuffer(requestBody))
	rr := httptest.NewRecorder()

	mockFacade.EXPECT().BanDriver(req.Context(), testID).Return(testErr).Once()

	// --------
	// ACT
	// --------
	handler.ServeHTTP(rr, req)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, http.StatusBadRequest, rr.Code, "Response codes don't match, should be 400")
	assert.Contains(t, rr.Body.String(), "User 'no' not found\n", "Response bodies don't match")
	mockFacade.AssertExpectations(t)

	// Assert logger
	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	assert.Equal(t, "failed to ban user", hook.LastEntry().Message)
	assert.Equal(t, testErr, hook.LastEntry().Data["error"])
}

func TestDriverStatusHandler_ServeHTTP_UnhappyPath_InternalError(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockFacade := &mocks.UserFacade{}
	handler := handlers.NewDriverStatusHandler(mockFacade, testLogger)

	testStatusRequest := models.UpdateStatusRequest{
		Status: "ban",
	}

	requestBody, _ := json.Marshal(testStatusRequest)
	testID := uuid.New()
	testErr := repositories.NewInternalError("no")

	// Because we are directly calling the handler in the test without going through a router that parses the URL parameters,
	// we have to manually insert the URL parameters into the request context.
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", testID.String())

	// Create our request context with the formatted chi context.
	reqCtx := context.WithValue(context.Background(), chi.RouteCtxKey, rctx)
	req, _ := http.NewRequestWithContext(reqCtx, "POST", "/drivers/"+testID.String()+"/status", bytes.NewBuffer(requestBody))
	rr := httptest.NewRecorder()

	mockFacade.EXPECT().BanDriver(req.Context(), testID).Return(testErr).Once()

	// --------
	// ACT
	// --------
	handler.ServeHTTP(rr, req)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, http.StatusInternalServerError, rr.Code, "Response codes don't match, should be 500")
	assert.Contains(t, rr.Body.String(), "Internal error: no\n", "Response bodies don't match")
	mockFacade.AssertExpectations(t)

	// Assert logger
	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	assert.Equal(t, "failed to ban user", hook.LastEntry().Message)
	assert.Equal(t, testErr, hook.LastEntry().Data["error"])
}

func TestDriverStatusHandler_ServeHTTP_UnhappyPath_InvalidStatus(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockFacade := &mocks.UserFacade{}
	handler := handlers.NewDriverStatusHandler(mockFacade, testLogger)

	// Set invalid status in the request.
	testStatusRequest := models.UpdateStatusRequest{
		Status: "boom",
	}

	requestBody, _ := json.Marshal(testStatusRequest)
	testID := uuid.New()

	// Because we are directly calling the handler in the test without going through a router that parses the URL parameters,
	// we have to manually insert the URL parameters into the request context.
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", testID.String())

	// Create our request context with the formatted chi context.
	reqCtx := context.WithValue(context.Background(), chi.RouteCtxKey, rctx)
	req, _ := http.NewRequestWithContext(reqCtx, "POST", "/drivers/"+testID.String()+"/status", bytes.NewBuffer(requestBody))
	rr := httptest.NewRecorder()

	// --------
	// ACT
	// --------
	handler.ServeHTTP(rr, req)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, http.StatusNotImplemented, rr.Code, "Response codes don't match, should be 501")
	assert.Contains(t, rr.Body.String(), "unimplemented status", "Response bodies don't match")
	mockFacade.AssertExpectations(t)

	// Assert logger
	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.WarnLevel, hook.LastEntry().Level)
	assert.Equal(t, "unimplemented status update workflow called", hook.LastEntry().Message)
	assert.Equal(t, testStatusRequest.Status, hook.LastEntry().Data["status"])
}

func TestDriverStatusHandler_ServeHTTP_UnhappyPath_ParsingID(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockFacade := &mocks.UserFacade{}
	handler := handlers.NewDriverStatusHandler(mockFacade, testLogger)

	testStatusRequest := models.UpdateStatusRequest{
		Status: "ban",
	}

	requestBody, _ := json.Marshal(testStatusRequest)
	// Create invalid uuid.
	testInvalidID := "wroom"

	// Because we are directly calling the handler in the test without going through a router that parses the URL parameters,
	// we have to manually insert the URL parameters into the request context.
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("id", testInvalidID)

	// Create our request context with the formatted chi context.
	reqCtx := context.WithValue(context.Background(), chi.RouteCtxKey, rctx)
	req, _ := http.NewRequestWithContext(reqCtx, "POST", "/drivers/"+testInvalidID+"/status", bytes.NewBuffer(requestBody))
	rr := httptest.NewRecorder()

	// --------
	// ACT
	// --------
	handler.ServeHTTP(rr, req)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, http.StatusBadRequest, rr.Code, "Response codes don't match, should be 400")
	assert.Contains(t, rr.Body.String(), "invalid user id", "Response bodies don't match")
	mockFacade.AssertExpectations(t)

	// Assert logger
	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	assert.Equal(t, "failed to parse driver id", hook.LastEntry().Message)
	assert.NotEmpty(t, hook.LastEntry().Data["error"])
}
