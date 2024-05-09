package mysql

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Proposal struct {
	ID         uuid.UUID      `gorm:"primaryKey;not null" json:"id"`
	Subject    string         `gorm:"type:varchar(255);not null" json:"subject"`
	Content    string         `gorm:"type:text;not null" json:"content"`
	Attachment string         `gorm:"type:text;" json:"attachment"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}
