package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Type string

const (
	Community    Type = "community"
	Company      Type = "company"
	Institution  Type = "institution"
	NGO          Type = "ngo"
	Agency       Type = "agency"
)

type Organization struct {
	ID           uuid.UUID
	Name         string
	Description  string
	Type         Type
	ProfileImage string
	FoundingDate time.Time
	Email        string
	Website      string
	Phone        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}

type OrganizationRepository interface {
	GetAll() ([]Organization, error)
	GetByID(id string) (Organization, error)
	Create(organization *Organization) error
	Update(organization *Organization) error
	Delete(organization *Organization) error
}

type OrganizationUsecase interface {
	GetAll() ([]Organization, error)
	GetByID(id uuid.UUID) (Organization, error)
	Create(organization *Organization) error
	Update(organization *Organization) error
	Delete(id uuid.UUID) error
}
