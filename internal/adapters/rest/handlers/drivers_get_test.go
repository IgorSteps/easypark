package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/IgorSteps/easypark/internal/adapters/rest/handlers"
	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	mocks "github.com/IgorSteps/easypark/mocks/adapters/rest/handlers"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestDriverUsersGetHander_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockFacade := &mocks.UserFacade{}
	handler := handlers.NewDriverUsersGetHandler(testLogger, mockFacade)

	expectedUsers := []entities.User{
		{
			FirstName: "John",
			LastName:  "Smith",
			Username:  "testuser",
			Password:  "password",
			Email:     "test@example.com",
		},
		{
			FirstName: "Boom",
			LastName:  "Bam",
			Username:  "testuser",
			Password:  "password",
			Email:     "test@example.com",
		},
	}

	req, _ := http.NewRequest("GET", "/drivers", bytes.NewBuffer(nil))
	rr := httptest.NewRecorder()

	mockFacade.EXPECT().GetAllDriverUsers(req.Context()).Return(expectedUsers, nil).Once()

	// --------
	// ACT
	// --------
	handler.ServeHTTP(rr, req)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, http.StatusOK, rr.Code, "Response codes don't match, should be 200 OK")

	// Unmarshal response body into slice of Users.
	var actualUsers []entities.User
	err := json.Unmarshal(rr.Body.Bytes(), &actualUsers)
	assert.NoError(t, err, "Must have no error unmarshaling response body")
	assert.Equal(t, expectedUsers, actualUsers, "Users don't match")

	mockFacade.AssertExpectations(t)
}

func TestDriverUsersGetHander_UnhappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockFacade := &mocks.UserFacade{}
	handler := handlers.NewDriverUsersGetHandler(testLogger, mockFacade)

	testError := repositories.NewInternalError("boom")

	req, _ := http.NewRequest("GET", "/drivers", bytes.NewBuffer(nil))
	rr := httptest.NewRecorder()

	mockFacade.EXPECT().GetAllDriverUsers(req.Context()).Return(nil, testError).Once()

	// --------
	// ACT
	// --------
	handler.ServeHTTP(rr, req)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, http.StatusInternalServerError, rr.Code, "Response codes don't match, should be 500")

	// Assert logger
	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	assert.Equal(t, "failed to get all drivers", hook.LastEntry().Message)
	assert.Equal(t, testError, hook.LastEntry().Data["error"])

	mockFacade.AssertExpectations(t)
}
