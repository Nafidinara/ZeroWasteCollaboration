package usecases

import (
	"time"

	"redoocehub/domains/organization/entities"
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

func (o *organizationUsecase) GetByID(id string) (entities.Organization, error) {
	return o.organizationRepository.GetByID(id)
}

func (o *organizationUsecase) Create(organization *entities.Organization) error {
	return o.organizationRepository.Create(organization)
}

func (o *organizationUsecase) Update(organization *entities.Organization) error {
	return o.organizationRepository.Update(organization)
}

func (o *organizationUsecase) Delete(organization *entities.Organization) error {
	return o.organizationRepository.Delete(organization)
}
