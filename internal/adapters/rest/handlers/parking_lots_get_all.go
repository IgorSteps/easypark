package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/sirupsen/logrus"
)

// ParkingLotGetAllHandler provides a REST Handler implementation to get all parking lots and
// implements http.Handler interface.
type ParkingLotGetAllHandler struct {
	logger *logrus.Logger
	facade ParkingLotFacade
}

// NewParkingLotGetAllHandler returns a new instance of ParkingLotGetAllHandler.
func NewParkingLotGetAllHandler(l *logrus.Logger, f ParkingLotFacade) *ParkingLotGetAllHandler {
	return &ParkingLotGetAllHandler{
		logger: l,
		facade: f,
	}
}

// ServeHTTP handles incoming HTTP request to get all parking lots.
func (s *ParkingLotGetAllHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	lots, err := s.facade.GetAllParkingLots(r.Context())
	if err != nil {
		s.logger.WithError(err).Error("failed to get all parking lotss")

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
	json.NewEncoder(w).Encode(lots)
}
