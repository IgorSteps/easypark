package client

import (
	"encoding/json"

	"github.com/IgorSteps/easypark/internal/adapters/websocket/models"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

// Client is setup for every websocket connection, it acts as 'transmitter' between the Hub and Websocket connection, ie.
// it transfers client's outgoing messages to the WS Connection and incoming messages to the Hub.
type Client struct {
	Logger *logrus.Logger
	Hub    *Hub
	Conn   *websocket.Conn
	Send   chan *models.Message
	UserID uuid.UUID
}

// Read transfers incoming messages from the Websocket connection to the Hub.
//
// Must run as a go-routine for every connection to avoid blocking other operations, like writing or handling other conns.
func (c *Client) Read() {
	// TODO: Add timeout and rate limiting for message reading.
	defer func() {
		c.Hub.unregister <- c
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

		// Send to the hub
		c.Hub.broadcast <- &target
	}
}

// Write transfers outgoing messages the Hub to the Websocket connection.
//
// Run as a go-routine for every connection to avoid blocking other operations, like writing or handling other conns.
func (c *Client) Write() {
	// TODO:
	// 1) Add a timeout or rate limiting for message sending.
	// 2) Add ping/pong mechanism - find out if this required? How our UI can handle this?
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
