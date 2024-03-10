package repositories

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
)

// UserRepository provides an interface for CRUD operations on users.
type UserRepository interface {
	// CreateUser creates a new user.
	CreateUser(ctx context.Context, user entities.User) (entities.User, error)

	// CheckUserExists checs if a user with given email exists.
	CheckUserExistsByEmail(ctx context.Context, email string) (bool, error)
}
