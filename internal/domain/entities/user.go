package entities

import (
	"github.com/google/uuid"
)

type UserRole string

const (
	RoleAdmin  UserRole = "admin"
	RoleDriver UserRole = "driver"
)

// User represents a User in EasyPark.
type User struct {
	ID        uuid.UUID `gorm:"primary_key"`
	Username  string
	Email     string
	Password  string
	FirstName string
	LastName  string
	Role      UserRole
}
