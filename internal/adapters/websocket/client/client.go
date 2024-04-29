package client

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type Client struct {
	logger *logrus.Logger
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
				c.logger.WithError(err).Error("unexpected websocket closure")
			}
			break
		}
		c.Hub.Broadcast <- message
	}
}

func (c *Client) WritePump() {
	defer func() {
		c.Conn.Close()
	}()
	for {
		message, ok := <-c.Send
		if !ok {
			c.logger.Info("hub closed the channel")
			c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
			break
		}
		c.Conn.WriteMessage(websocket.TextMessage, message)
	}
}
