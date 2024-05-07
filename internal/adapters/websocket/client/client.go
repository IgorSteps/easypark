package client

import (
	"encoding/json"

	"github.com/IgorSteps/easypark/internal/adapters/websocket/models"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type Client struct {
	Logger *logrus.Logger
	Hub    *Hub
	Conn   *websocket.Conn
	Send   chan *models.Message
	UserID uuid.UUID
}

// TODO: Add a timeout or rate limiting for message sending to prevent abuse or resource exhaustion.
func (c *Client) ReadPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				c.Logger.WithError(err).Error("unexpected websocket closure")
			}
			break
		}

		// Deserialize JSON.
		var target models.Message
		err = json.Unmarshal(message, &target)
		if err != nil {
			c.Logger.WithError(err).WithField("client id", c.UserID).Error("failed to unmarshal websocket message")
			break
		}
		c.Logger.WithField("data", target).Debug("reading message")
		c.Hub.Broadcast <- &target
	}
}

// TODO: Add a timeout or rate limiting for message sending to prevent abuse or resource exhaustion.
func (c *Client) WritePump() {
	defer func() {
		c.Conn.Close()
	}()
	for {
		message, ok := <-c.Send
		if !ok {
			c.Logger.WithField("client id", c.UserID).Debug("hub closed the channel for client")
			c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
			break
		}

		// Serialize message to JSON.
		data, err := json.Marshal(message)
		if err != nil {
			c.Logger.WithError(err).WithField("client id", c.UserID).Error("failed to marshal websocket message")
			break
		}
		c.Logger.WithField("message", message).Debug("writing message")

		w, err := c.Conn.NextWriter(websocket.TextMessage)
		if err != nil {
			c.Logger.WithError(err).WithField("client id", c.UserID).Error("failed to setup a write for next message")
			break
		}
		w.Write(data)

		if err := w.Close(); err != nil {
			c.Logger.WithError(err).WithField("client id", c.UserID).Error("failed to close websocket write")
			break
		}
	}
}
