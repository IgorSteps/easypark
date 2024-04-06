package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// DriversParkingRequestsGetHandler provides a REST Handler implementation to get particular driver's parking requests and
// implements http.Handler interface.
type DriversParkingRequestsGetHandler struct {
	logger *logrus.Logger
	facade ParkingRequestFacade
}

// NewParkingRequestsGetHandler returns a new instance of DriversParkingRequestsGetHandler.
func NewDriversParkingRequestsGetHandler(l *logrus.Logger, f ParkingRequestFacade) *DriversParkingRequestsGetHandler {
	return &DriversParkingRequestsGetHandler{
		logger: l,
		facade: f,
	}
}

// ServeHTTP handles incoming HTTP request to get parking requests for a particular driver. Method name matches the http.Handler interface.
func (s *DriversParkingRequestsGetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	driverID := chi.URLParam(r, "id")

	parsedID, err := uuid.Parse(driverID)
	if err != nil {
		s.logger.WithError(err).Error("failed to parse driver id")
		http.Error(w, "invalid driver id", http.StatusBadRequest)
		return
	}

	parkingRequests, err := s.facade.GetDriversParkingRequests(r.Context(), parsedID)
	if err != nil {
		s.logger.WithError(err).Error("failed to get all parking requests")

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
	json.NewEncoder(w).Encode(parkingRequests)
}
