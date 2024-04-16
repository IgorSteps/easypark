package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/IgorSteps/easypark/internal/adapters/rest/models"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type NotificationCreateHandler struct {
	logger *logrus.Logger
	facade NotificationFacade
}

func NewNotificationCreateHandler(l *logrus.Logger, f NotificationFacade) *NotificationCreateHandler {
	return &NotificationCreateHandler{
		logger: l,
		facade: f,
	}
}

func (s *NotificationCreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var createReq models.CreateNotificationRequest

	err := json.NewDecoder(r.Body).Decode(&createReq)
	if err != nil {
		s.logger.Error("failed to decode notification creation request: ", err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	driverID := chi.URLParam(r, "id")
	parsedID, err := uuid.Parse(driverID)
	if err != nil {
		s.logger.WithError(err).Error("failed to parse driver id")
		http.Error(w, "invalid driver id", http.StatusBadRequest)
		return
	}
	notification, err := s.facade.CreateNotification(
		r.Context(),
		parsedID,
		createReq.ParkingRequestID,
		createReq.ParkingSpaceID,
		createReq.Location,
		createReq.NotificationType,
	)
	if err != nil {
		s.logger.WithError(err).Error("failed to create a notification")

		switch err.(type) {
		case *repositories.InvalidInputError:
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		case *repositories.InternalError:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		default:
			s.logger.WithError(err).Warn("unknown error type")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(notification)
}
