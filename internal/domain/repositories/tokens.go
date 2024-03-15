package repositories

import "github.com/IgorSteps/easypark/internal/domain/entities"

// TokenRepository provides an interface to validate users' tokens.
type TokenRepository interface {
	// GenereateToken creates an auth token for a given user.
	GenerateToken(user *entities.User) (string, error)

	// ParseToken validates and parses the JWT token, returning our claims.
	ParseToken(tokenStr string) (*entities.Claims, error)
}
