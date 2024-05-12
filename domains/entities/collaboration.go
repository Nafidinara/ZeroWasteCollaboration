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
	GetAllByUserId(userId uuid.UUID) ([]Collaboration, error)
	Create(collaboration *Collaboration) (*Collaboration, error)
	Update(collaboration *Collaboration) error
	Delete(collaboration *Collaboration) error
}

type CollaborationUsecase interface {
	GetByID(id uuid.UUID) (Collaboration, error)
	GetAllByUserId(userId uuid.UUID) ([]Collaboration, error)
	Create(request *dto.CollaborationRequest) (*Collaboration, error)
	Update(id uuid.UUID, user_id uuid.UUID, request *dto.CollaborationUpdateStatusRequest) (*Collaboration, error)
	Delete(id uuid.UUID) error
}

func ToResponseCollaboration(collaboration *Collaboration) *dto.Collaboration {
	return &dto.Collaboration{
		ID: collaboration.ID,
		User: dto.User{
			ID:       collaboration.UserId,
			Username: collaboration.User.Username,
			Email:    collaboration.User.Email,
			FullName: collaboration.User.FullName,
			Gender:   collaboration.User.Gender,
			ProfileImage: collaboration.User.ProfileImage,
		},
		Organization: dto.Organization{
			Name:    collaboration.Organization.Name,
			Email:   collaboration.Organization.Email,
			Website: collaboration.Organization.Website,
			Phone:   collaboration.Organization.Phone,
			Description: collaboration.Organization.Description,
			Type:    collaboration.Organization.Type,
			ProfileImage: collaboration.Organization.ProfileImage,
		},
		Proposal: dto.Proposal{
			Subject:    collaboration.Proposal.Subject,
			Content:    collaboration.Proposal.Content,
			Attachment: collaboration.Proposal.Attachment,
		},
		Status: collaboration.Status,
	}
}

func ToResponseCollaborations(collaborations []Collaboration) []dto.Collaboration {
	var responseCollaborations []dto.Collaboration
	for _, collaboration := range collaborations {
		responseCollaborations = append(responseCollaborations, *ToResponseCollaboration(&collaboration))
	}
	return responseCollaborations
}
