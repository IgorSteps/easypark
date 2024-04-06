package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/sirupsen/logrus"
)

type ParkingRequestsGetHandler struct {
	logger *logrus.Logger
	facade ParkingRequestFacade
}

func NewParkingRequestsGetHandler(l *logrus.Logger, f ParkingRequestFacade) *ParkingRequestsGetHandler {
	return &ParkingRequestsGetHandler{
		logger: l,
		facade: f,
	}
}

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
