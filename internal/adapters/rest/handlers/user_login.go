package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/IgorSteps/easypark/internal/adapters/rest/models"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/sirupsen/logrus"
)

// UserLoginHandler provides a REST Handler implementation to login users and
// implements http.Handler interface.
type UserLoginHandler struct {
	logger *logrus.Logger
	facade UserFacade
}

// NewUserLoginHandler creates new instance of UserCreateHandler.
func NewUserLoginHandler(f UserFacade, l *logrus.Logger) *UserLoginHandler {
	return &UserLoginHandler{
		logger: l,
		facade: f,
	}
}

// ServeHTTP handles incoming HTTP request to login users.
func (s *UserLoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var request models.LoginUserRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		s.logger.Error("failed to decode user login request: ", err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	token, err := s.facade.AuthoriseUser(r.Context(), request.Username, request.Password)
	if err != nil {
		s.logger.WithError(err).Error("failed to login user")
		switch err.(type) {
		case *repositories.InvalidCredentialsError:
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		case *repositories.UserNotFoundError:
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		case *repositories.InternalError:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		default:
			s.logger.Error("unknown error type")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	resp := models.LoginUserResponse{Message: "User logged in successfully", Token: token}
	json.NewEncoder(w).Encode(resp)
}
