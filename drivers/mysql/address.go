package mysql

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Address struct {
	ID             uuid.UUID      `gorm:"primaryKey;not null" json:"id"`
	// User           User           `gorm:"foreignKey:UserId;"`
	UserId         *uuid.UUID     `gorm:"type:varchar(191);nullable" json:"user_id"`
	// Organization   Organization   `gorm:"foreignKey:OrganizationId;"`
	OrganizationId *uuid.UUID     `gorm:"type:varchar(191);nullable" json:"organization_id"`
	Country        string         `gorm:"type:varchar(255);not null" json:"country"`
	State          string         `gorm:"type:varchar(255);not null" json:"state"`
	City           string         `gorm:"type:varchar(255);not null" json:"city"`
	PostalCode     string         `gorm:"type:varchar(255);not null" json:"postal_code"`
	Street         string         `gorm:"type:varchar(255);not null" json:"street"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at"`
}
