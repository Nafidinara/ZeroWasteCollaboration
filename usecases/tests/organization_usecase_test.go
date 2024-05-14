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
	"redoocehub/domains/types"
	"redoocehub/usecases"
)

func TestOrganizationUsecase(t *testing.T) {
	t.Run("GetAll - success get all", func(t *testing.T) {

		mockOrganizationRepo := new(mocks.OrganizationRepositoryMock)
		uc := usecases.NewOrganizationUsecase(mockOrganizationRepo, 1*time.Second)

		mockOrganizationRepo.On("GetAll").Return([]entities.Organization{}, nil)

		organizations, err := uc.GetAll()

		assert.NoError(t, err)
		assert.NotNil(t, organizations)
		assert.Equal(t, 0, len(organizations))
	})

	t.Run("GetAll - error get all", func(t *testing.T) {

		mockOrganizationRepo := new(mocks.OrganizationRepositoryMock)
		uc := usecases.NewOrganizationUsecase(mockOrganizationRepo, 1*time.Second)

		mockOrganizationRepo.On("GetAll").Return([]entities.Organization{}, assert.AnError)

		_, err := uc.GetAll()

		assert.Error(t, err)
		mockOrganizationRepo.AssertExpectations(t)
	})

	t.Run("GetAllByUserId - success get all by user id", func(t *testing.T) {

		userId := uuid.New()

		mockOrganizationRepo := new(mocks.OrganizationRepositoryMock)
		uc := usecases.NewOrganizationUsecase(mockOrganizationRepo, 1*time.Second)

		mockOrganizationRepo.On("GetAllByUserId", userId).Return([]entities.Organization{}, nil)

		organizations, err := uc.GetAllByUserId(userId)

		assert.NoError(t, err)
		assert.NotNil(t, organizations)
		assert.Equal(t, 0, len(organizations))
	})

	t.Run("GetAllByUserId - user not found", func(t *testing.T) {

		userId := uuid.New()

		mockOrganizationRepo := new(mocks.OrganizationRepositoryMock)
		uc := usecases.NewOrganizationUsecase(mockOrganizationRepo, 1*time.Second)

		mockOrganizationRepo.On("GetAllByUserId", userId).Return([]entities.Organization{}, assert.AnError)

		_, err := uc.GetAllByUserId(userId)

		assert.Error(t, err)
	})

	t.Run("GetByID - success get by id", func(t *testing.T) {

		organizationId := uuid.New()
		existOrganization := entities.Organization{
			ID:           organizationId,
			UserID:       uuid.New(),
			Name:         "test",
			Description:  "test",
			Type:         types.Community,
			ProfileImage: "test.png",
			FoundingDate: time.Now(),
			Email:        "test@gmail.com",
			Website:      "test.com",
			Phone:        "628976585743",
		}

		mockOrganizationRepo := new(mocks.OrganizationRepositoryMock)
		uc := usecases.NewOrganizationUsecase(mockOrganizationRepo, 1*time.Second)

		mockOrganizationRepo.On("GetByID", organizationId).Return(existOrganization, nil)

		organization, err := uc.GetByID(organizationId)

		assert.NoError(t, err)
		assert.NotNil(t, organization)
		assert.Equal(t, organizationId, organization.ID)
		assert.Equal(t, existOrganization.Name, organization.Name)
		assert.Equal(t, existOrganization.Description, organization.Description)
		assert.Equal(t, existOrganization.Type, organization.Type)
		assert.Equal(t, existOrganization.ProfileImage, organization.ProfileImage)
		assert.Equal(t, existOrganization.FoundingDate, organization.FoundingDate)
		assert.Equal(t, existOrganization.Email, organization.Email)
		assert.Equal(t, existOrganization.Website, organization.Website)
		assert.Equal(t, existOrganization.Phone, organization.Phone)
	})

	t.Run("GetByID - organization not found", func(t *testing.T) {

		organizationId := uuid.New()

		mockOrganizationRepo := new(mocks.OrganizationRepositoryMock)
		uc := usecases.NewOrganizationUsecase(mockOrganizationRepo, 1*time.Second)

		mockOrganizationRepo.On("GetByID", organizationId).Return(entities.Organization{}, assert.AnError)

		_, err := uc.GetByID(organizationId)

		assert.Error(t, err)
		mockOrganizationRepo.AssertExpectations(t)
	})

	t.Run("Create - success create", func(t *testing.T) {

		organizationReq := &dto.OrganizationRequest{
			UserID:       uuid.New(),
			Name:         "test",
			Description:  "test",
			Type:         types.Community,
			ProfileImage: "test.png",
			FoundingDate: "2006-01-02",
			Email:        "test@gmail.com",
			Website:      "test.com",
			Phone:        "628976585743",
		}

		foundingDate, _ := time.Parse("2006-01-02", organizationReq.FoundingDate)

		newOrganization := &entities.Organization{
			ID:           uuid.New(),
			UserID:       organizationReq.UserID,
			Name:         organizationReq.Name,
			Description:  organizationReq.Description,
			Type:         organizationReq.Type,
			Email:        organizationReq.Email,
			ProfileImage: organizationReq.ProfileImage.(string),
			FoundingDate: foundingDate,
			Website:      organizationReq.Website,
			Phone:        organizationReq.Phone,
		}

		mockOrganizationRepo := new(mocks.OrganizationRepositoryMock)
		uc := usecases.NewOrganizationUsecase(mockOrganizationRepo, 1*time.Second)

		mockOrganizationRepo.On("Create", mock.Anything).Return(&entities.Organization{
			ID:           newOrganization.ID,
			UserID:       newOrganization.UserID,
			Name:         newOrganization.Name,
			Description:  newOrganization.Description,
			Type:         newOrganization.Type,
			ProfileImage: newOrganization.ProfileImage,
			FoundingDate: newOrganization.FoundingDate,
			Email:        newOrganization.Email,
			Website:      newOrganization.Website,
			Phone:        newOrganization.Phone,
		}, nil)

		organization, err := uc.Create(organizationReq)

		assert.NoError(t, err)
		assert.NotNil(t, organization)
		assert.Equal(t, organizationReq.UserID, organization.UserID)
		assert.Equal(t, organizationReq.Name, organization.Name)
	})

	t.Run("Create - error create", func(t *testing.T) {

		organizationReq := &dto.OrganizationRequest{
			UserID:       uuid.New(),
			Name:         "test",
			Description:  "test",
			Type:         types.Community,
			ProfileImage: "test.png",
			FoundingDate: "2006-01-02",
			Email:        "test@gmail.com",
			Website:      "test.com",
			Phone:        "628976585743",
		}

		mockOrganizationRepo := new(mocks.OrganizationRepositoryMock)
		uc := usecases.NewOrganizationUsecase(mockOrganizationRepo, 1*time.Second)

		mockOrganizationRepo.On("Create", mock.Anything).Return(&entities.Organization{}, assert.AnError)

		_, err := uc.Create(organizationReq)

		assert.Error(t, err)
		mockOrganizationRepo.AssertExpectations(t)
	})

	t.Run("Update - success update", func(t *testing.T) {

		organization := &entities.Organization{
			ID:           uuid.New(),
			UserID:       uuid.New(),
			Name:         "test",
			Description:  "test",
			Type:         types.Community,
			ProfileImage: "test.png",
			FoundingDate: time.Now(),
			Email:        "test@gmail.com",
			Website:      "test.com",
			Phone:        "628976585743",
		}

		mockOrganizationRepo := new(mocks.OrganizationRepositoryMock)
		uc := usecases.NewOrganizationUsecase(mockOrganizationRepo, 1*time.Second)

		mockOrganizationRepo.On("Update", organization).Return(nil)

		err := uc.Update(organization)

		assert.NoError(t, err)
		mockOrganizationRepo.AssertExpectations(t)
	})

	t.Run("Update - error update", func(t *testing.T) {

		organization := &entities.Organization{
			ID:           uuid.New(),
			UserID:       uuid.New(),
			Name:         "test",
			Description:  "test",
			Type:         types.Community,
			ProfileImage: "test.png",
			FoundingDate: time.Now(),
			Email:        "test@gmail.com",
			Website:      "test.com",
			Phone:        "628976585743",
		}

		mockOrganizationRepo := new(mocks.OrganizationRepositoryMock)
		uc := usecases.NewOrganizationUsecase(mockOrganizationRepo, 1*time.Second)

		mockOrganizationRepo.On("Update", organization).Return(assert.AnError)

		err := uc.Update(organization)

		assert.Error(t, err)
		mockOrganizationRepo.AssertExpectations(t)
	})

	t.Run("Delete - success delete", func(t *testing.T) {
		organizationId := uuid.New()
		existOrganization := entities.Organization{
			ID:           organizationId,
			UserID:       uuid.New(),
			Name:         "test",
			Description:  "test",
			Type:         types.Community,
			ProfileImage: "test.png",
			FoundingDate: time.Now(),
			Email:        "test@gmail.com",
			Website:      "test.com",
			Phone:        "628976585743",
		}

		mockOrganizationRepo := new(mocks.OrganizationRepositoryMock)
		uc := usecases.NewOrganizationUsecase(mockOrganizationRepo, 1*time.Second)

		mockOrganizationRepo.On("GetByID", organizationId).Return(existOrganization, nil)
		mockOrganizationRepo.On("Delete", &existOrganization).Return(nil)

		err := uc.Delete(organizationId)

		assert.NoError(t, err)
		assert.Equal(t, organizationId, existOrganization.ID)
		mockOrganizationRepo.AssertExpectations(t)
	})

	t.Run("Delete - error delete", func(t *testing.T) {

		organizationId := uuid.New()
		existOrganization := entities.Organization{
			ID:           organizationId,
			UserID:       uuid.New(),
			Name:         "test",
			Description:  "test",
			Type:         types.Community,
			ProfileImage: "test.png",
			FoundingDate: time.Now(),
			Email:        "test@gmail.com",
			Website:      "test.com",
			Phone:        "628976585743",
		}

		mockOrganizationRepo := new(mocks.OrganizationRepositoryMock)
		uc := usecases.NewOrganizationUsecase(mockOrganizationRepo, 1*time.Second)

		mockOrganizationRepo.On("GetByID", organizationId).Return(existOrganization, nil)
		mockOrganizationRepo.On("Delete", &existOrganization).Return(assert.AnError)

		err := uc.Delete(organizationId)

		assert.Error(t, err)
		mockOrganizationRepo.AssertExpectations(t)
	})

	t.Run("GetUser - success get user", func(t *testing.T) {

		userId := uuid.New()
		existUser := entities.User{
			ID:       userId,
			Email:    "test",
			Password: "test",
		}

		mockOrganizationRepo := new(mocks.OrganizationRepositoryMock)
		uc := usecases.NewOrganizationUsecase(mockOrganizationRepo, 1*time.Second)

		mockOrganizationRepo.On("GetUser", userId).Return(existUser, nil)

		user, err := uc.GetUser(userId)

		assert.NoError(t, err)
		assert.Equal(t, userId, user.ID)
		mockOrganizationRepo.AssertExpectations(t)
	})

	t.Run("GetUser - error get user", func(t *testing.T) {

		userId := uuid.New()

		mockOrganizationRepo := new(mocks.OrganizationRepositoryMock)
		uc := usecases.NewOrganizationUsecase(mockOrganizationRepo, 1*time.Second)

		mockOrganizationRepo.On("GetUser", userId).Return(entities.User{}, assert.AnError)

		user, err := uc.GetUser(userId)

		assert.Error(t, err)
		assert.Equal(t, user, entities.User{})
		mockOrganizationRepo.AssertExpectations(t)
	})
}
