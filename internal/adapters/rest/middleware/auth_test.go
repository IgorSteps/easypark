package middleware_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/IgorSteps/easypark/internal/adapters/rest/middleware"
	"github.com/IgorSteps/easypark/internal/domain/entities"
	mocks "github.com/IgorSteps/easypark/mocks/domain/repositories"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestMiddlewareAuthorise_ValidToken(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	testLogger, _ := test.NewNullLogger()
	mockTokenRepo := &mocks.TokenRepository{}
	middleware := middleware.NewAuthMiddleware(mockTokenRepo, testLogger)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// This assertion will be triggered at the ACT stage.
		claims, ok := r.Context().Value("claims").(*entities.Claims)
		assert.True(t, ok)
		assert.Equal(t, "driver", claims.Role)
	})

	tests := []struct {
		name          string
		tokenStr      string
		setupMocks    func()
		expectedCode  int
		expectedClaim *entities.Claims
	}{
		{
			name:     "Valid token",
			tokenStr: "Bearer token",
			setupMocks: func() {
				mockTokenRepo.EXPECT().ParseToken("token").Return(&entities.Claims{Role: "driver"}, nil).Once()
			},
			expectedCode: http.StatusOK,
		},
		{
			name:     "Invalid token",
			tokenStr: "Bearer invalidToken",
			setupMocks: func() {
				mockTokenRepo.EXPECT().ParseToken("invalidToken").Return(nil, errors.New("boom")).Once()
			},
			expectedCode: http.StatusUnauthorized,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.setupMocks()

			req, _ := http.NewRequest("GET", "/", nil)
			req.Header.Set("Authorization", tc.tokenStr)
			recorder := httptest.NewRecorder()

			// ----
			// ACT
			// ----
			middleware.Authorise(handler).ServeHTTP(recorder, req)

			// ------
			// ASSERT
			// ------
			assert.Equal(t, tc.expectedCode, recorder.Code)
			mockTokenRepo.AssertExpectations(t)
		})
	}
}

func TestAuthMiddleware_RequireRole(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	logger, _ := test.NewNullLogger()

	midleware := middleware.NewAuthMiddleware(nil, logger)
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	tests := []struct {
		name            string
		roleInContext   entities.UserRole
		requiredRole    entities.UserRole
		expectedCode    int
		expectedMessage string
	}{
		{
			name:            "Allow admin to access admin route",
			roleInContext:   entities.RoleAdmin,
			requiredRole:    entities.RoleAdmin,
			expectedCode:    http.StatusOK,
			expectedMessage: "",
		},
		{
			name:            "Forbid driver to access admin route",
			roleInContext:   entities.RoleDriver,
			requiredRole:    entities.RoleAdmin,
			expectedCode:    http.StatusForbidden,
			expectedMessage: "Forbidden",
		},
		{
			name:            "Allow driver to access driver route",
			roleInContext:   entities.RoleDriver,
			requiredRole:    entities.RoleDriver,
			expectedCode:    http.StatusOK,
			expectedMessage: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/", nil)

			// Setup role in the context
			ctx := context.WithValue(req.Context(), "claims", &entities.Claims{Role: string(tt.roleInContext)})
			req = req.WithContext(ctx)

			recorder := httptest.NewRecorder()

			// -----
			// ACT
			// -----
			execute := midleware.RequireRole(tt.requiredRole)
			execute(nextHandler).ServeHTTP(recorder, req)

			// ------
			// ASSERT
			// ------
			assert.Equal(t, tt.expectedCode, recorder.Code)

			// If a message is expected, check it
			if tt.expectedMessage != "" {
				responseBody := recorder.Body.String()
				assert.Contains(t, responseBody, tt.expectedMessage)
			}
		})
	}
}
