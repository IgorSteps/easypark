package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/IgorSteps/easypark/internal/adapters/rest/models"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/sirupsen/logrus"
)

// UserCreateHandler provides a REST Handler implementation to create users and
// implements http.Handler interface.
type UserCreateHandler struct {
	logger *logrus.Logger
	facade UserFacade
}

// NewUserCreateHandler creates new instance of UserCreateHandler.
func NewUserCreateHandler(f UserFacade, l *logrus.Logger) *UserCreateHandler {
	return &UserCreateHandler{
		logger: l,
		facade: f,
	}
}

// ServeHTTP handles incoming HTTP request to create users.
func (s *UserCreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request models.UserCreationRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		s.logger.Error("failed to decode user creation request: ", err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	domainUser := request.ToDomain()
	err = s.facade.CreateUser(r.Context(), domainUser)
	if err != nil {
		s.logger.WithError(err).Error("failed to create user")

		switch err.(type) {
		case *repositories.UserAlreadyExistsError:
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
	resp := models.CreateUserResponse{Message: "user created successfully"}
	json.NewEncoder(w).Encode(resp)
}
