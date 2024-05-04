package repositories

import (
	"gorm.io/gorm"

	"redoocehub/domains"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) domains.UserRepository {
	return &userRepository{DB: db}
}

func (u *userRepository) Create(user *domains.User) error {
	return u.DB.Create(user).Error
}

func (u *userRepository) GetByEmail(email string) (domains.User, error) {
	var user domains.User
	err := u.DB.Where("email = ?", email).First(&user).Error
	return user, err
}

func (u *userRepository) GetByID(id string) (domains.User, error) {
	var user domains.User
	err := u.DB.Where("id = ?", id).First(&user).Error
	return user, err
}
