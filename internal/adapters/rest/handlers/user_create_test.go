package handlers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/IgorSteps/easypark/internal/adapters/rest/handlers"
	"github.com/IgorSteps/easypark/internal/adapters/rest/models"
	mocks "github.com/IgorSteps/easypark/mocks/adapters/rest/handlers"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestUserCreateHandler_ServeHTTP_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockFacade := &mocks.UserFacade{}
	handler := handlers.NewUserCreateHandler(mockFacade, testLogger)

	testUserReq := models.UserCreationRequest{
		Firstname: "John",
		Lastname:  "Smith",
		Username:  "testuser",
		Password:  "password",
		Email:     "test@example.com",
	}
	testDomainUser := testUserReq.ToDomain()

	requestBody, _ := json.Marshal(testUserReq)

	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(requestBody))
	rr := httptest.NewRecorder()

	mockFacade.EXPECT().CreateUser(req.Context(), testDomainUser).Return(nil).Once()

	// --------
	// ACT
	// --------
	handler.ServeHTTP(rr, req)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, http.StatusCreated, rr.Code, "Response codes don't match, should be 201 CREATED")
	assert.Contains(t, rr.Body.String(), "user created successfully", "Reponse bodies don't match")
	mockFacade.AssertExpectations(t)
}

func TestUserCreateHandler_ServeHTTP_UnhappyPath_DecoderFailure(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockFacade := &mocks.UserFacade{}
	handler := handlers.NewUserCreateHandler(mockFacade, testLogger)
	testUserReq := []byte(`{"invalid":"testuser"}`)

	requestBody, _ := json.Marshal(testUserReq)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(requestBody))
	rr := httptest.NewRecorder()

	// --------
	// ACT
	// --------
	handler.ServeHTTP(rr, req)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, http.StatusBadRequest, rr.Code, "Response codes don't match, should be 400 Bad Request")
	assert.Contains(t, rr.Body.String(), "invalid request body", "Reponse bodies don't match")

	// Assert logger.
	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	assert.Equal(t,
		"failed to decode user creation request: json: cannot unmarshal string into Go value of type models.UserCreationRequest",
		hook.LastEntry().Message,
		"Messages are not equal",
	)

	hook.Reset()
	assert.Nil(t, hook.LastEntry())

	mockFacade.AssertExpectations(t)
}

func TestUserCreateHandler_ServeHTTP_UnhappyPath_FacadeFailure(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockFacade := &mocks.UserFacade{}
	handler := handlers.NewUserCreateHandler(mockFacade, testLogger)

	testError := errors.New("boom")
	testUserReq := models.UserCreationRequest{
		Firstname: "John",
		Lastname:  "Smith",
		Username:  "testuser",
		Password:  "password",
		Email:     "test@example.com",
	}
	testDomainUser := testUserReq.ToDomain()

	requestBody, _ := json.Marshal(testUserReq)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(requestBody))
	rr := httptest.NewRecorder()

	mockFacade.EXPECT().CreateUser(req.Context(), testDomainUser).Return(testError).Once()

	// --------
	// ACT
	// --------
	handler.ServeHTTP(rr, req)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, http.StatusInternalServerError, rr.Code, "Response codes don't match, should be 500 Internal Server Error")
	assert.Contains(t, rr.Body.String(), "failed to create user", "Reponse bodies don't match")

	// Assert logger.
	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	assert.Equal(t,
		fmt.Sprintf("failed to create user: %s", testError.Error()),
		hook.LastEntry().Message,
		"Messages are not equal",
	)

	hook.Reset()
	assert.Nil(t, hook.LastEntry())

	mockFacade.AssertExpectations(t)
}
