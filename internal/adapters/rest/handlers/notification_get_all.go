package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/sirupsen/logrus"
)

// NotificationGetAllHandler provides a REST Handler implementation to get all notifications and
// implements http.Handler interface.
type NotificationGetAllHandler struct {
	logger *logrus.Logger
	facade NotificationFacade
}

// NewNotificationGetAllHandler returns a new instance of NotificationGetAllHandler.
func NewNotificationGetAllHandler(l *logrus.Logger, f NotificationFacade) *NotificationGetAllHandler {
	return &NotificationGetAllHandler{
		logger: l,
		facade: f,
	}
}

// ServeHTTP handles incoming HTTP request to get all notifications.
// Method name matches the http.Handler interface.
func (s *NotificationGetAllHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	notifications, err := s.facade.GetAllNotifications(r.Context())
	if err != nil {
		s.logger.WithError(err).Error("failed to get all notifications")

		switch err.(type) {
		case *repositories.InternalError:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		default:
			s.logger.WithError(err).Warn("unknown error type")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(notifications)
}
