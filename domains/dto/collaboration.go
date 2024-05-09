package dto

import (
	"github.com/google/uuid"

	"redoocehub/domains/types"
)

type Collaboration struct {
	ID             uuid.UUID
	User           User
	Organization   Organization
	Proposal       Proposal
	Status         types.StatusCollaboration
}

type CollaborationRequest struct {
	UserId       uuid.UUID              `json:"user_id" binding:"required"`
	OrganizationId uuid.UUID              `json:"organization_id" binding:"required" validate:"required"`
	ProposalId     uuid.UUID              `json:"proposal_id" binding:"required"`
	Status         types.StatusCollaboration `json:"status" binding:"required" validate:"required,oneof=accepted rejected waiting running"`
	Subject        string                 `json:"subject" binding:"required" validate:"required"`
	Content        string                 `json:"content" binding:"required" validate:"required"`
	Attachment     string                 `json:"attachment" binding:"required" validate:"required"`
}