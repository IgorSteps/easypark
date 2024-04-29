package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/IgorSteps/easypark/internal/adapters/websocket/models"
	usecases "github.com/IgorSteps/easypark/internal/usecases/message"
	"github.com/gorilla/websocket"
	goriallaWebsocket "github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var upgrader = goriallaWebsocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type WebsocketHandler struct {
	logger         *logrus.Logger
	messageCreator *usecases.CreateMessage
}

func NewWebsocketHandler(l *logrus.Logger, mc *usecases.CreateMessage) *WebsocketHandler {
	return &WebsocketHandler{
		logger:         l,
		messageCreator: mc,
	}
}

func (s *WebsocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		return
	}
	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			return
		}

		// Parse the JSON message
		var msg models.ChatMessage
		err = json.Unmarshal(p, &msg)
		if err != nil {
			// Send an error message back to the client
			errMsg, _ := json.Marshal(map[string]string{"error": "Invalid message format"})
			conn.WriteMessage(websocket.TextMessage, errMsg)
			continue
		}

		// Process the message
		_, err = s.messageCreator.Execute(context.TODO(), msg.SenderID, msg.ReceiverID, msg.Content)
		if err != nil {
			// TODO: Proper error handling
			errMsg, _ := json.Marshal(map[string]string{"error": err.Error()})
			conn.WriteMessage(websocket.TextMessage, errMsg)
			continue
		}

		if err = conn.WriteMessage(messageType, p); err != nil {
			return
		}
	}
}
