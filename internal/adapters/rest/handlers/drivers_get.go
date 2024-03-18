package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/sirupsen/logrus"
)

type DriverUsersGetHander struct {
	logger *logrus.Logger
	facade UserFacade
}

func NewDriverUsersGetHandler(l *logrus.Logger, f UserFacade) *DriverUsersGetHander {
	return &DriverUsersGetHander{
		logger: l,
		facade: f,
	}
}

func (s *DriverUsersGetHander) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	users, err := s.facade.GetAllDriverUsers(r.Context())
	if err != nil {
		s.logger.WithError(err).Error("failed to get all drivers")

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
	json.NewEncoder(w).Encode(users)
}
