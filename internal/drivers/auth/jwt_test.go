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

func TestJWTTokenService_GenerateToken(t *testing.T) {
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
	expiresAt := time.Now().Add(1 * time.Hour).Unix() // 1 hour

	// ----
	// ACT
	// ----
	tokenString, err := tokenService.GenerateToken(testUser, expiresAt)

	// ------
	// ASSERT
	// ------
	assert.NoError(t, err, "Token generation should not return an error")

	// Parse the token to get the claims.
	token, parseErr := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenService.SecretKey), nil
	})
	assert.NoError(t, parseErr, "Token parsing should not return an error")

	claims, ok := token.Claims.(*jwt.MapClaims)
	assert.True(t, ok, "Token claims should be of type jwt.MapClaims")
	assert.Equal(t, testUser.Username, (*claims)["username"], "Username should match")
	assert.Equal(t, string(testUser.Role), (*claims)["role"], "Role should match")
	assert.Equal(t, float64(expiresAt), (*claims)["exp"], "Expiery should match")
}
