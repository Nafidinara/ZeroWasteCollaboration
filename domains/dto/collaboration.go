package dto

import (
	"github.com/google/uuid"

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
}

type CollaborationRequest struct {
	UserId       uuid.UUID              `json:"user_id" binding:"required"`
	OrganizationId uuid.UUID              `json:"organization_id" binding:"required"`
	ProposalId     uuid.UUID              `json:"proposal_id" binding:"required"`
	Status         types.StatusCollaboration `json:"status" binding:"required" validate:"required"`
	Subject        string                 `json:"subject" binding:"required" validate:"required"`
	Content        string                 `json:"content" binding:"required" validate:"required"`
	Attachment     string                 `json:"attachment" binding:"required" validate:"required"`
}