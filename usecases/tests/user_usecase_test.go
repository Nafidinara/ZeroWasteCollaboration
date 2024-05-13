package tests

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"redoocehub/domains/dto"
	"redoocehub/domains/entities"
	"redoocehub/domains/mocks"
	"redoocehub/domains/types"
	"redoocehub/usecases"
)

func TestGetUserByID(t *testing.T) {
	userId := uuid.New()

	existUser := entities.User{
		ID:           userId,
		Email:        "user@example.com",
		Username:     "testuser",
		Password:     "testpassword",
		FullName:     "Test User",
		Gender:       types.Male,
		ProfileImage: "image.png",
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepositoryMock)
		uc := usecases.NewUserUsecase(mockUserRepo, 1*time.Second)

		mockUserRepo.On("GetByID", userId.String()).Return(existUser, nil)

		user, err := uc.GetUserByID(context.Background(), userId.String())

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, existUser.ID, user.ID)
		assert.Equal(t, existUser.Email, user.Email)
		assert.Equal(t, existUser.Username, user.Username)
		assert.Equal(t, existUser.FullName, user.FullName)
		assert.Equal(t, existUser.Gender, user.Gender)
		assert.Equal(t, existUser.ProfileImage, user.ProfileImage)
	})

	t.Run("error", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepositoryMock)
		uc := usecases.NewUserUsecase(mockUserRepo, 1*time.Second)

		mockUserRepo.On("GetByID", userId.String()).Return(entities.User{}, assert.AnError)

		_, err := uc.GetUserByID(context.Background(), userId.String())

		assert.Error(t, err)
		assert.Equal(t, assert.AnError, err)
		mockUserRepo.AssertExpectations(t)
	})
}

func TestCreateUser(t *testing.T) {
	request := &dto.RegisterRequest{
		Email:        "user@example.com",
		Username:     "testuser",
		FullName:     "Test User",
		Gender:       types.Male,
		Password:     "testpassword",
		ProfileImage: "image.png",
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepositoryMock)
		uc := usecases.NewUserUsecase(mockUserRepo, 1*time.Second)

		mockUserRepo.On("Create", mock.Anything).Return(&entities.User{
			ID:           uuid.New(),
			Email:        request.Email,
			Username:     request.Username,
			FullName:     request.FullName,
			Gender:       request.Gender,
			Password:     request.Password,
			ProfileImage: request.ProfileImage,
		}, nil)

		user, err := uc.Create(context.Background(), request)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, request.Email, user.Email)
		assert.Equal(t, request.Username, user.Username)
		assert.Equal(t, request.FullName, user.FullName)
		assert.Equal(t, request.Gender, user.Gender)
	})

	t.Run("error", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepositoryMock)
		uc := usecases.NewUserUsecase(mockUserRepo, 1*time.Second)

		mockUserRepo.On("Create", mock.Anything).Return(&entities.User{}, assert.AnError)

		_, err := uc.Create(context.Background(), request)

		assert.Error(t, err)
		assert.Equal(t, assert.AnError, err)
		mockUserRepo.AssertExpectations(t)
	})
}

func TestGetUserByEmail(t *testing.T) {
	userId := uuid.New()

	existUser := entities.User{
		ID:           userId,
		Email:        "user@example.com",
		Username:     "testuser",
		Password:     "testpassword",
		FullName:     "Test User",
		Gender:       types.Male,
		ProfileImage: "image.png",
	}

	t.Run("success", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepositoryMock)
		uc := usecases.NewUserUsecase(mockUserRepo, 1*time.Second)

		mockUserRepo.On("GetByEmail", existUser.Email).Return(existUser, nil)

		user, err := uc.GetUserByEmail(context.Background(), existUser.Email)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, existUser.ID, user.ID)
		assert.Equal(t, existUser.Email, user.Email)
		assert.Equal(t, existUser.Username, user.Username)
		assert.Equal(t, existUser.FullName, user.FullName)
		assert.Equal(t, existUser.Gender, user.Gender)
		assert.Equal(t, existUser.ProfileImage, user.ProfileImage)
	})

	t.Run("error", func(t *testing.T) {
		mockUserRepo := new(mocks.UserRepositoryMock)
		uc := usecases.NewUserUsecase(mockUserRepo, 1*time.Second)

		mockUserRepo.On("GetByEmail", existUser.Email).Return(entities.User{}, assert.AnError)

		_, err := uc.GetUserByEmail(context.Background(), existUser.Email)

		assert.Error(t, err)
		assert.Equal(t, assert.AnError, err)
		mockUserRepo.AssertExpectations(t)
	})
}
