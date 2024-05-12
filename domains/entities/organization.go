package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"redoocehub/domains/dto"
	"redoocehub/domains/types"
)

type Organization struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	Name         string
	Description  string
	Type         types.OrganizationType
	ProfileImage string
	FoundingDate time.Time
	Email        string
	Website      string
	Phone        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
	User         User
	Addresses    []Address `gorm:"foreignKey:organization_id"`
}

type OrganizationRepository interface {
	GetAll() ([]Organization, error)
	GetAllByUserId(userId uuid.UUID) ([]Organization, error)
	GetByID(id uuid.UUID) (Organization, error)
	Create(organization *Organization) (*Organization, error)
	Update(organization *Organization) error
	Delete(organization *Organization) error
	GetUser(userId uuid.UUID) (User, error)
}

type OrganizationUsecase interface {
	GetAll() ([]Organization, error)
	GetAllByUserId(userId uuid.UUID) ([]Organization, error)
	GetByID(id uuid.UUID) (Organization, error)
	Create(request *dto.OrganizationRequest) (*Organization, error)
	Update(organization *Organization) error
	Delete(id uuid.UUID) error
	GetUser(userId uuid.UUID) (User, error)
}

func EntityToDtoOrganizationDetail(organization *Organization) dto.OrganizationResponseDetail {
	return dto.OrganizationResponseDetail{
		ID:           organization.ID,
		Name:         organization.Name,
		Description:  organization.Description,
		Type:         organization.Type,
		ProfileImage: organization.ProfileImage,
		FoundingDate: organization.FoundingDate,
		Email:        organization.Email,
		Website:      organization.Website,
		Phone:        organization.Phone,
		CreatedAt:    organization.CreatedAt,
		UpdatedAt:    organization.UpdatedAt,
		User: dto.User{
			ID:           organization.UserID,
			Username:     organization.User.Username,
			ProfileImage: organization.User.ProfileImage,
			Email:        organization.User.Email,
			FullName:     organization.User.FullName,
			Gender:       organization.User.Gender,
		},
	}
}

func EntityToDtoOrganization(organization *Organization) dto.OrganizationResponse {
	return dto.OrganizationResponse{
		ID:           organization.ID,
		Name:         organization.Name,
		Description:  organization.Description,
		Type:         organization.Type,
		ProfileImage: organization.ProfileImage,
		FoundingDate: organization.FoundingDate,
		Email:        organization.Email,
		Website:      organization.Website,
		Phone:        organization.Phone,
		CreatedAt:    organization.CreatedAt,
		UpdatedAt:    organization.UpdatedAt,
	}
}

func ToGetAllResponseOrganizations(organizations []Organization) []dto.OrganizationResponse {
	var response []dto.OrganizationResponse
	for _, organization := range organizations {
		response = append(response, EntityToDtoOrganization(&organization))
	}
	return response
}

func ToResponseOrganization(organization *Organization) dto.OrganizationResponse {
	return dto.OrganizationResponse{
		ID:           organization.ID,
		Name:         organization.Name,
		Description:  organization.Description,
		Type:         organization.Type,
		ProfileImage: organization.ProfileImage,
		FoundingDate: organization.FoundingDate,
		Email:        organization.Email,
		Website:      organization.Website,
		Phone:        organization.Phone,
		CreatedAt:    organization.CreatedAt,
		UpdatedAt:    organization.UpdatedAt,
	}
}

func ToResponseOrganizationDetail(organization *Organization, user *User) dto.OrganizationResponseDetail {

	var addresses []dto.Address

	for _, address := range organization.Addresses {
		addresses = append(addresses, dto.Address{
			Street:     address.Street,
			City:       address.City,
			Country:    address.Country,
			State:      address.State,
			PostalCode: address.PostalCode,
		})
	}

	return dto.OrganizationResponseDetail{
		ID:           organization.ID,
		Name:         organization.Name,
		Description:  organization.Description,
		Type:         organization.Type,
		ProfileImage: organization.ProfileImage,
		FoundingDate: organization.FoundingDate,
		Email:        organization.Email,
		Website:      organization.Website,
		Phone:        organization.Phone,
		CreatedAt:    organization.CreatedAt,
		UpdatedAt:    organization.UpdatedAt,
		User: dto.User{
			ID:           user.ID,
			Username:     user.Username,
			ProfileImage: user.ProfileImage,
			Email:        user.Email,
			FullName:     user.FullName,
			Gender:       user.Gender,
		},
		Addresses: addresses,
	}
}
