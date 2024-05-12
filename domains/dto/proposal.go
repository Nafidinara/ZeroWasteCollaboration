package dto

type Proposal struct {
	Subject    string `json:"subject"`
	Content    string `json:"content"`
	Attachment string `json:"attachment"`
}

type ProposalRequest struct {
	Subject    string `json:"subject" binding:"required" validate:"required"`
	Content    string `json:"content" binding:"required" validate:"required"`
	Attachment string `json:"attachment" binding:"required" validate:"required"`
}
