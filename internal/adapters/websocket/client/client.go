package client

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type Client struct {
	Logger *logrus.Logger
	Hub    *Hub
	Conn   *websocket.Conn
	Send   chan []byte
	UserID uuid.UUID
}

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
		c.Hub.Broadcast <- message
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
		err := c.Conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			c.Logger.WithError(err).WithField("client id", c.UserID).Error("failed to send websocket message")
			break
		}
	}
}
