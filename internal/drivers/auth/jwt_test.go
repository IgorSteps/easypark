package auth_test

import (
	"testing"
	"time"

	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/drivers/auth"
	"github.com/IgorSteps/easypark/internal/drivers/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestJWTTokenService(t *testing.T) {
	// --------
	// ASSEMBLE
	// --------
	config := config.AuthConfig{
		SecretKey: "supersecret",
	}
	tokenService, err := auth.NewJWTTokenServiceFromConfig(config)
	assert.NoError(t, err, "Instantiating new token service shouldn't return an error")
	testUser := &entities.User{
		Username: "flob",
		Role:     "blob",
	}

	// ----
	// ACT
	// ----
	tokenString, err := tokenService.GenerateToken(testUser)
	assert.NoError(t, err, "Token generation should not return an error")

	// Parse the token to get the claims.
	token, parseErr := jwt.ParseWithClaims(tokenString, &entities.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenService.SecretKey), nil
	})
	assert.NoError(t, parseErr, "Token parsing should not return an error")

	// ------
	// ASSERT
	// ------
	claims := token.Claims.(*entities.Claims)
	assert.Equal(t, testUser.Username, claims.Username, "Username should match")
	assert.Equal(t, string(testUser.Role), claims.Role, "Role should match")
	precision := 5 * time.Second
	assert.WithinDuration(t, time.Now().Add(1*time.Hour), claims.ExpiresAt.Time, precision, "Expiration time should be roughly 5 sec from now")
}
