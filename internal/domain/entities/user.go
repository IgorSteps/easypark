package entities

import (
	"github.com/google/uuid"
)

// Represents a User in EasyPark: a driver or an admin.
type User struct {
	ID        uuid.UUID
	Username  string
	Email     string
	Password  string
	FirstName string
	LastName  string
	Role      UserRole
}
