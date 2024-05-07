package usecasefacades_test

import (
	"context"
	"testing"

	"github.com/IgorSteps/easypark/internal/adapters/usecasefacades"
	"github.com/IgorSteps/easypark/internal/domain/entities"
	mocks "github.com/IgorSteps/easypark/mocks/adapters/usecasefacades"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestMessageFacade_Enqueue(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	enqueuer := &mocks.MessageEnqueuer{}
	dequeuer := &mocks.MessageDequeuer{}
	facade := usecasefacades.NewMessageFacade(enqueuer, dequeuer)
	testDriverID := uuid.New()
	testAdminID := uuid.New()
	content := "aaaa"

	testCtx := context.Background()
	msg := entities.Message{
		ID: uuid.New(),
	}
	enqueuer.EXPECT().Execute(testCtx, testDriverID, testAdminID, content).Return(msg, nil).Once()

	// --------
	// ACT
	// --------
	actualMsg, err := facade.EnqueueMessage(testCtx, testDriverID, testAdminID, content)

	// --------
	// ASSERT
	// --------
	assert.NoError(t, err)
	assert.Equal(t, actualMsg, msg)
}

func TestMessageFacade_Dequeue(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	enqueuer := &mocks.MessageEnqueuer{}
	dequeuer := &mocks.MessageDequeuer{}
	facade := usecasefacades.NewMessageFacade(enqueuer, dequeuer)
	testDriverID := uuid.New()

	testCtx := context.Background()
	msgs := []entities.Message{
		{ID: uuid.New()},
	}
	dequeuer.EXPECT().Execute(testCtx, testDriverID).Return(msgs, nil).Once()

	// --------
	// ACT
	// --------
	actualMsgs, err := facade.DequeueMessages(testCtx, testDriverID)

	// --------
	// ASSERT
	// --------
	assert.NoError(t, err)
	assert.Equal(t, len(actualMsgs), 1)
}
