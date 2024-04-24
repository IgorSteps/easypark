package datastore

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// UserPostgresRepository implements UserRepository interface to provide database operation on Users.
type UserPostgresRepository struct {
	Logger *logrus.Logger
	DB     Datastore
}

// NewUserPostgresRepository retuns new UserPostgresRepository.
func NewUserPostgresRepository(db Datastore, lgr *logrus.Logger) *UserPostgresRepository {
	return &UserPostgresRepository{
		Logger: lgr,
		DB:     db,
	}
}

// CreateUser creates a record in the database with a given user.
func (s *UserPostgresRepository) CreateUser(ctx context.Context, user *entities.User) error {
	result := s.DB.WithContext(ctx).Create(user)
	err := result.Error()
	if err != nil {
		s.Logger.WithError(err).Error("failed to insert user into the database")
		return repositories.NewInternalError("failed to insert user into the database")
	}

	return nil
}

// CheckUserExists queries DB to check if user with given email or username exists.
func (s *UserPostgresRepository) CheckUserExists(ctx context.Context, email, uname string) (bool, error) {
	var user entities.User

	result := s.DB.WithContext(ctx).Where("email = ? OR username = ?", email, uname).First(&user)
	err := result.Error()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil // User not found
		}

		s.Logger.WithError(err).WithFields(logrus.Fields{
			"email":    email,
			"username": uname,
		}).Error("failed to query for user in the database")
		return false, repositories.NewInternalError("failed to query for user in the database")
	}

	return true, nil // User found
}

// GetDriverByUsername queries DB to find user with given username.
func (s *UserPostgresRepository) GetDriverByUsername(ctx context.Context, username string) (*entities.User, error) {
	var user entities.User

	result := s.DB.WithContext(ctx).Where("username = ?", username).First(&user)
	err := result.Error()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			s.Logger.WithField("username", username).Warn("failed to find user with given username in the database")
			return nil, repositories.NewNotFoundError(username)
		}

		s.Logger.WithError(err).Error("failed to query for user in the database")
		return nil, repositories.NewInternalError("failed to query for user in the database")
	}

	return &user, nil // User found
}

// GetAllDriverUsers gets all the driver users from the databas.
func (s *UserPostgresRepository) GetAllDriverUsers(ctx context.Context) ([]entities.User, error) {
	var users []entities.User

	// Get all users with driver role.
	result := s.DB.WithContext(ctx).Where("role <> ?", entities.RoleAdmin).FindAll(&users)
	err := result.Error()
	if err != nil {
		s.Logger.WithError(err).Error("failed to query for all drivers in the database")
		return users, repositories.NewInternalError("failed to query for all drivers in the database")
	}

	return users, nil
}

// GetSingle queries the DB to find a user by their ID.
func (s *UserPostgresRepository) GetSingle(ctx context.Context, id uuid.UUID, user *entities.User) error {
	result := s.DB.WithContext(ctx).First(user, "id = ?", id)
	err := result.Error()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			s.Logger.WithField("id", id).Error("failed to find driver with given id in the database")
			return repositories.NewNotFoundError(id.String())
		}

		s.Logger.WithError(err).Error("failed to query for user in the database")
		return repositories.NewInternalError("failed to query for user in the database")
	}

	return nil
}

// Save saves updated user into the DB.
func (s *UserPostgresRepository) Save(ctx context.Context, user *entities.User) error {
	result := s.DB.WithContext(ctx).Save(user)
	err := result.Error()
	if err != nil {
		s.Logger.WithError(err).Error("failed to save updated user in the database")
		return repositories.NewInternalError("failed to save updated user in the database")
	}

	return nil
}
