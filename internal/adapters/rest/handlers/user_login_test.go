package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/IgorSteps/easypark/internal/adapters/rest/handlers"
	"github.com/IgorSteps/easypark/internal/adapters/rest/models"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	mocks "github.com/IgorSteps/easypark/mocks/adapters/rest/handlers"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestUserLoginHandler_ServeHTTP_HappyPath(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockFacade := &mocks.UserFacade{}
	handler := handlers.NewUserLoginHandler(mockFacade, testLogger)

	testToken := "haha"
	testUserReq := models.LoginUserRequest{
		Username: "testuser",
		Password: "password",
	}

	requestBody, _ := json.Marshal(testUserReq)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(requestBody))
	rr := httptest.NewRecorder()

	user := CreateTestUser()
	mockFacade.EXPECT().AuthoriseUser(req.Context(), testUserReq.Username, testUserReq.Password).Return(user, testToken, nil).Once()

	// --------
	// ACT
	// --------
	handler.ServeHTTP(rr, req)

	// --------
	// ASSERT
	// --------
	var target models.LoginUserResponse
	err := json.Unmarshal(rr.Body.Bytes(), &target)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rr.Code, "Response codes don't match, should be 200")
	assert.Equal(t, target.User, *user, "User bodies don't match")
	assert.Equal(t, target.Token, testToken, "Token bodies don't match")
	mockFacade.AssertExpectations(t)
}

func TestUserLoginHandler_ServeHTTP_UnhappyPath_InvalidCredentials(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockFacade := &mocks.UserFacade{}
	handler := handlers.NewUserLoginHandler(mockFacade, testLogger)

	emptyToken := ""
	testError := repositories.NewInvalidInputError("invalid password")
	testUserReq := models.LoginUserRequest{
		Username: "testuser",
		Password: "password",
	}

	requestBody, _ := json.Marshal(testUserReq)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(requestBody))
	rr := httptest.NewRecorder()

	mockFacade.EXPECT().AuthoriseUser(req.Context(), testUserReq.Username, testUserReq.Password).Return(nil, emptyToken, testError).Once()

	// --------
	// ACT
	// --------
	handler.ServeHTTP(rr, req)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, http.StatusUnauthorized, rr.Code, "Response codes don't match, should be 401")
	assert.Contains(t, rr.Body.String(), "invalid password", "Reponse bodies don't match")
	mockFacade.AssertExpectations(t)
}

func TestUserLoginHandler_ServeHTTP_UnhappyPath_UserNotFoundError(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockFacade := &mocks.UserFacade{}
	handler := handlers.NewUserLoginHandler(mockFacade, testLogger)

	emptyToken := ""
	testUserReq := models.LoginUserRequest{
		Username: "testuser",
		Password: "password",
	}
	testError := repositories.NewNotFoundError(testUserReq.Username)

	requestBody, _ := json.Marshal(testUserReq)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(requestBody))
	rr := httptest.NewRecorder()

	mockFacade.EXPECT().AuthoriseUser(req.Context(), testUserReq.Username, testUserReq.Password).Return(nil, emptyToken, testError).Once()

	// --------
	// ACT
	// --------
	handler.ServeHTTP(rr, req)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, http.StatusUnauthorized, rr.Code, "Response codes don't match, should be 401")
	assert.Contains(t, rr.Body.String(), "Resource 'testuser' not found", "Reponse bodies don't match")
	mockFacade.AssertExpectations(t)
}

func TestUserLoginHandler_ServeHTTP_UnhappyPath_InternalError(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockFacade := &mocks.UserFacade{}
	handler := handlers.NewUserLoginHandler(mockFacade, testLogger)

	emptyToken := ""
	testUserReq := models.LoginUserRequest{
		Username: "testuser",
		Password: "password",
	}
	testError := repositories.NewInternalError("boom")

	requestBody, _ := json.Marshal(testUserReq)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(requestBody))
	rr := httptest.NewRecorder()

	mockFacade.EXPECT().AuthoriseUser(req.Context(), testUserReq.Username, testUserReq.Password).Return(nil, emptyToken, testError).Once()

	// --------
	// ACT
	// --------
	handler.ServeHTTP(rr, req)

	// --------
	// ASSERT
	// --------
	assert.Equal(t, http.StatusInternalServerError, rr.Code, "Response codes don't match, should be 500")
	assert.Contains(t, rr.Body.String(), "Internal error: boom", "Reponse bodies don't match")
	mockFacade.AssertExpectations(t)
}
