package auth

import (
	"github.com/IgorSteps/easypark/internal/domain/entities"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

const (
	configPath = "../../../" // located in project root.
	configName = "config"
	configType = "yaml"
	configKey  = "auth"
)

type AuthConfig struct {
	SecretKey string
}

func NewJWTTokenServiceFromConfig() (*JWTTokenService, error) {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	var config AuthConfig

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	err := viper.UnmarshalKey(configKey, &config)
	if err != nil {
		return nil, err
	}

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