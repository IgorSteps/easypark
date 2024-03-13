package repositories

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
)

// UserRepository provides an interface for CRUD operations on users.
type UserRepository interface {
	// CreateUser creates a new user.
	CreateUser(ctx context.Context, user *entities.User) error

	// CheckUserExists checks if a user exists with given username and email.
	CheckUserExists(ctx context.Context, email, uname string) (bool, error)

	// FindByUsername gets the user(if it exists) from the database using their username.
	FindByUsername(ctx context.Context, username string) (*entities.User, error)
}
