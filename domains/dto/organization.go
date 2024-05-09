package dto

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"redoocehub/domains/types"
)

type Organization struct {
	Name         string                 `json:"name"`
	Description  string                 `json:"description"`
	Type         types.OrganizationType `json:"type"`
	ProfileImage string                 `json:"profile_image"`
	FoundingDate time.Time              `json:"founding_date"`
	Email        string                 `json:"email"`
	Website      string                 `json:"website"`
	Phone        string                 `json:"phone"`
	// User         User                   `json:"owner"`
}

type OrganizationRequest struct {
	Name         string                 `json:"name" binding:"required" validate:"required"`
	UserID       uuid.UUID              `json:"user_id" binding:"required"`
	Description  string                 `json:"description" binding:"required" validate:"required"`
	Type         types.OrganizationType `json:"type" binding:"required" validate:"required"`
	ProfileImage string                 `json:"profile_image" binding:"required" validate:"required"`
	FoundingDate string                 `json:"founding_date" binding:"required" validate:"required,customDateFormat"`
	Email        string                 `json:"email" binding:"required" validate:"required,email"`
	Website      string                 `json:"website" binding:"required" validate:"required,url"`
	Phone        string                 `json:"phone" binding:"required" validate:"required,max=15"`
}

type OrganizationResponse struct {
	ID           uuid.UUID              `json:"id"`
	Name         string                 `json:"name"`
	Description  string                 `json:"description"`
	Type         types.OrganizationType `json:"type"`
	ProfileImage string                 `json:"profile_image"`
	FoundingDate time.Time              `json:"founding_date"`
	Email        string                 `json:"email"`
	Website      string                 `json:"website"`
	Phone        string                 `json:"phone"`
	CreatedAt    time.Time              `json:"created_at"`
	UpdatedAt    time.Time              `json:"updated_at"`
	DeletedAt    gorm.DeletedAt         `json:"deleted_at"`
}

type OrganizationResponseDetail struct {
	ID           uuid.UUID              `json:"id"`
	Name         string                 `json:"name"`
	Description  string                 `json:"description"`
	Type         types.OrganizationType `json:"type"`
	ProfileImage string                 `json:"profile_image"`
	FoundingDate time.Time              `json:"founding_date"`
	Email        string                 `json:"email"`
	Website      string                 `json:"website"`
	Phone        string                 `json:"phone"`
	CreatedAt    time.Time              `json:"created_at"`
	UpdatedAt    time.Time              `json:"updated_at"`
	DeletedAt    gorm.DeletedAt         `json:"deleted_at"`
	User         User                   `json:"owner"`
	Addresses    []Address              `json:"addresses"`
}
