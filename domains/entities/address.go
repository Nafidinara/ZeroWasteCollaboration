package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"redoocehub/domains/dto"
)

type Address struct {
	ID             uuid.UUID
	UserID         uuid.UUID `json:"user_id"`
	OrganizationID uuid.UUID `json:"organization_id"`
	Country        string
	State          string
	City           string
	PostalCode     string
	Street         string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}

type AddressRepository interface {
	Create(address *Address) (*Address, error)
	Delete(address *Address) error
	GetAllUserAddress(userId uuid.UUID) ([]Address, error)
	GetAllOrganizationAddress(organizationId uuid.UUID) ([]Address, error)
	GetByID(id uuid.UUID) (Address, error)
}

type AddressUsecase interface {
	CreateUserAddress(request *dto.UserAddressRequest) (*Address, error)
	CreateOrganizationAddress(request *dto.OrganizationAddressRequest) (*Address, error)
	Delete(id uuid.UUID) error
	GetAllUserAddress(userId uuid.UUID) ([]Address, error)
	GetAllOrganizationAddress(organizationId uuid.UUID) ([]Address, error)
}
