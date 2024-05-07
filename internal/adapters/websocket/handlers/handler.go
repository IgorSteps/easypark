package handlers

import (
	"net/http"

	"github.com/IgorSteps/easypark/internal/adapters/websocket/client"
	"github.com/IgorSteps/easypark/internal/adapters/websocket/models"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

// TODO for the upgrader:
// 1) Move these values to config?
// 2) Research how we can determine the best buffer sizes...

// upgrader specifies parameters to update HTTP connection to websocket one.
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Disable CORS - NEVER EVER EVER DO THIS IN PRODUCTION!!!
	},
}

// WebsocketHandler handles WebSocket connections.
type WebsocketHandler struct {
	logger *logrus.Logger
	hub    *client.Hub
}

// NewWebsocketHandler returns new instance of WebsocketHandler.
func NewWebsocketHandler(l *logrus.Logger, hub *client.Hub) *WebsocketHandler {
	return &WebsocketHandler{
		logger: l,
		hub:    hub,
	}
}

// ServeHTTP handles upgrading incoming HTTP connections to Websocket connections,
// registers the Client with the Hub and starts Read/Write go routines for the Client.
func (s *WebsocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		s.logger.WithError(err).Error("failed to upgrade HTTP connection to a Websocket one")
		return
	}

	// Parse sender's ID.
	senderID := chi.URLParam(r, "id")
	parsedID, err := uuid.Parse(senderID)
	if err != nil {
		s.logger.WithError(err).Error("failed to parse sender's id")
		return
	}

	// Create a client for the sender.
	client := &client.Client{
		Logger: s.logger,
		Hub:    s.hub,
		Conn:   conn,
		Send:   make(chan *models.Message),
		UserID: parsedID,
	}
	// Register client with the Hub.
	client.Hub.Register <- client

	// Start read/write in their own go-routines.
	go client.Write()
	go client.Read()
}
