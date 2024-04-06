package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/IgorSteps/easypark/internal/adapters/rest/models"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/sirupsen/logrus"
)

// DriverCreateHandler provides a REST Handler implementation to create driver users and
// implements http.Handler interface.
type DriverCreateHandler struct {
	logger *logrus.Logger
	facade UserFacade
}

// NewDriverCreateHandler creates new instance of DriverCreateHandler.
func NewDriverCreateHandler(f UserFacade, l *logrus.Logger) *DriverCreateHandler {
	return &DriverCreateHandler{
		logger: l,
		facade: f,
	}
}

// ServeHTTP handles incoming HTTP request to create users.
func (s *DriverCreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request models.UserCreationRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		s.logger.Error("failed to decode user creation request: ", err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	domainUser := request.ToDomain()
	driver, err := s.facade.CreateDriver(r.Context(), domainUser)
	if err != nil {
		s.logger.WithError(err).Error("failed to create user")

		switch err.(type) {
		case *repositories.ResourceAlreadyExistsError:
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
	json.NewEncoder(w).Encode(driver)
}
