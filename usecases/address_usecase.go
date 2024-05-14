package usecases

import (
	"time"

	"github.com/google/uuid"

	"redoocehub/domains/dto"
	"redoocehub/domains/entities"
)

type addressUsecase struct {
	addressRepository entities.AddressRepository
	contextTimeout    time.Duration
}

func NewAddressUsecase(addressRepository entities.AddressRepository, timeout time.Duration) entities.AddressUsecase {
	return &addressUsecase{
		addressRepository: addressRepository,
		contextTimeout:    timeout,
	}
}

func (a *addressUsecase) CreateUserAddress(request *dto.UserAddressRequest) (*entities.Address, error) {
	address := &entities.Address{
		ID:         uuid.New(),
		UserID:     request.UserId,
		Country:    request.Country,
		State:      request.State,
		City:       request.City,
		Street:     request.Street,
		PostalCode: request.PostalCode,
	}

	newAddress, err := a.addressRepository.Create(address)

	if err != nil {
		return nil, err
	}

	return newAddress, nil
}

func (a *addressUsecase) CreateOrganizationAddress(request *dto.OrganizationAddressRequest) (*entities.Address, error) {
	address := &entities.Address{
		ID:             uuid.New(),
		OrganizationID: request.OrganizationId,
		Country:        request.Country,
		State:          request.State,
		City:           request.City,
		Street:         request.Street,
		PostalCode:     request.PostalCode,
	}

	newAddress, err := a.addressRepository.Create(address)

	if err != nil {
		return nil, err
	}

	return newAddress, nil
}

// delete
func (a *addressUsecase) Delete(id uuid.UUID) error {
	address, err := a.addressRepository.GetByID(id)

	if err != nil {
		return err
	}

	return a.addressRepository.Delete(&address)
}

// get user
func (a *addressUsecase) GetAllUserAddress(userId uuid.UUID) ([]entities.Address, error) {
	return a.addressRepository.GetAllUserAddress(userId)
}

// get organization
func (a *addressUsecase) GetAllOrganizationAddress(organizationId uuid.UUID) ([]entities.Address, error) {
	return a.addressRepository.GetAllOrganizationAddress(organizationId)
}
