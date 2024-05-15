package entities

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"redoocehub/domains/dto"
	"redoocehub/domains/types"
)

type User struct {
	ID            uuid.UUID
	Username      string
	ProfileImage  string
	Email         string
	FullName      string
	Gender        types.Gender
	Password      string
	RefreshToken  string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
	Organizations []Organization `gorm:"foreignKey:user_id"`
	Addresses     []Address      `gorm:"foreignKey:user_id"`
}

type UserRepository interface {
	Create(user *User) (*User, error)
	GetByEmail(email string) (User, error)
	GetByID(id string) (User, error)
	Update(user *User) error
	GetDashboardData(id string) (*dto.DashboardData, error)
}

type UserUsecase interface {
	GetUserByID(c context.Context, id string) (User, error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
	ExtractIDFromToken(requestToken string, secret string) (string, error)
	GetProfileByID(c context.Context, userID string) (*User, error)
	Create(c context.Context, request *dto.RegisterRequest) (*User, error)
	GetUserByEmail(c context.Context, email string) (User, error)
	Update(id uuid.UUID, request *dto.UpdateUserRequest) (*User, error)
	GetDashboardData(id string) (*dto.DashboardData, error)
}

func EntityToDtoUser(user *User) dto.User {
	return dto.User{
		ID:           user.ID,
		Username:     user.Username,
		ProfileImage: user.ProfileImage,
		Email:        user.Email,
		FullName:     user.FullName,
		Gender:       user.Gender,
	}
}

func ToRegisterResponseUser(user *User, accessToken string) *dto.RegisterResponse {
	return &dto.RegisterResponse{
		ID:           user.ID,
		Email:        user.Email,
		Username:     user.Username,
		FullName:     user.FullName,
		Gender:       user.Gender,
		ProfileImage: user.ProfileImage,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		AccessToken:  accessToken,
		RefreshToken: user.RefreshToken,
	}
}

func ToLoginResponseUser(user dto.User, accessToken string, refreshToken string) *dto.LoginResponse {
	return &dto.LoginResponse{
		User: dto.User{
			ID:           user.ID,
			Username:     user.Username,
			ProfileImage: user.ProfileImage,
			Email:        user.Email,
			FullName:     user.FullName,
			Gender:       user.Gender,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

func ToRefreshTokenResponseUser(accessToken string, refreshToken string) *dto.RefreshTokenResponse {
	return &dto.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

func ToProfileResponseUser(user *User) *dto.ProfileResponse {

	var orgs []dto.Organization

	for _, org := range user.Organizations {
		orgs = append(orgs, dto.Organization{
			Name:         org.Name,
			Description:  org.Description,
			Type:         org.Type,
			ProfileImage: org.ProfileImage,
			FoundingDate: org.FoundingDate,
			Email:        org.Email,
			Website:      org.Website,
			Phone:        org.Phone,
		})
	}

	var addresses []dto.Address

	for _, address := range user.Addresses {
		addresses = append(addresses, dto.Address{
			Country:    address.Country,
			City:       address.City,
			Street:     address.Street,
			PostalCode: address.PostalCode,
			State:      address.State,
		})
	}

	return &dto.ProfileResponse{
		ID:            user.ID,
		Username:      user.Username,
		ProfileImage:  user.ProfileImage,
		Email:         user.Email,
		FullName:      user.FullName,
		Gender:        user.Gender,
		Organizations: orgs,
		Addresses:     addresses,
	}
}
