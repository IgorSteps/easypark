package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const ctxClaimKey = "claims"

// StatusChecker provides an interface to check if the driver is banned.
// Implemented by the CheckDriverStatus usecase.
//
// TODO: Consider if we need to wrap it in a facade.
type StatusChecker interface {
	Execute(ctx context.Context, id uuid.UUID) (bool, error)
}

// Middleware provides middleware for JWT authorisation and RBAC.
type Middleware struct {
	tokenService  repositories.TokenRepository
	logger        *logrus.Logger
	statusChekcer StatusChecker
}

// NewMiddleware returns new instance of Middleware.
func NewMiddleware(ts repositories.TokenRepository, l *logrus.Logger, sc StatusChecker) *Middleware {
	return &Middleware{
		tokenService:  ts,
		logger:        l,
		statusChekcer: sc,
	}
}

// Authorise parses and validates JWT token and extracts claims to the request context for use down the stack.
func (s *Middleware) Authorise(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		claims, err := s.tokenService.ParseToken(tokenStr)
		if err != nil {
			s.logger.WithError(err).Warn("failed to parse and validate auth token")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Put user's claims in the request context.
		ctx := context.WithValue(r.Context(), ctxClaimKey, claims)

		// Move to the next middleware...
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// RequireRole enforces RBAC by extracting claims from ctx and checking against required role.
func (s *Middleware) RequireRole(requiredRole entities.UserRole) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Convert ctx claims to our domain claims type.
			claims, ok := r.Context().Value(ctxClaimKey).(*entities.Claims)

			// TODO: Could probably split into 2 conditionals to return more meaningful error msgs
			if !ok || claims.Role != string(requiredRole) {
				s.logger.WithField("required role", requiredRole).Warn("access is forbidden")
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			// Move to the next middleware...
			next.ServeHTTP(w, r)
		})
	}
}

// CheckStatus checks whether user sending the request is banned.
func (s *Middleware) CheckStatus(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Ignore "ok" flag, becasue we assume it was converted in the previous handler with no errors.
		claims, _ := r.Context().Value(ctxClaimKey).(*entities.Claims)

		userID := claims.UserID
		isBanned, err := s.statusChekcer.Execute(r.Context(), userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if isBanned {
			s.logger.WithField("id", userID).Warn("banned driver tried to access the service")
			http.Error(w, "Account is banned.", http.StatusForbidden)
			return
		}

		// Move to the next middleware...
		next.ServeHTTP(w, r)
	})
}
