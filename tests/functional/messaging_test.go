package functional

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/IgorSteps/easypark/internal/adapters/websocket/models"
	"github.com/IgorSteps/easypark/tests/functional/client"
	"github.com/IgorSteps/easypark/tests/functional/utils"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/suite"
)

type TestMessagingSuite struct {
	client.RestClientSuite
}

func (s *TestMessagingSuite) TestDriverToAdminChat() {
	// --------
	// ASSEMBLE
	// --------
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	senderDriver, _ := utils.CreateAndLoginDriver(ctx, &s.RestClientSuite, nil)
	_ = utils.CreateAndLoginAdmin(ctx, &s.RestClientSuite)

	// Driver WebSocket connection
	driverURL := "ws://localhost:8081/ws/" + senderDriver.ID.String()
	driverConn, _, err := websocket.DefaultDialer.Dial(driverURL, nil)
	s.Require().NoError(err, "Failed to establish WebSocket connection for driver")
	defer driverConn.Close()

	// Admin WebSocket connection
	adminURL := "ws://localhost:8081/ws/" + "a131a9a0-8d09-4166-b6fc-f8a08ba549e9"
	adminConn, _, err := websocket.DefaultDialer.Dial(adminURL, nil)
	s.Require().NoError(err, "Failed to establish WebSocket connection for admin")
	defer adminConn.Close()

	// Prepare message
	content := "Hello WebSocket"
	message := models.Message{
		SenderID:   senderDriver.ID,
		ReceiverID: uuid.MustParse("a131a9a0-8d09-4166-b6fc-f8a08ba549e9"),
		Content:    content,
	}
	msgData, err := json.Marshal(message)
	s.Require().NoError(err, "Failed to serialize message")

	// ------
	// ACT
	// ------
	// Send message from driver to admin
	err = driverConn.WriteMessage(websocket.TextMessage, msgData)
	s.Require().NoError(err, "Failed to send message from driver")

	// Receive message on admin's connection
	_, response, err := adminConn.ReadMessage()
	s.Require().NoError(err, "Failed to read message on admin's connection")

	var receivedMsg models.Message
	err = json.Unmarshal(response, &receivedMsg)
	s.Require().NoError(err)
	// ------
	// ASSERT
	// ------
	s.Require().Equal(message, receivedMsg, "Message received doesn't match message that was sent")
}

func (s *TestMessagingSuite) TestFullConversation() {
	// --------
	// ASSEMBLE
	// --------
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	senderDriver, _ := utils.CreateAndLoginDriver(ctx, &s.RestClientSuite, nil)
	_ = utils.CreateAndLoginAdmin(ctx, &s.RestClientSuite)

	// Establish connections for admin and driver
	driverConn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8081/ws/"+senderDriver.ID.String(), nil)
	s.Require().NoError(err, "Failed to establish WebSocket connection for driver")
	defer driverConn.Close()

	adminConn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8081/ws/"+"a131a9a0-8d09-4166-b6fc-f8a08ba549e9", nil)
	s.Require().NoError(err, "Failed to establish WebSocket connection for admin")
	defer adminConn.Close()

	adminUserID := uuid.MustParse("a131a9a0-8d09-4166-b6fc-f8a08ba549e9")
	// Simulate conversation
	conversation := []models.Message{
		{SenderID: senderDriver.ID, ReceiverID: adminUserID, Content: "Hello admin"},
		{SenderID: adminUserID, ReceiverID: senderDriver.ID, Content: "Hello driver, what?"},
		{SenderID: senderDriver.ID, ReceiverID: adminUserID, Content: "I need help."},
		{SenderID: adminUserID, ReceiverID: senderDriver.ID, Content: "No."},
	}

	// --------
	// ACT
	// --------
	for _, msg := range conversation {
		s.T().Log(msg)
		senderConn := driverConn
		receiverConn := adminConn
		if msg.SenderID == adminUserID {
			senderConn = adminConn
			receiverConn = driverConn
		}

		// Serialize message to JSON
		msgData, err := json.Marshal(msg)
		s.Require().NoError(err, "Failed to serialize message")

		// Send message
		err = senderConn.WriteMessage(websocket.TextMessage, msgData)
		s.Require().NoError(err, "Failed to send message")

		// Receive message on the sender connection
		_, senderResponse, err := senderConn.ReadMessage()
		s.Require().NoError(err, "Failed to read message from sender")

		// Receive message on the receiver connection
		_, receiverResponse, err := receiverConn.ReadMessage()
		s.Require().NoError(err, "Failed to read message from receiver")

		// --------
		// ASSERT
		// --------
		var senderReceivedMsg models.Message
		var receiverReceivedMsg models.Message

		err = json.Unmarshal(senderResponse, &senderReceivedMsg)
		s.Require().NoError(err)

		err = json.Unmarshal(receiverResponse, &receiverReceivedMsg)
		s.Require().NoError(err)

		// Ensure both sender and receiver get the correct messages
		s.Require().Equal(msg.Content, senderReceivedMsg.Content, "Sender message content does not match")
		s.Require().Equal(msg.Content, receiverReceivedMsg.Content, "Receiver message content does not match")
	}
}

func (s *TestMessagingSuite) TestMessagesAreDequeued() {
	// --------
	// ASSEMBLE
	// --------
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	senderDriver, _ := utils.CreateAndLoginDriver(ctx, &s.RestClientSuite, nil)
	_ = utils.CreateAndLoginAdmin(ctx, &s.RestClientSuite)

	// Establish connections for driver
	driverConn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8081/ws/"+senderDriver.ID.String(), nil)
	s.Require().NoError(err, "Failed to establish WebSocket connection for driver")
	defer driverConn.Close()

	adminUserID := uuid.MustParse("a131a9a0-8d09-4166-b6fc-f8a08ba549e9")

	driversMessages := []models.Message{
		{SenderID: senderDriver.ID, ReceiverID: adminUserID, Content: "Hello admin"},
		{SenderID: senderDriver.ID, ReceiverID: adminUserID, Content: "I need help."},
		{SenderID: senderDriver.ID, ReceiverID: adminUserID, Content: "Hello."},
		{SenderID: senderDriver.ID, ReceiverID: adminUserID, Content: "Answer me."},
	}

	for _, msg := range driversMessages {
		// Serialize message to JSON
		msgData, err := json.Marshal(msg)
		s.Require().NoError(err, "Failed to serialize message")

		// Send message
		err = driverConn.WriteMessage(websocket.TextMessage, msgData)
		s.Require().NoError(err, "Failed to send message")

		// Don't do anything, messages should be persisted in the db.
	}

	// Wait for server to process and queue messages
	time.Sleep(3 * time.Second)

	// Now connect the admin
	adminConn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8081/ws/"+"a131a9a0-8d09-4166-b6fc-f8a08ba549e9", nil)
	s.Require().NoError(err, "Failed to establish WebSocket connection for admin")
	defer adminConn.Close()

	// --------
	// ACT
	// --------
	for _, msg := range driversMessages {
		// Receive message on admin's connection
		_, response, err := adminConn.ReadMessage()
		s.Require().NoError(err, "Failed to read message on admin's connection")

		// --------
		// ASSERT
		// --------
		var message models.Message
		err = json.Unmarshal(response, &message)
		s.Require().Equal(msg, message, "Message received doesn't match messages that were sent")
	}
}

func TestMessagingSuiteInit(t *testing.T) {
	suite.Run(t, new(TestMessagingSuite))
}
