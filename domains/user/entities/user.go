package entities

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
)

type User struct {
	ID           uuid.UUID
	Username     string
	ProfileImage string
	Email        string
	FullName     string
	Gender       Gender
	Password     string
	RefreshToken string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}

type UserRepository interface {
	Create(user *User) error
	GetByEmail(email string) (User, error)
	GetByID(id string) (User, error)
}

type UserUsecase interface {
	GetUserByID(c context.Context, id string) (User, error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
	ExtractIDFromToken(requestToken string, secret string) (string, error)
	GetProfileByID(c context.Context, userID string) (*User, error)
	Create(c context.Context, user *User) error
	GetUserByEmail(c context.Context, email string) (User, error)
}
