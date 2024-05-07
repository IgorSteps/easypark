package routes

import (
	"net/http"

	"github.com/IgorSteps/easypark/internal/adapters/websocket/client"
	"github.com/IgorSteps/easypark/internal/adapters/websocket/handlers"
	chiLogger "github.com/chi-middleware/logrus-logger"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

// Middleware defines an interfaces for middleware that authorises users' tokens.
type Middleware interface {
	Authorise(next http.Handler) http.Handler
}

// NewRouter constructs routes for our Websocket API.
func NewRouter(logger *logrus.Logger, hub *client.Hub) chi.Router {
	router := chi.NewRouter()
	router.Use(chiLogger.Logger("router", logger))

	router.Group(func(r chi.Router) {
		//r.Use(middleware.Authorise)
		r.Method(http.MethodGet, "/ws/{id}", handlers.NewWebsocketHandler(logger, hub))
	})

	return router
}
