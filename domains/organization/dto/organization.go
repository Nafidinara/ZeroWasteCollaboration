package dto

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"redoocehub/domains/organization/entities"
)

type OrganizationRequest struct {
	Name         string        `form:"name" binding:"required"`
	Description  string        `form:"description" binding:"required"`
	Type         entities.Type `form:"type" binding:"required"`
	ProfileImage string        `form:"profile_image" binding:"required"`
	FoundingDate time.Time     `form:"founding_date" binding:"required"`
	Email        string        `form:"email" binding:"required"`
	Website      string        `form:"website" binding:"required"`
	Phone        string        `form:"phone" binding:"required"`
}

type OrganizationResponse struct {
	ID uuid.UUID `json:"id"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Type         entities.Type `json:"type"`
	ProfileImage string        `json:"profile_image"`
	FoundingDate time.Time     `json:"founding_date"`
	Email        string        `json:"email"`
	Website      string        `json:"website"`
	Phone        string        `json:"phone"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
	DeletedAt    gorm.DeletedAt  `json:"deleted_at"`
}
