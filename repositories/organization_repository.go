package repositories

import (
	"gorm.io/gorm"

	"redoocehub/domains/organization/entities"
)

type organizationRepository struct {
	DB *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) entities.OrganizationRepository {
	return &organizationRepository{DB: db}
}

func (o *organizationRepository) Create(organization *entities.Organization) error {
	return o.DB.Create(organization).Error
}

func (o *organizationRepository) GetByID(id string) (entities.Organization, error) {
	var organization entities.Organization
	err := o.DB.Where("id = ?", id).First(&organization).Error
	return organization, err
}

func (o *organizationRepository) GetAll() ([]entities.Organization, error) {
	var organizations []entities.Organization
	err := o.DB.Find(&organizations).Error
	return organizations, err
}

func (o *organizationRepository) Update(organization *entities.Organization) error {
	return o.DB.Save(organization).Error
}

func (o *organizationRepository) Delete(organization *entities.Organization) error {
	return o.DB.Delete(organization).Error
}
