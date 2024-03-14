package auth

import (
	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/IgorSteps/easypark/internal/drivers/config"
	"github.com/golang-jwt/jwt/v5"
)

// NewJWTTokenServiceFromConfig returns new JWTTokenService from config.
func NewJWTTokenServiceFromConfig(config config.AuthConfig) (*JWTTokenService, error) {
	return newJWTTokenService(config.SecretKey), nil
}

// JWTTokenService provides functionality to generate JWT tokens.
type JWTTokenService struct {
	secretKey string
}

// NewJWTTokenService returns new instance of JWTTokenService.
func newJWTTokenService(secretKey string) *JWTTokenService {
	return &JWTTokenService{secretKey: secretKey}
}

// GenerateToken generates JWT token for a given user for a given time.
func (s *JWTTokenService) GenerateToken(user *entities.User, expiresAt int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"role":     user.Role,
		"exp":      expiresAt,
	})

	return token.SignedString([]byte(s.secretKey))
}
