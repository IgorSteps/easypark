package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/domain/repositories"
	"github.com/sirupsen/logrus"
)

const ctxClaimKey = "claims"

// AuthMiddleware provides middleware for JWT authorisation and RBAC.
type AuthMiddleware struct {
	tokenService repositories.TokenRepository
	logger       *logrus.Logger
}

// NewAuthMiddleware returns new instance of AuthMiddleware.
func NewAuthMiddleware(ts repositories.TokenRepository, l *logrus.Logger) *AuthMiddleware {
	return &AuthMiddleware{
		tokenService: ts,
		logger:       l,
	}
}

// Authorise parses and validates JWT token and extracts claims to the request context for use down the stack.
func (s *AuthMiddleware) Authorise(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		claims, err := s.tokenService.ParseToken(tokenStr)
		if err != nil {
			s.logger.WithError(err).Warn("failed to parse and validate auth token")
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Put user's claims in the request context
		ctx := context.WithValue(r.Context(), ctxClaimKey, claims)

		// Move to the next middleware...
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// RequireRole enforces RBAC by extracting claims from ctx and checking against required role.
func (s *AuthMiddleware) RequireRole(requiredRole entities.UserRole) func(next http.Handler) http.Handler {
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
