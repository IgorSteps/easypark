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

// ParkingSpaceStatusHandler provides a REST Handler implementation to change status of parking space and
// implements http.Handler interface.
type ParkingSpaceStatusHandler struct {
	logger *logrus.Logger
	facade ParkingSpaceFacade
}

// NewParkingSpaceStatusHandler creates new instance of ParkingSpaceStatusHandler.
func NewParkingSpaceStatusHandler(f ParkingSpaceFacade, l *logrus.Logger) *ParkingSpaceStatusHandler {
	return &ParkingSpaceStatusHandler{
		logger: l,
		facade: f,
	}
}

func (s *ParkingSpaceStatusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var updateSpaceStatusRequest models.UpdateParkingSpaceStatus

	spaceID := chi.URLParam(r, "id")

	err := json.NewDecoder(r.Body).Decode(&updateSpaceStatusRequest)
	if err != nil {
		s.logger.WithError(err).Error("failed to decode request")
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	parsedID, err := uuid.Parse(spaceID)
	if err != nil {
		s.logger.WithError(err).Error("failed to parse parking space id")
		http.Error(w, "invalid parking space id", http.StatusBadRequest)
		return
	}

	updatedSpace, err := s.facade.UpdateParkingSpaceStatus(r.Context(), parsedID, updateSpaceStatusRequest.Status)
	if err != nil {
		s.logger.WithError(err).Error("failed to update parking space status")

		switch err.(type) {
		case *repositories.InvalidInputError:
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		case *repositories.NotFoundError:
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		case *repositories.InternalError:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	s.logger.Debug(updatedSpace)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedSpace)
}
