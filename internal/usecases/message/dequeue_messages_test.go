package usecases_test

import (
	"context"
	"testing"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	usecases "github.com/IgorSteps/easypark/internal/usecases/message"
	mocks "github.com/IgorSteps/easypark/mocks/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDequeueMessageExecute(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	logger, _ := test.NewNullLogger()
	ctx := context.Background()
	driverID := uuid.New()
	mockMessageRepo := &mocks.MessageRepository{}
	usecase := usecases.NewDequeueMessages(logger, mockMessageRepo)

	driversMsges := []entities.Message{
		{
			ID:         uuid.New(),
			SenderID:   uuid.New(),
			ReceiverID: driverID,
			Content:    "hi",
		},
	}
	mockMessageRepo.EXPECT().GetManyForUser(driverID).Return(driversMsges, nil).Once()
	mockMessageRepo.EXPECT().Delete(mock.Anything).Return(nil).Once()

	// --------
	// ACT
	// --------
	msges, err := usecase.Execute(ctx, driverID)

	// --------
	// ASSERT
	// --------
	assert.NoError(t, err)
	assert.Equal(t, len(msges), 1)
	mockMessageRepo.AssertExpectations(t)
}
