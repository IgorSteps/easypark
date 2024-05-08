package datastore_test

import (
	"testing"

	"github.com/IgorSteps/easypark/internal/adapters/datastore"
	"github.com/IgorSteps/easypark/internal/domain/entities"
	mocks "github.com/IgorSteps/easypark/mocks/adapters/datastore"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestMessageRepo_Create(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	logger, _ := test.NewNullLogger()
	db := &mocks.Datastore{}
	repo := datastore.NewMessagePostgresRepository(logger, db)

	msg := &entities.Message{
		ID: uuid.New(),
	}
	db.EXPECT().Create(msg).Return(db).Once()
	db.EXPECT().Error().Return(nil).Once()

	// ----
	// ACT
	// ----
	err := repo.Create(msg)

	// ------
	// ASSERT
	// ------
	assert.NoError(t, err)
}

func TestMessageRepo_GetManyForUser(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	logger, _ := test.NewNullLogger()
	db := &mocks.Datastore{}
	repo := datastore.NewMessagePostgresRepository(logger, db)
	userID := uuid.New()
	var messages []entities.Message

	db.EXPECT().Where("receiver_id = ?", userID).Return(db).Once()
	db.EXPECT().FindAll(&messages).Return(db).Once()
	db.EXPECT().Error().Return(nil).Once()

	// ----
	// ACT
	// ----
	_, err := repo.GetManyForUser(userID)

	// ------
	// ASSERT
	// ------
	assert.NoError(t, err)
}

func TestMessageRepo_Delete(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	logger, _ := test.NewNullLogger()
	db := &mocks.Datastore{}
	repo := datastore.NewMessagePostgresRepository(logger, db)

	messages := []entities.Message{
		{ID: uuid.New()},
	}

	var ids []uuid.UUID
	for _, message := range messages {
		ids = append(ids, message.ID)
	}

	db.EXPECT().Where("id IN ?", ids).Return(db).Once()
	db.EXPECT().Delete(&messages).Return(db).Once()
	db.EXPECT().Error().Return(nil).Once()

	// ----
	// ACT
	// ----
	err := repo.Delete(messages)

	// ------
	// ASSERT
	// ------
	assert.NoError(t, err)
}
