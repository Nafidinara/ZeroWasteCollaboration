package dto

import (
	"time"

	"github.com/google/uuid"

	"redoocehub/domains/types"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	ProfileImage string    `json:"profile_image"`
	Email        string    `json:"email"`
	FullName     string    `json:"fullname"`
	Gender       types.Gender
	// Organizations []Organization `gorm:"foreignKey:user_id"`
}

type LoginRequest struct {
	Email    string `form:"email" binding:"required,email" validate:"required,email"`
	Password string `form:"password" binding:"required" validate:"required"`
}

type LoginResponse struct {
	User         User   `json:"user"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RegisterRequest struct {
	FullName     string       `form:"fullname" binding:"required" validate:"required"`
	Gender       types.Gender `form:"gender" binding:"required" validate:"required,oneof=male female"`
	Username     string       `form:"username" binding:"required" validate:"required"`
	Email        string       `form:"email" binding:"required,email" validate:"required,email"`
	Password     string       `form:"password" binding:"required" validate:"required"`
	ProfileImage string       `form:"profile_image" binding:"required"`
}

type UpdateUserRequest struct {
	FullName     string       `form:"fullname" binding:"required" validate:"required"`
	Gender       types.Gender `form:"gender" binding:"required" validate:"required,oneof=male female"`
	Username     string       `form:"username" binding:"required" validate:"required"`
	Email        string       `form:"email" binding:"required,email" validate:"required,email"`
	Password     string       `form:"password" binding:"required"`
	ProfileImage string       `form:"profile_image" binding:"required"`
}

type RegisterResponse struct {
	ID           uuid.UUID    `json:"id"`
	Email        string       `json:"email"`
	FullName     string       `json:"fullname"`
	ProfileImage string       `json:"profile_image"`
	Gender       types.Gender `json:"gender"`
	Username     string       `json:"username"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
	AccessToken  string       `json:"accessToken"`
	RefreshToken string       `json:"refreshToken"`
}

type ProfileResponse struct {
	ID            uuid.UUID `json:"id"`
	Username      string    `json:"username"`
	ProfileImage  string    `json:"profile_image"`
	Email         string    `json:"email"`
	FullName      string    `json:"fullname"`
	Gender        types.Gender
	Organizations []Organization `json:"organizations"`
	Addresses     []Address      `json:"addresses"`
}

type RefreshTokenRequest struct {
	RefreshToken string `form:"refreshToken" binding:"required"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
type CollaborationStatusCount struct {
	Accepted int64 `json:"accepted"`
	Rejected int64 `json:"rejected"`
	Waiting  int64 `json:"waiting"`
	Running  int64 `json:"running"`
}

type DashboardData struct {
	OrganizationCount int64 `json:"organization_count"`
	AddressCount      int64 `json:"address_count"`
	CollaborationSend  CollaborationStatusCount `json:"collaboration_send"`
	CollaborationReceive  CollaborationStatusCount `json:"collaboration_receive"`
}