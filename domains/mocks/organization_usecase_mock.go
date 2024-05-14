package mocks

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"

	"redoocehub/domains/entities"
)

type OrganizationRepositoryMock struct {
	mock.Mock
}

func (o *OrganizationRepositoryMock) Create(organization *entities.Organization) (*entities.Organization, error) {
	args := o.Called(organization)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entities.Organization), args.Error(1)
}

func (o *OrganizationRepositoryMock) GetAllByUserId(userId uuid.UUID) ([]entities.Organization, error) {
	args := o.Called(userId)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]entities.Organization), args.Error(1)
}

func (o *OrganizationRepositoryMock) GetByID(id uuid.UUID) (entities.Organization, error) {
	args := o.Called(id)

	if args.Get(0) == nil {
		return entities.Organization{}, args.Error(1)
	}

	return args.Get(0).(entities.Organization), args.Error(1)
}

func (o *OrganizationRepositoryMock) GetAll() ([]entities.Organization, error) {
	args := o.Called()

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]entities.Organization), args.Error(1)
}

func (o *OrganizationRepositoryMock) Update(organization *entities.Organization) error {
	args := o.Called(organization)

	return args.Error(0)
}

func (o *OrganizationRepositoryMock) Delete(organization *entities.Organization) error {
	args := o.Called(organization)

	return args.Error(0)
}

func (o *OrganizationRepositoryMock) GetUser(userId uuid.UUID) (entities.User, error) {
	args := o.Called(userId)

	if args.Get(0) == nil {
		return entities.User{}, args.Error(1)
	}

	return args.Get(0).(entities.User), args.Error(1)
}
