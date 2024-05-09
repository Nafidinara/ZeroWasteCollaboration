package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Proposal struct {
	ID         uuid.UUID
	Subject    string
	Content    string
	Attachment string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}

type ProposalRepository interface {
	GetAll() ([]Proposal, error)
	GetByID(id uuid.UUID) (Proposal, error)
	Create(proposal *Proposal) (*Proposal, error)
	Update(proposal *Proposal) error
	Delete(proposal *Proposal) error
}

type ProposalUsecase interface {
	GetAll() ([]Proposal, error)
	GetByID(id uuid.UUID) (Proposal, error)
	Create(proposal *Proposal) (*Proposal, error)
	Update(proposal *Proposal) error
	Delete(id uuid.UUID) error
}
