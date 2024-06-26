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

func (o *organizationRepository) Create(organization *entities.Organization) (*entities.Organization, error) {
	if err := o.DB.Create(organization).Error; err != nil {
		return nil, err
	}
	return organization, nil
}

func (o *organizationRepository) GetAllByUserId(userId uuid.UUID) ([]entities.Organization, error) {
	var organizations []entities.Organization
	err := o.DB.Where("user_id = ?", userId).Preload("User").
		Find(&organizations).Error
	return organizations, err
}

func (o *organizationRepository) GetByID(id uuid.UUID) (entities.Organization, error) {
	var organization entities.Organization
	err := o.DB.Where("id = ?", id).Preload("User").
		First(&organization).Error

	if err != nil {
		return organization, err
	}

	err = o.DB.Table("addresses").Where("organization_id = ?", id).Find(&organization.Addresses).Error
	return organization, err
}

func (o *organizationRepository) GetAll() ([]entities.Organization, error) {
	var organizations []entities.Organization
	err := o.DB.Preload("User").Find(&organizations).Error
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
