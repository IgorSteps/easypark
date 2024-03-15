package routes

import (
	"net/http"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/sirupsen/logrus"

	lgr "github.com/chi-middleware/logrus-logger"
	"github.com/go-chi/chi/v5"
)

// HandlerFactory defines an interface for creating different REST handlers.
type HandlerFactory interface {
	UserCreate() http.Handler
	UserAuthorise() http.Handler
}

// RequestAuthoriser defines an interfaces for middleware that authorises users' tokens.
//
// This could be a factory similar to handlers, but we aren't using many middlewares
// that require same dependencies.
type RequestAuthoriser interface {
	Authorise(next http.Handler) http.Handler
	RequireRole(requiredRole entities.UserRole) func(next http.Handler) http.Handler
}

// NewRouter constructs routes for our REST API.
func NewRouter(handlerFactory HandlerFactory, requestAuthoriser RequestAuthoriser, logger *logrus.Logger) chi.Router {
	router := chi.NewRouter()
	router.Use(lgr.Logger("router", logger))

	// Public routes
	router.Method(http.MethodPost, "/register", handlerFactory.UserCreate())
	router.Method(http.MethodPost, "/login", handlerFactory.UserAuthorise())

	// Driver routes
	router.Group(func(r chi.Router) {
		r.Use(requestAuthoriser.Authorise, requestAuthoriser.RequireRole(entities.RoleDriver))
		// Placeholder:
		r.Get("/driver", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Welcome, Driver!"))
		})
	})

	// Admin routes
	router.Group(func(r chi.Router) {
		r.Use(requestAuthoriser.Authorise, requestAuthoriser.RequireRole(entities.RoleAdmin))
		// Placeholder:
		r.Get("/admin", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Welcome, Admin!"))
		})
	})

	return router
}
