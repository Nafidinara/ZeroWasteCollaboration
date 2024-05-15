package dto

type EmailRequest struct {
	OrganizationEmail  string `json:"organization_email" validate:"required,email"`
	UserFullName       string `json:"user_full_name" validate:"required"`
	ProposalSubject    string `json:"proposal_subject" validate:"required"`
	ProposalContent    string `json:"proposal_content" validate:"required"`
	ProposalAttachment string `json:"proposal_attachment" validate:"required"`
}
