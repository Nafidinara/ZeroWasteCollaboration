package infra

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"

	"redoocehub/domains/types"
)

type JwtCustomClaims struct {
	FullName string       `json:"fullname"`
	Email    string       `json:"email"`
	Username string       `json:"username"`
	Gender   types.Gender `json:"gender"`
	ID       uuid.UUID    `json:"id"`
	jwt.StandardClaims
}

type JwtCustomRefreshClaims struct {
	ID uuid.UUID `json:"id"`
	jwt.StandardClaims
}
