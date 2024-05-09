package dto

import (
	"github.com/google/uuid"
)

type Proposal struct {
	ID         uuid.UUID
	Subject    string
	Content    string
	Attachment string
}
