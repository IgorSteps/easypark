package handlers

import (
	"net/http"

	usecases "github.com/IgorSteps/easypark/internal/usecases/message"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

// TODO: Move these values to config?
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type WebsocketHandler struct {
	logger         *logrus.Logger
	messageCreator *usecases.CreateMessage
	hub            *Hub
}

func NewWebsocketHandler(l *logrus.Logger, mc *usecases.CreateMessage, hub *Hub) *WebsocketHandler {
	return &WebsocketHandler{
		logger:         l,
		messageCreator: mc,
		hub:            hub,
	}
}

func (s *WebsocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.logger.WithError(err).Error("failed to upgrade HTTP connection to a Websocket one")
		return
	}

	// Parse sender ID.
	senderID := chi.URLParam(r, "id")
	parsedID, err := uuid.Parse(senderID)
	if err != nil {
		s.logger.WithError(err).Error("failed to parse sender's id")
		return
	}

	// Create a client for the sender.
	client := &Client{
		hub:    s.hub,
		conn:   conn,
		send:   make(chan []byte, 256),
		userID: parsedID,
	}
	// Register client with the Hub.
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}
