package enitities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Username  string
	Email     string
	Password  string
	FirstName string
	LastName  string
	CreatedAt time.Time
	Role      UserRole
}
