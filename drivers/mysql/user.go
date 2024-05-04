package mysql

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID      `gorm:"primaryKey;not null" json:"id"`
	Email        string         `gorm:"type:varchar(255);not null" json:"email"`
	Name         string         `gorm:"type:varchar(255);not null" json:"name"`
	Password     string         `gorm:"type:varchar(255);not null" json:"password"`
	RefreshToken string         `json:"refresh_token"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}
