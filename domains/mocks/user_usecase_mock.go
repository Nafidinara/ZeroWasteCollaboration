package mocks

import (
	"github.com/stretchr/testify/mock"

	"redoocehub/domains/dto"
	"redoocehub/domains/entities"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) Create(user *entities.User) (*entities.User, error) {
	args := m.Called(user)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *UserRepositoryMock) GetByEmail(email string) (entities.User, error) {
	args := m.Called(email)

	if args.Get(0) == nil {
		return entities.User{}, args.Error(1)
	}

	return args.Get(0).(entities.User), args.Error(1)
}

func (m *UserRepositoryMock) GetByID(id string) (entities.User, error) {
	args := m.Called(id)

	if args.Get(0) == nil {
		return entities.User{}, args.Error(1)
	}

	return args.Get(0).(entities.User), args.Error(1)
}

func (m *UserRepositoryMock) Update(user *entities.User) error {
	args := m.Called(user)

	return args.Error(0)
}

func (m *UserRepositoryMock) GetDashboardData(id string) (*dto.DashboardData, error) {
	args := m.Called(id)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*dto.DashboardData), args.Error(1)
}
