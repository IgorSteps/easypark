package handlers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/IgorSteps/easypark/internal/adapters/rest/handlers"
	"github.com/IgorSteps/easypark/internal/adapters/rest/models"
	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	mocks "github.com/IgorSteps/easypark/mocks/adapters/rest/handlers"
	"github.com/google/uuid"
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
	handler := handlers.NewDriverCreateHandler(mockFacade, testLogger)

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

	testUser := CreateTestUser()
	mockFacade.EXPECT().CreateDriver(req.Context(), testDomainUser).Return(testUser, nil).Once()

	// --------
	// ACT
	// --------
	handler.ServeHTTP(rr, req)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, http.StatusCreated, rr.Code, "Response codes don't match, should be 201 CREATED")

	var targetModel entities.User
	err := json.Unmarshal(rr.Body.Bytes(), &targetModel)
	assert.NoError(t, err, "Failed to unmarshall response body")
	assert.Equal(t, testUser, &targetModel, "Reponse bodies don't match")
	mockFacade.AssertExpectations(t)
}

func TestUserCreateHandler_ServeHTTP_UnhappyPath_DecoderFailure(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockFacade := &mocks.UserFacade{}
	handler := handlers.NewDriverCreateHandler(mockFacade, testLogger)
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

func TestUserCreateHandler_ServeHTTP_UnhappyPath_UserAlreadyExistsError(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockFacade := &mocks.UserFacade{}
	handler := handlers.NewDriverCreateHandler(mockFacade, testLogger)
	testUsername := "testuser"
	testEmail := "test@example.com"
	testUserReq := models.UserCreationRequest{
		Firstname: "John",
		Lastname:  "Smith",
		Username:  testUsername,
		Password:  "password",
		Email:     testEmail,
	}
	testError := repositories.NewResourceAlreadyExistsError(testUsername)
	testDomainUser := testUserReq.ToDomain()

	requestBody, _ := json.Marshal(testUserReq)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(requestBody))
	rr := httptest.NewRecorder()

	mockFacade.EXPECT().CreateDriver(req.Context(), testDomainUser).Return(nil, testError).Once()

	// --------
	// ACT
	// --------
	handler.ServeHTTP(rr, req)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, http.StatusBadRequest, rr.Code, "Response codes don't match, should be 400")
	assert.Contains(t, rr.Body.String(), fmt.Sprintf("Resource '%s' already exists", testUsername), "Reponse bodies don't match")

	// Assert logger.
	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	assert.Equal(t, "failed to create user", hook.LastEntry().Message, "Error messages are not equal")
	assert.Equal(t, testError, hook.LastEntry().Data["error"], "Errors are different")

	hook.Reset()
	assert.Nil(t, hook.LastEntry())

	mockFacade.AssertExpectations(t)
}

func TestUserCreateHandler_ServeHTTP_UnhappyPath_InternalError(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, hook := test.NewNullLogger()
	mockFacade := &mocks.UserFacade{}
	handler := handlers.NewDriverCreateHandler(mockFacade, testLogger)

	testUserReq := models.UserCreationRequest{
		Firstname: "John",
		Lastname:  "Smith",
		Username:  "testUsername",
		Password:  "password",
		Email:     "testEmail",
	}
	errorMessage := "boom"
	testError := repositories.NewInternalError(errorMessage)
	testDomainUser := testUserReq.ToDomain()

	requestBody, _ := json.Marshal(testUserReq)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(requestBody))
	rr := httptest.NewRecorder()

	mockFacade.EXPECT().CreateDriver(req.Context(), testDomainUser).Return(nil, testError).Once()

	// --------
	// ACT
	// --------
	handler.ServeHTTP(rr, req)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, http.StatusInternalServerError, rr.Code, "Response codes don't match, should be 500")
	assert.Contains(t, rr.Body.String(), "Internal error: boom", "Reponse bodies don't match")

	// Assert logger.
	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	assert.Equal(t, "failed to create user", hook.LastEntry().Message, "Error messages are not equal")
	assert.Equal(t, testError, hook.LastEntry().Data["error"], "Errors are different")

	hook.Reset()
	assert.Nil(t, hook.LastEntry())

	mockFacade.AssertExpectations(t)
}

func CreateTestUser() *entities.User {
	return &entities.User{
		ID:        uuid.New(),
		Username:  "boom",
		Email:     "bom",
		Password:  "foo",
		FirstName: "john",
		LastName:  "smith",
		Role:      entities.RoleDriver,
	}
}
