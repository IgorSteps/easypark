package repositories

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/google/uuid"
)

// UserRepository provides an interface for CRUD operations on users.
type UserRepository interface {
	// CreateUser creates a new user.
	CreateUser(ctx context.Context, user *entities.User) error

	// CheckUserExists checks if a user exists with given username and email.
	CheckUserExists(ctx context.Context, email, uname string) (bool, error)

	// GetDriverByUsername gets the user(if it exists) from the database using their username.
	GetDriverByUsername(ctx context.Context, username string) (*entities.User, error)
	// GetAllDriverUsers gets all the driver users from the database.
	GetAllDriverUsers(ctx context.Context) ([]entities.User, error)

	// GetSingle gets a single user using their UUID for update or read.
	GetSingle(ctx context.Context, id uuid.UUID, user *entities.User) error

	GetAdmin(ctx context.Context, user *entities.User) error

	// Save saves the user when performing the Updating.
	Save(ctx context.Context, user *entities.User) error
}
