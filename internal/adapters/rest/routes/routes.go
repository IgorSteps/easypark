package routes

import (
	"net/http"

	"github.com/sirupsen/logrus"

	lgr "github.com/chi-middleware/logrus-logger"
	"github.com/go-chi/chi/v5"
)

// HandlerFactory defines an interface for creating different REST handlers.
type HandlerFactory interface {
	UserCreate() http.Handler
	UserAuthorise() http.Handler
}

// NewRouter constructs routes for our REST API.
func NewRouter(handlerFactory HandlerFactory, logger *logrus.Logger) chi.Router {
	router := chi.NewRouter()
	router.Use(lgr.Logger("router", logger))

	router.Method(http.MethodPost, "/register", handlerFactory.UserCreate())
	router.Method(http.MethodPost, "/login", handlerFactory.UserAuthorise())

	return router
}
