package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"redoocehub/domains/entities"
)

type addressRepository struct {
	DB *gorm.DB
}

func NewAddressRepository(db *gorm.DB) entities.AddressRepository {
	return &addressRepository{DB: db}
}

//get All
func (a *addressRepository) GetAll() ([]entities.Address, error) {
	var addresses []entities.Address
	err := a.DB.Preload("User").Preload("Organization").Find(&addresses).Error
	return addresses, err
}

//get by id
func (a *addressRepository) GetByID(id uuid.UUID) (entities.Address, error) {
	var address entities.Address
	err := a.DB.Where("id = ?", id).Preload("User").Preload("Organization").First(&address).Error
	return address, err
}

//create
func (a *addressRepository) Create(address *entities.Address) (*entities.Address, error) {
	if err := a.DB.Create(address).Error; err != nil {
		return nil, err
	}
	return address, nil
}

//update
func (a *addressRepository) Update(address *entities.Address) error {
	return a.DB.Save(address).Error
}

//delete
func (a *addressRepository) Delete(address *entities.Address) error {
	return a.DB.Delete(address).Error
}

//get user address
func (a *addressRepository) GetAllUserAddress(userId uuid.UUID) ([]entities.Address, error) {
	var address []entities.Address
	err := a.DB.Where("user_id = ?", userId).Find(&address).Error
	return address, err
}

//get organization address
func (a *addressRepository) GetAllOrganizationAddress(organizationId uuid.UUID) ([]entities.Address, error) {
	var address []entities.Address
	err := a.DB.Where("organization_id = ?", organizationId).Find(&address).Error
	return address, err
}