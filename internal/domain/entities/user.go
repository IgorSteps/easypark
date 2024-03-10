package entities

import (
	"github.com/google/uuid"
)

// Represents a User in EasyPark: a driver or an admin.
type User struct {
	ID        uuid.UUID `json:"-"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Role      UserRole  `json:"-"`
}
