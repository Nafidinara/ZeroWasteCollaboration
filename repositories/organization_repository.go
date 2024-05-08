package repositories

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"redoocehub/domains/entities"
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

func (o *organizationRepository) GetByID(id uuid.UUID) (entities.Organization, error) {
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
	organization.UpdatedAt = time.Now()

	return o.DB.Save(organization).Error
}

func (o *organizationRepository) Delete(organization *entities.Organization) error {
	return o.DB.Delete(organization).Error
}

func (o *organizationRepository) GetUser(userId uuid.UUID) (entities.User, error) {
	var user entities.User
	err := o.DB.Where("id = ?", userId).First(&user).Error
	return user, err
}
