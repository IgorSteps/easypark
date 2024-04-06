package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/sirupsen/logrus"
)

// ParkingRequestsGetHandler provides a REST Handler implementation to get all parking requests and
// implements http.Handler interface.
type ParkingRequestsGetHandler struct {
	logger *logrus.Logger
	facade ParkingRequestFacade
}

// NewParkingRequestsGetHandler returns a new instance of ParkingRequestsGetHandler.
func NewParkingRequestsGetHandler(l *logrus.Logger, f ParkingRequestFacade) *ParkingRequestsGetHandler {
	return &ParkingRequestsGetHandler{
		logger: l,
		facade: f,
	}
}

// ServeHTTP handles incoming HTTP request to get all parking requests. Method name matches the http.Handler interface.
func (s *ParkingRequestsGetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	parkingRequests, err := s.facade.GetAllParkingRequests(r.Context())
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
