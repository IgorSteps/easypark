package client_test

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/IgorSteps/easypark/internal/adapters/websocket/client"
	"github.com/IgorSteps/easypark/internal/adapters/websocket/models"
	"github.com/IgorSteps/easypark/internal/domain/entities"
	mocks "github.com/IgorSteps/easypark/mocks/adapters/websocket/client"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestHub_RegisterClient(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockFacade := &mocks.MessageFacade{}
	hub := client.NewHub(testLogger, mockFacade)
	testUserID := uuid.New()

	mockClient := &client.Client{
		Hub:    hub,
		UserID: testUserID,
		Send:   make(chan []byte, 256),
	}

	dequeuedMsgs := []entities.Message{
		{
			ID: uuid.New(),
		},
	}
	mockFacade.EXPECT().DequeueMessages(context.TODO(), testUserID).Return(dequeuedMsgs, nil).Once()

	// ----
	// ACT
	// ----
	// Simulate registering a client.
	go func() {
		hub.Register <- mockClient
	}()
	// Process messages.
	go hub.Run()

	// ------
	// ASSERT
	// ------
	// Allow 1 sec for the message to be processed.
	time.Sleep(1 * time.Second)
	assert.Equal(t, mockClient, hub.Clients[mockClient.UserID])
	mockFacade.AssertExpectations(t)
}

func TestHub_UnregisterClient(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockFacade := &mocks.MessageFacade{}
	hub := client.NewHub(testLogger, mockFacade)
	testUserID := uuid.New()

	mockClient := &client.Client{
		Hub:    hub,
		UserID: testUserID,
		Send:   make(chan []byte, 256),
	}
	// Register first.
	hub.Clients[mockClient.UserID] = mockClient

	// ----
	// ACT
	// ----
	// Simulate unregistering a client.
	go func() {
		hub.Unregister <- mockClient
	}()
	// Process messages.
	go hub.Run()

	// ------
	// ASSERT
	// ------
	// Allow 1 sec for the message to be processed.
	time.Sleep(1 * time.Second)
	_, exists := hub.Clients[mockClient.UserID]
	assert.False(t, exists)
	mockFacade.AssertExpectations(t)
}

func TestHub_Broadcast_WithRegisteredReceiver(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockFacade := &mocks.MessageFacade{}
	hub := client.NewHub(testLogger, mockFacade)

	// Sender client.
	testSenderUserID := uuid.New()
	mockSenderClient := &client.Client{
		Hub:    hub,
		UserID: testSenderUserID,
		Send:   make(chan []byte, 256),
	}
	hub.Clients[testSenderUserID] = mockSenderClient

	// Receiver client.
	testReceiverUserID := uuid.New()
	mockReceiverClient := &client.Client{
		Hub:    hub,
		UserID: testSenderUserID,
		Send:   make(chan []byte, 256),
	}
	hub.Clients[testReceiverUserID] = mockReceiverClient

	// Sender's msg.
	message := models.Message{
		SenderID:   testSenderUserID,
		ReceiverID: testReceiverUserID,
		Content:    "Boom",
	}
	marshalledMessage, _ := json.Marshal(message)

	// ----
	// ACT
	// ----
	go func() {
		hub.Broadcast <- marshalledMessage
	}()
	// Process messages.
	go hub.Run()

	// ------
	// ASSERT
	// ------
	// Allow 1 sec for the message to be processed.
	time.Sleep(1 * time.Second)

	receivedMessage := <-mockReceiverClient.Send
	assert.Equal(t, "Boom", string(receivedMessage))

	mockFacade.AssertExpectations(t)
}

func TestHub_Broadcast_WithUnregisteredReceiver(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockFacade := &mocks.MessageFacade{}
	hub := client.NewHub(testLogger, mockFacade)

	// Sender client.
	testSenderUserID := uuid.New()
	mockSenderClient := &client.Client{
		Hub:    hub,
		UserID: testSenderUserID,
		Send:   make(chan []byte, 256),
	}
	hub.Clients[testSenderUserID] = mockSenderClient

	// No registered receiver.
	testReceiverID := uuid.New()

	// Sender's msg.
	message := models.Message{
		SenderID:   testSenderUserID,
		ReceiverID: testReceiverID,
		Content:    "Boom",
	}
	marshalledMessage, _ := json.Marshal(message)

	mockFacade.EXPECT().EnqueueMessage(context.TODO(), testSenderUserID, testReceiverID, message.Content).Return(entities.Message{}, nil).Once()

	// ----
	// ACT
	// ----
	go func() {
		hub.Broadcast <- marshalledMessage
	}()
	// Process messages.
	go hub.Run()

	// ------
	// ASSERT
	// ------
	// Allow 1 sec for the message to be processed.
	time.Sleep(1 * time.Second)
	mockFacade.AssertExpectations(t)
}

func TestHub_Broadcast_EnqueueFailed(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockFacade := &mocks.MessageFacade{}
	hub := client.NewHub(testLogger, mockFacade)

	// Sender client.
	testSenderUserID := uuid.New()
	mockSenderClient := &client.Client{
		Hub:    hub,
		UserID: testSenderUserID,
		Send:   make(chan []byte, 256),
	}
	hub.Clients[testSenderUserID] = mockSenderClient

	// No registered receiver.
	testReceiverID := uuid.New()

	// Sender's msg.
	message := models.Message{
		SenderID:   testSenderUserID,
		ReceiverID: testReceiverID,
		Content:    "Boom",
	}
	marshalledMessage, _ := json.Marshal(message)

	testError := errors.New("boom")
	mockFacade.EXPECT().EnqueueMessage(context.TODO(), testSenderUserID, testReceiverID, message.Content).Return(entities.Message{}, testError).Once()

	// ----
	// ACT
	// ----
	go func() {
		hub.Broadcast <- marshalledMessage
	}()
	// Process messages.
	go hub.Run()

	// ------
	// ASSERT
	// ------
	// Allow 1 sec for the message to be processed.
	time.Sleep(1 * time.Second)
	receivedMessage := <-mockSenderClient.Send
	assert.Equal(t, "failed to enqueue message", string(receivedMessage))
	mockFacade.AssertExpectations(t)
}

func TestHub_Broadcast_DequeueFailed(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockFacade := &mocks.MessageFacade{}
	hub := client.NewHub(testLogger, mockFacade)
	testUserID := uuid.New()

	mockClient := &client.Client{
		Hub:    hub,
		UserID: testUserID,
		Send:   make(chan []byte, 256),
	}

	testError := errors.New("boom")
	mockFacade.EXPECT().DequeueMessages(context.TODO(), testUserID).Return([]entities.Message{}, testError).Once()

	// ----
	// ACT
	// ----
	// Simulate registering a client.
	go func() {
		hub.Register <- mockClient
	}()
	// Process messages.
	go hub.Run()

	// ------
	// ASSERT
	// ------
	// Allow 1 sec for the message to be processed.
	time.Sleep(1 * time.Second)
	receivedMessage := <-mockClient.Send
	assert.Equal(t, "boom", string(receivedMessage))
	mockFacade.AssertExpectations(t)
}

func TestHub_Broadcast_FullSendBuffer(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockFacade := &mocks.MessageFacade{}
	hub := client.NewHub(testLogger, mockFacade)

	// Sender client.
	testSenderUserID := uuid.New()
	mockSenderClient := &client.Client{
		Hub:    hub,
		UserID: testSenderUserID,
		Send:   make(chan []byte, 256),
	}
	hub.Clients[testSenderUserID] = mockSenderClient

	// Receiver client
	testReceiverUserID := uuid.New()
	mockReceiverClient := &client.Client{
		Hub:    hub,
		UserID: testReceiverUserID,
		Send:   make(chan []byte, 1), // Make it very small.
	}
	mockReceiverClient.Send <- []byte("boom") // Fill the buffer.
	hub.Clients[testReceiverUserID] = mockReceiverClient

	// Sender's msg.
	message := models.Message{
		SenderID:   testSenderUserID,
		ReceiverID: testReceiverUserID,
		Content:    "Boom",
	}
	marshalledMessage, _ := json.Marshal(message)

	// ----
	// ACT
	// ----
	// Simulate registering a client.
	go func() {
		hub.Broadcast <- marshalledMessage
	}()
	// Process messages.
	go hub.Run()

	// ------
	// ASSERT
	// ------
	// Allow 1 sec for the message to be processed.
	time.Sleep(1 * time.Second)

	assert.NotContains(t, hub.Clients, testReceiverUserID, "Receiver client should be unregistered")
	_, ok := <-mockReceiverClient.Send
	assert.False(t, ok, "Receiver client send channel should be closed")

	mockFacade.AssertExpectations(t)
}
