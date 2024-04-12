package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type ParkingSpaceGetSingleHandler struct {
	logger *logrus.Logger
	facade ParkingSpaceFacade
}

func NewParkingSpaceGetSingleHandler(l *logrus.Logger, f ParkingSpaceFacade) *ParkingSpaceGetSingleHandler {
	return &ParkingSpaceGetSingleHandler{
		logger: l,
		facade: f,
	}
}

func (s *ParkingSpaceGetSingleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	parkingSpaceID := chi.URLParam(r, "id")
	parsedID, err := uuid.Parse(parkingSpaceID)
	if err != nil {
		s.logger.WithError(err).Error("failed to parse parking space id")
		http.Error(w, "invalid parking space id", http.StatusBadRequest)
		return
	}

	parkingSpace, err := s.facade.GetSingleParkingSpace(r.Context(), parsedID)
	if err != nil {
		s.logger.WithError(err).Error("failed to get single parking space")

		switch err.(type) {
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
	json.NewEncoder(w).Encode(parkingSpace)
}
