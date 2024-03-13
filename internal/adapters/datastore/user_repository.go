package datastore

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
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
			"email":    user.Email,
			"username": user.Username,
		}).Error("failed to query for user in the database")
		return false, repositories.NewInternalError("failed to query for user in the database")
	}

	return true, nil // User found
}

// FindByUsername queries DB to find user with given username.
func (s *UserPostgresRepository) FindByUsername(ctx context.Context, username string) (*entities.User, error) {
	var user entities.User

	result := s.DB.WithContext(ctx).Where("username = ?", username).First(&user)
	err := result.Error()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			s.Logger.WithField("username", username).Warn("failed to find user with given username in the database")
			return &entities.User{}, repositories.NewUserNotFoundError(username)
		}

		s.Logger.WithError(err).Error("failed to query for user in the database")
		return &entities.User{}, repositories.NewInternalError("failed to query for user in the database")
	}

	return &user, nil // User found
}
