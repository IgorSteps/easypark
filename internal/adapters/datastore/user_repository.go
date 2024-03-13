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
	result := s.DB.WithContext(ctx).Create(&user)
	err := result.Error()
	if err != nil {
		s.Logger.WithError(err).Error("failed to insert user into the database")

		return err
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
			s.Logger.Debug("User not found")
			return false, nil // User not found
		}

		s.Logger.WithError(err).Error("failed to query for user in the database")
		return false, err
	}
	s.Logger.Debug("User found")
	return true, nil // User found
}
