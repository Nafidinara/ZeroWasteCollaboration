package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Type string

const (
	Community    Type = "community"
	Company      Type = "company"
	Institution  Type = "institution"
	NGO          Type = "ngo"
	Agency       Type = "agency"
)

type Organization struct {
	ID           uuid.UUID
	Name         string
	Description  string
	Type         Type
	ProfileImage string
	FoundingDate time.Time
	Email        string
	Website      string
	Phone        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}

type OrganizationRepository interface {
	Create(organization *Organization) error
	GetByID(id string) (Organization, error)
	Update(organization *Organization) error
	Delete(organization *Organization) error
	GetAll() ([]Organization, error)
}

// type UserRepository interface {
// 	Create(user *User) error
// 	GetByEmail(email string) (User, error)
// 	GetByID(id string) (User, error)
// }

// type UserUsecase interface {
// 	GetUserByID(c context.Context, id string) (User, error)
// 	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
// 	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
// 	ExtractIDFromToken(requestToken string, secret string) (string, error)
// 	GetProfileByID(c context.Context, userID string) (*User, error)
// 	Create(c context.Context, user *User) error
// 	GetUserByEmail(c context.Context, email string) (User, error)
// }
