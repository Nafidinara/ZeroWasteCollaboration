package dto

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"redoocehub/domains/entities"
)

type OrganizationRequest struct {
	Name         string        `json:"name" binding:"required" validate:"required"`
	Description  string        `json:"description" binding:"required" validate:"required"`
	Type         entities.Type `json:"type" binding:"required" validate:"required"`
	ProfileImage string        `json:"profile_image" binding:"required" validate:"required"`
	FoundingDate string        `json:"founding_date" binding:"required" validate:"required,customDateFormat"`
	Email        string        `json:"email" binding:"required" validate:"required,email"`
	Website      string        `json:"website" binding:"required" validate:"required,url"`
	Phone        string        `json:"phone" binding:"required" validate:"required,max=15"`
}

type OrganizationResponse struct {
	ID           uuid.UUID      `json:"id"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	Type         entities.Type  `json:"type"`
	ProfileImage string         `json:"profile_image"`
	FoundingDate time.Time      `json:"founding_date"`
	Email        string         `json:"email"`
	Website      string         `json:"website"`
	Phone        string         `json:"phone"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
	User         entities.User  `json:"user"`
}
