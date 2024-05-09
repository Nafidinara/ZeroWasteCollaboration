package usecases

import (
	"time"

	"github.com/google/uuid"

	"redoocehub/domains/dto"
	"redoocehub/domains/entities"
)

type organizationUsecase struct {
	organizationRepository entities.OrganizationRepository
	contextTimeout         time.Duration
}

func NewOrganizationUsecase(organizationRepository entities.OrganizationRepository, timeout time.Duration) entities.OrganizationUsecase {
	return &organizationUsecase{
		organizationRepository: organizationRepository,
		contextTimeout:         timeout,
	}
}

func (o *organizationUsecase) GetAll() ([]entities.Organization, error) {
	return o.organizationRepository.GetAll()
}

func (o *organizationUsecase) GetByID(id uuid.UUID) (entities.Organization, error) {
	return o.organizationRepository.GetByID(id)
}

func (o *organizationUsecase) Create(request *dto.OrganizationRequest) (*entities.Organization, error) {

	foundingDate, err := time.Parse("2006-01-02", request.FoundingDate)

	if err != nil {
		return nil, err
	}

	organization := &entities.Organization{
		ID:           uuid.New(),
		UserID:       request.UserID,
		Name:         request.Name,
		Description:  request.Description,
		Type:         request.Type,
		Email:        request.Email,
		ProfileImage: request.ProfileImage,
		FoundingDate: foundingDate,
		Website:      request.Website,
		Phone:        request.Phone,
	}

	NewOrganization, err := o.organizationRepository.Create(organization)

	if err != nil {
		return nil, err
	}

	return NewOrganization, nil
}

func (o *organizationUsecase) Update(organization *entities.Organization) error {

	return o.organizationRepository.Update(organization)
}

func (o *organizationUsecase) Delete(id uuid.UUID) error {
	org, err := o.organizationRepository.GetByID(id)
	if err != nil {
		return err
	}

	return o.organizationRepository.Delete(&org)
}

func (o *organizationUsecase) GetUser(userId uuid.UUID) (entities.User, error) {
	return o.organizationRepository.GetUser(userId)
}
