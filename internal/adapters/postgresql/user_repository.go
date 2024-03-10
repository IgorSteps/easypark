package postgresql

import (
	"context"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/sirupsen/logrus"
)

type PostgreSQLUserRepository struct {
	Logger *logrus.Logger
	DB     DBHandler
}

func NewPostgreSQLUserRepository(db DBHandler, lgr *logrus.Logger) *PostgreSQLUserRepository {
	return &PostgreSQLUserRepository{
		Logger: lgr,
		DB:     db,
	}
}

func (s *PostgreSQLUserRepository) CreateUser(ctx context.Context, user entities.User) (entities.User, error) {
	result := s.DB.Create(ctx, &user)
	if result.Error != nil {
		s.Logger.WithError(result.Error).Error("failed to insert user into the database")

		return entities.User{}, result.Error
	}

	return user, nil
}
