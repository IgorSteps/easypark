package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// ParkingLotGetSingleHandler provides a REST Handler implementation to get single parking lot and
// implements http.Handler interface.
type ParkingLotGetSingleHandler struct {
	logger *logrus.Logger
	facade ParkingLotFacade
}

// NewParkingLotGetSingleHandler creates new instance of ParkingLotGetSingleHandler.
func NewParkingLotGetSingleHandler(f ParkingLotFacade, l *logrus.Logger) *ParkingLotGetSingleHandler {
	return &ParkingLotGetSingleHandler{
		logger: l,
		facade: f,
	}
}

// ServeHTTP handles incoming HTTP request to get a single parking lot.
// Method name matches the http.Handler interface.
func (s *ParkingLotGetSingleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	lotID := chi.URLParam(r, "id")
	parsedID, err := uuid.Parse(lotID)
	if err != nil {
		s.logger.WithError(err).Error("failed to parse parking lot id")
		http.Error(w, "invalid parking lot id", http.StatusBadRequest)
		return
	}

	lot, err := s.facade.GetSingleParkingLot(r.Context(), parsedID)
	if err != nil {
		s.logger.WithError(err).Error("failed to get single parking lot")

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
		default:
			s.logger.Warn("unknown error type, returning internal server error")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(lot)
}
