package entities

import (
	"github.com/google/uuid"
)

type UserRole string

const (
	RoleAdmin  UserRole = "admin"
	RoleDriver UserRole = "driver"
)

type UserStatus string

const (
	StatusActive UserStatus = "active"
	StatusBanned UserStatus = "banned"
)

// User represents a User in EasyPark.
type User struct {
	ID        uuid.UUID `gorm:"primary_key"`
	Username  string
	Email     string
	Password  string
	FirstName string
	LastName  string
	Status    UserStatus
	Role      UserRole
}

// SetOnCreate sets internally managed fileds on user creation.
func (s *User) SetOnCreate() {
	s.ID = uuid.New()
	s.Role = RoleDriver
	s.Status = StatusActive
}

// Ban bans the user.
func (s *User) Ban() {
	s.Status = StatusBanned
}
