package mysql

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"redoocehub/domains/user/entities"
)

type Organization struct {
	ID           uuid.UUID      `gorm:"primaryKey;not null" json:"id"`
	Name         string         `gorm:"type:varchar(255);not null" json:"name"`
	Description  string         `gorm:"type:varchar(255);" json:"description"`
	Type         entities.Type  `gorm:"type:enum('community','company','institution','ngo','agency');not null" json:"type"`
	ProfileImage string         `gorm:"type:varchar(255);" json:"profile_image"`
	FoundingDate time.Time      `json:"founding_date"`
	Email        string         `gorm:"type:varchar(255);not null" json:"email"`
	Website      string         `gorm:"type:varchar(255);" json:"website"`
	Phone        string         `gorm:"type:varchar(255);" json:"phone"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}
