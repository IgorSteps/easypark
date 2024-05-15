package entities

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUser_SetOnCreate(t *testing.T) {
	user := User{
		Username:  "testuser",
		Email:     "test@example.com",
		Password:  "securepassword",
		FirstName: "Test",
		LastName:  "User",
	}

	// Before calling SetOnCreate
	assert.Empty(t, user.ID, "ID should be initially empty")
	assert.NotEqual(t, RoleDriver, user.Role, "Role should not be preset as Driver")
	assert.NotEqual(t, StatusActive, user.Status, "Status should not be preset as Active")

	// Action
	user.SetOnCreate()

	// After calling SetOnCreate
	assert.NotEmpty(t, user.ID, "ID should be set")
	assert.Equal(t, RoleDriver, user.Role, "Role should be set to Driver")
	assert.Equal(t, StatusActive, user.Status, "Status should be set to Active")
}

func TestUser_Ban(t *testing.T) {
	user := User{
		ID:        uuid.New(),
		Username:  "bannableuser",
		Email:     "user@example.com",
		Password:  "securepassword",
		FirstName: "Bannable",
		LastName:  "User",
		Status:    StatusActive, // initially active
	}

	// Before calling Ban
	assert.Equal(t, StatusActive, user.Status, "User should be active before ban")

	// Action
	user.Ban()

	// After calling Ban
	assert.Equal(t, StatusBanned, user.Status, "User status should be set to Banned")
}
