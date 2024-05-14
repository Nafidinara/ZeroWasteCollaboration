package tests

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"redoocehub/domains/dto"
	"redoocehub/domains/entities"
	"redoocehub/domains/mocks"
	"redoocehub/usecases"
)

func TestAddressUsecase(t *testing.T) {

	t.Run("CreateUserAddress - success create user address", func(t *testing.T) {

		addressReq := &dto.UserAddressRequest{
			UserId:     uuid.New(),
			Country:    "Indonesia",
			State:      "Jawa Barat",
			City:       "Bandung",
			Street:     "Jl. Cipaganti",
			PostalCode: "40123",
		}

		newAddress := &entities.Address{
			ID:         uuid.New(),
			UserID:     addressReq.UserId,
			Country:    addressReq.Country,
			State:      addressReq.State,
			City:       addressReq.City,
			Street:     addressReq.Street,
			PostalCode: addressReq.PostalCode,
		}

		mockAddressRepo := new(mocks.AddressRepositoryMock)
		uc := usecases.NewAddressUsecase(mockAddressRepo, 1*time.Second)

		mockAddressRepo.On("Create", mock.Anything).Return(newAddress, nil)

		address, err := uc.CreateUserAddress(addressReq)

		assert.NoError(t, err)
		assert.NotNil(t, address)
		mockAddressRepo.AssertExpectations(t)
	})

	t.Run("CreateOrganizationAddress - success create organization address", func(t *testing.T) {

		addressReq := &dto.OrganizationAddressRequest{
			OrganizationId: uuid.New(),
			Country:        "Indonesia",
			State:          "Jawa Barat",
			City:           "Bandung",
			Street:         "Jl. Cipaganti",
			PostalCode:     "40123",
		}

		newAddress := &entities.Address{
			ID:             uuid.New(),
			OrganizationID: addressReq.OrganizationId,
			Country:        addressReq.Country,
			State:          addressReq.State,
			City:           addressReq.City,
			Street:         addressReq.Street,
			PostalCode:     addressReq.PostalCode,
		}

		mockAddressRepo := new(mocks.AddressRepositoryMock)
		uc := usecases.NewAddressUsecase(mockAddressRepo, 1*time.Second)

		mockAddressRepo.On("Create", mock.Anything).Return(newAddress, nil)

		address, err := uc.CreateOrganizationAddress(addressReq)

		assert.NoError(t, err)
		assert.NotNil(t, address)
		mockAddressRepo.AssertExpectations(t)
	})

	t.Run("Delete - success delete", func(t *testing.T) {

		existAddress := entities.Address{
			ID:             uuid.New(),
			UserID:         uuid.New(),
			OrganizationID: uuid.New(),
			Country:        "Indonesia",
			State:          "Jawa Barat",
			City:           "Bandung",
			Street:         "Jl. Cipaganti",
			PostalCode:     "40123",
		}

		mockAddressRepo := new(mocks.AddressRepositoryMock)
		uc := usecases.NewAddressUsecase(mockAddressRepo, 1*time.Second)

		mockAddressRepo.On("GetByID", mock.Anything).Return(existAddress, nil)
		mockAddressRepo.On("Delete", &existAddress).Return(nil)

		err := uc.Delete(uuid.New())

		assert.NoError(t, err)
	})

	t.Run("Delete - get by id not found", func(t *testing.T) {
		mockAddressRepo := new(mocks.AddressRepositoryMock)
		uc := usecases.NewAddressUsecase(mockAddressRepo, 1*time.Second)

		mockAddressRepo.On("GetByID", mock.Anything).Return(entities.Address{}, assert.AnError)

		err := uc.Delete(uuid.New())

		assert.Error(t, err)
	})

	t.Run("Get All User Address - success get all", func(t *testing.T) {

		mockAddressRepo := new(mocks.AddressRepositoryMock)
		uc := usecases.NewAddressUsecase(mockAddressRepo, 1*time.Second)

		mockAddressRepo.On("GetAllUserAddress", mock.Anything).Return([]entities.Address{}, nil)

		addresses, err := uc.GetAllUserAddress(uuid.New())

		assert.NoError(t, err)
		assert.NotNil(t, addresses)
		assert.Equal(t, 0, len(addresses))
	})

	t.Run("Get All Organization Address - success get all", func(t *testing.T) {

		mockAddressRepo := new(mocks.AddressRepositoryMock)
		uc := usecases.NewAddressUsecase(mockAddressRepo, 1*time.Second)

		mockAddressRepo.On("GetAllOrganizationAddress", mock.Anything).Return([]entities.Address{}, nil)

		addresses, err := uc.GetAllOrganizationAddress(uuid.New())

		assert.NoError(t, err)
		assert.NotNil(t, addresses)
		assert.Equal(t, 0, len(addresses))
	})
}
