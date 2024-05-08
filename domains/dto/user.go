package dto

import (
	"time"

	"github.com/google/uuid"

	"redoocehub/domains/entities"
)

type LoginRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RegisterRequest struct {
	FullName string          `form:"fullname" binding:"required"`
	Gender   entities.Gender `form:"gender" binding:"required"`
	Username string          `form:"username" binding:"required"`
	Email    string          `form:"email" binding:"required,email"`
	Password string          `form:"password" binding:"required"`
}

type RegisterResponse struct {
	ID           uuid.UUID       `json:"id"`
	Email        string          `json:"email"`
	FullName     string          `json:"fullname"`
	ProfileImage string          `json:"profile_image"`
	Gender       entities.Gender `json:"gender"`
	Username     string          `json:"username"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
	AccessToken  string          `json:"accessToken"`
	RefreshToken string          `json:"refreshToken"`
}

type RefreshTokenRequest struct {
	RefreshToken string `form:"refreshToken" binding:"required"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
