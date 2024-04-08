package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// DeleteParkingLotHandler provides a REST Handler implementation to delete a parking lot and
// implements http.Handler interface.
type DeleteParkingLotHandler struct {
	logger *logrus.Logger
	facade ParkingLotFacade
}

// NewDeleteParkingLotHandler returns a new instance of DeleteParkingLotHandler.
func NewDeleteParkingLotHandler(l *logrus.Logger, f ParkingLotFacade) *DeleteParkingLotHandler {
	return &DeleteParkingLotHandler{
		logger: l,
		facade: f,
	}
}

// ServeHTTP handles incoming HTTP request to delete a parking lot.
func (s *DeleteParkingLotHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	lotID := chi.URLParam(r, "id")
	parsedID, err := uuid.Parse(lotID)
	if err != nil {
		s.logger.WithError(err).Error("failed to parse lot id")
		http.Error(w, "invalid lot id", http.StatusBadRequest)
		return
	}

	err = s.facade.DeleteParkingLot(r.Context(), parsedID)
	if err != nil {
		s.logger.WithError(err).Error("failed to delete a parking lot")

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

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("successfully deleted parking lot")
}
