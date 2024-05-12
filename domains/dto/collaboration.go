package dto

import (
	"mime/multipart"

	"github.com/google/uuid"

	"redoocehub/domains/types"
)

type Collaboration struct {
	ID           uuid.UUID `json:"id"`
	User         User `json:"user"`
	Organization Organization `json:"organization"`
	Proposal     Proposal `json:"proposal"`
	Status       types.StatusCollaboration `json:"status"`
}

type CollaborationRequest struct {
	UserId         uuid.UUID                 `json:"user_id" form:"user_id" binding:"required"`
	OrganizationId uuid.UUID                 `json:"organization_id" form:"organization_id" binding:"required" validate:"required"`
	ProposalId     uuid.UUID                 `json:"proposal_id" form:"proposal_id" binding:"required"`
	Status         types.StatusCollaboration `json:"status" form:"status" binding:"required" validate:"required,oneof=accepted rejected waiting running"`
	Subject        string                    `json:"subject" form:"subject" binding:"required" validate:"required"`
	Content        string                    `json:"content" form:"content" binding:"required" validate:"required"`
	Attachment     multipart.File            `json:"attachment,omitempty" form:"attachment" binding:"required" validate:"required"`
	// Attachment string `json:"attachment" form:"attachment" binding:"required"`
}

type CollaborationUpdateStatusRequest struct {
	Status types.StatusCollaboration `json:"status" form:"status" binding:"required" validate:"required,oneof=accepted rejected waiting running"`
}
