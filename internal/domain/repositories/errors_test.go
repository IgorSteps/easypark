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
	expectedMsg := fmt.Sprintf("Resource '%s' not found", username)

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
	err := repositories.NewResourceAlreadyExistsError(username)
	expectedMsg := fmt.Sprintf("Resource '%s' already exists", username)

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
