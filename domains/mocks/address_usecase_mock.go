package mocks

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"

	"redoocehub/domains/entities"
)

type AddressRepositoryMock struct {
	mock.Mock
}

func (a *AddressRepositoryMock) Create(address *entities.Address) (*entities.Address, error) {
	args := a.Called(address)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entities.Address), args.Error(1)
}

func (a *AddressRepositoryMock) GetByID(id uuid.UUID) (entities.Address, error) {
	args := a.Called(id)

	if args.Get(0) == nil {
		return entities.Address{}, args.Error(1)
	}

	return args.Get(0).(entities.Address), args.Error(1)
}

func (a *AddressRepositoryMock) Delete(address *entities.Address) error {
	args := a.Called(address)

	return args.Error(0)
}

func (a *AddressRepositoryMock) GetAllUserAddress(userId uuid.UUID) ([]entities.Address, error) {
	args := a.Called(userId)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]entities.Address), args.Error(1)
}

func (a *AddressRepositoryMock) GetAllOrganizationAddress(organizationId uuid.UUID) ([]entities.Address, error) {
	args := a.Called(organizationId)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]entities.Address), args.Error(1)
}
