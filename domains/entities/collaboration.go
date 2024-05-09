package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"redoocehub/domains/dto"
	"redoocehub/domains/types"
)

type Collaboration struct {
	ID             uuid.UUID
	User           User
	UserId         uuid.UUID
	Organization   Organization
	OrganizationId uuid.UUID
	Proposal       Proposal
	ProposalId     uuid.UUID
	Status         types.StatusCollaboration
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}

type CollaborationRepository interface {
	GetByID(id uuid.UUID) (Collaboration, error)
	CreateCollaboration(collaboration *Collaboration) (*Collaboration, error)
	CreateProposal(proposal *Proposal) (*Proposal, error)
	Update(collaboration *Collaboration) error
	Delete(collaboration *Collaboration) error
}

type CollaborationUsecase interface {
	GetByID(id uuid.UUID) (Collaboration, error)
	Create(request *dto.CollaborationRequest) (*Collaboration, error)
	Update(collaboration *Collaboration) error
	Delete(id uuid.UUID) error
}
