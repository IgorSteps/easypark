package entities

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// Claims is our custom claims carried by JWTs.
type Claims struct {
	UserID   uuid.UUID `json:"userid"`
	Username string    `json:"username"`
	Role     string    `json:"role"`
	jwt.RegisteredClaims
}
