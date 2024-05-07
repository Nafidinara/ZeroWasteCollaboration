package mysql

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"redoocehub/domains/user/entities"
)

type User struct {
	ID           uuid.UUID       `gorm:"primaryKey;not null" json:"id"`
	Username     string          `gorm:"type:varchar(255);not null" json:"username"`
	ProfileImage string          `gorm:"type:varchar(255);" json:"profile_image"`
	Email        string          `gorm:"type:varchar(255);not null" json:"email"`
	FullName     string          `gorm:"type:varchar(255);not null" json:"fullname"`
	Gender       entities.Gender `gorm:"type:enum('male','female')" json:"gender"`
	Password     string          `gorm:"type:varchar(255);not null" json:"password"`
	RefreshToken string          `json:"refresh_token"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
	DeletedAt    gorm.DeletedAt  `json:"deleted_at"`
}
