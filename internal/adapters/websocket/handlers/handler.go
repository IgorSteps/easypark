package handlers

import (
	"net/http"

	"github.com/IgorSteps/easypark/internal/adapters/websocket/client"
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
	logger *logrus.Logger
	hub    *client.Hub
}

func NewWebsocketHandler(l *logrus.Logger, hub *client.Hub) *WebsocketHandler {
	return &WebsocketHandler{
		logger: l,
		hub:    hub,
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
	client := &client.Client{
		Hub:    s.hub,
		Conn:   conn,
		Send:   make(chan []byte, 256),
		UserID: parsedID,
	}
	// Register client with the Hub.
	client.Hub.Register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WritePump()
	go client.ReadPump()
}
