package dto

import (
	"github.com/google/uuid"
)

type Address struct {
	Country        string         `gorm:"type:varchar(255);not null" json:"country"`
	State          string         `gorm:"type:varchar(255);not null" json:"state"`
	City           string         `gorm:"type:varchar(255);not null" json:"city"`
	PostalCode     string         `gorm:"type:varchar(255);not null" json:"postal_code"`
	Street         string         `gorm:"type:varchar(255);not null" json:"street"`
}

type OrganizationAddressRequest struct {
	OrganizationId uuid.UUID `json:"organization_id" binding:"required" validate:"required"`
	Country        string    `json:"country" binding:"required" validate:"required"`
	State          string    `json:"state" binding:"required" validate:"required"`
	City           string    `json:"city" binding:"required" validate:"required"`
	PostalCode     string    `json:"postal_code" binding:"required" validate:"required"`
	Street         string    `json:"street" binding:"required" validate:"required"`
}

type UserAddressRequest struct {
	UserId         uuid.UUID `json:"user_id" binding:"required"`
	Country        string    `json:"country" binding:"required" validate:"required"`
	State          string    `json:"state" binding:"required" validate:"required"`
	City           string    `json:"city" binding:"required" validate:"required"`
	PostalCode     string    `json:"postal_code" binding:"required" validate:"required"`
	Street         string    `json:"street" binding:"required" validate:"required"`
}

type AddressResponse struct {
	ID uuid.UUID `json:"id"`
	// UserId         uuid.UUID `json:"user_id"`
	// OrganizationId uuid.UUID `json:"organization_id"`
	Country      string       `json:"country"`
	State        string       `json:"state"`
	City         string       `json:"city"`
	PostalCode   string       `json:"postal_code"`
	Street       string       `json:"street"`
	User         User         `json:"user"`
	Organization Organization `json:"organization"`
}

