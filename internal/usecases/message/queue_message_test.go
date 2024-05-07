package usecases_test

import (
	"context"
	"testing"
	"time"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	usecases "github.com/IgorSteps/easypark/internal/usecases/message"
	mocks "github.com/IgorSteps/easypark/mocks/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestQueueMessageExecute(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	logger, _ := test.NewNullLogger()
	ctx := context.Background()
	driverID := uuid.New()
	adminID := uuid.New()
	driver := entities.User{ID: driverID, Role: "driver"}
	admin := entities.User{ID: adminID, Role: "admin"}
	mockUserRepo := &mocks.UserRepository{}
	mockMessageRepo := &mocks.MessageRepository{}
	usecase := usecases.NewQueueMessage(logger, mockUserRepo, mockMessageRepo)

	mockUserRepo.EXPECT().GetDriverByID(ctx, driverID, mock.Anything).Return(nil).Once().Run(func(args mock.Arguments) {
		arg := args.Get(2).(*entities.User)
		*arg = driver
	})
	mockUserRepo.EXPECT().GetDriverByID(ctx, adminID, mock.Anything).Return(nil).Once().Run(func(args mock.Arguments) {
		arg := args.Get(2).(*entities.User)
		*arg = admin
	})
	content := "hi"
	expectedMessage := entities.Message{
		ID:         uuid.New(),
		SenderID:   driverID,
		ReceiverID: adminID,
		Content:    content,
		Delivered:  false,
		Timestamp:  time.Now(),
	}
	mockMessageRepo.EXPECT().Create(mock.AnythingOfType("*entities.Message")).Return(nil)

	// --------
	// ACT
	// --------
	message, err := usecase.Execute(ctx, driverID, adminID, content)

	// --------
	// ASSERT
	// --------
	assert.NoError(t, err)
	assert.Equal(t, expectedMessage.Content, message.Content)
	mockMessageRepo.AssertExpectations(t)
	mockUserRepo.AssertExpectations(t)
}
