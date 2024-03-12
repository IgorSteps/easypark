package datastore

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
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

func (s *UserPostgresRepository) CreateUser(ctx context.Context, user *entities.User) error {
	result := s.DB.Create(ctx, &user)
	if result.Error != nil {
		s.Logger.WithError(result.Error).Error("failed to insert user into the database")

		return result.Error
	}

	return nil
}

// CheckUserExists queries DB to check if user with given email or username exists.
func (s *UserPostgresRepository) CheckUserExists(ctx context.Context, email, uname string) (bool, error) {
	var user entities.User
	result := s.DB.Where(ctx, "email = ? OR username = ?", email, uname).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, nil
		}

		s.Logger.WithError(result.Error).Error("failed to query for user in the database")
		return false, result.Error
	}

	return true, nil // User found
}
