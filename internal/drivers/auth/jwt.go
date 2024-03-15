package auth

import (
	"errors"
	"time"

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
	SecretKey string
}

// NewJWTTokenService returns new instance of JWTTokenService.
func newJWTTokenService(secretKey string) *JWTTokenService {
	return &JWTTokenService{SecretKey: secretKey}
}

// GenerateToken generates JWT token for a given user for a given time.
func (s *JWTTokenService) GenerateToken(user *entities.User) (string, error) {
	claims := entities.Claims{
		UserID:   user.ID,
		Username: user.Username,
		Role:     string(user.Role),
		RegisteredClaims: jwt.RegisteredClaims{
			// TODO: Move timeout to config
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(s.SecretKey))
}

// ParseToken validates and parses the JWT token, returning our claims.
func (s *JWTTokenService) ParseToken(tokenStr string) (*entities.Claims, error) {
	claims := &entities.Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(s.SecretKey), nil
		},
	)
	if err != nil {
		// TODO: Could these errors be wrapped in our domain error type?
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
