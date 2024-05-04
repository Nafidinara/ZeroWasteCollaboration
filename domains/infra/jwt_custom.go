package infra

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type JwtCustomClaims struct {
	Name string    `json:"name"`
	ID   uuid.UUID `json:"id"`
	jwt.StandardClaims
}

type JwtCustomRefreshClaims struct {
	ID uuid.UUID `json:"id"`
	jwt.StandardClaims
}
