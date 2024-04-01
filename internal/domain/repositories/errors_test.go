package repositories_test

import (
	"fmt"
	"testing"

	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/stretchr/testify/assert"
)

func TestUserNotFoundError(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	username := "testuser"
	err := repositories.NewNotFoundError(username)
	expectedMsg := fmt.Sprintf("User '%s' not found", username)

	// ----
	// ACT
	// ----
	errString := err.Error()

	// ------
	// ASSERT
	// ------
	assert.Equal(t, expectedMsg, errString, "Expected message to be: %s, got: %s", expectedMsg, err.Error())
}

func TestUserAlreadyExistsError(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	username := "testuser"
	email := "testemail"
	err := repositories.NewUserAlreadyExistsError(username, email)
	expectedMsg := fmt.Sprintf("User '%s'/'%s' already exists", username, email)

	// ----
	// ACT
	// ----
	errString := err.Error()

	// ------
	// ASSERT
	// ------
	assert.Equal(t, expectedMsg, errString, "Expected message to be: %s, got: %s", expectedMsg, err.Error())
}

func TestInvalidCredentialsError(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	errorMsg := "boom"
	err := repositories.NewInvalidInputError(errorMsg)

	// ----
	// ACT
	// ----
	errString := err.Error()

	// ------
	// ASSERT
	// ------
	assert.Equal(t, errorMsg, errString, "Expected message to be: %s, got: %s", errorMsg, err.Error())
}

func TestInternalError(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testError := "boom"
	err := repositories.NewInternalError(testError)
	expectedMsg := fmt.Sprintf("Internal error: %s", testError)

	// ----
	// ACT
	// ----
	errString := err.Error()

	// ------
	// ASSERT
	// ------
	assert.Equal(t, expectedMsg, errString, "Expected message to be: %s, got: %s", expectedMsg, err.Error())
}
