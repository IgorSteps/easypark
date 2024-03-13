package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// HandlerFactory defines an interface for creating different REST handlers.
type HandlerFactory interface {
	UserCreate() http.Handler
}

// NewRouter constructs routes for our REST API.
func NewRouter(handlerFactory HandlerFactory) chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Method(http.MethodPost, "/register", handlerFactory.UserCreate())

	return router
}
