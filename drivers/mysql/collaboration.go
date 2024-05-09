package mysql

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"redoocehub/domains/types"
)

type Collaboration struct {
	ID             uuid.UUID                 `gorm:"primaryKey;not null" json:"id"`
	User           User                      `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	UserId         uuid.UUID                 `gorm:"type:varchar(191);index" json:"user_id"`
	Organization   Organization              `gorm:"foreignKey:OrganizationId;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	OrganizationId uuid.UUID                 `gorm:"type:varchar(191);index" json:"organization_id"`
	Proposal       Proposal                  `gorm:"foreignKey:ProposalId;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	ProposalId     uuid.UUID                 `gorm:"type:varchar(191);index" json:"proposal_id"`
	Status         types.StatusCollaboration `gorm:"type:enum('accepted','rejected','waiting','running');" json:"status"`
	CreatedAt      time.Time                 `json:"created_at"`
	UpdatedAt      time.Time                 `json:"updated_at"`
	DeletedAt      gorm.DeletedAt            `json:"deleted_at"`
}
