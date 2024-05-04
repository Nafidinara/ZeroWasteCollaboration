package repositories

import (
	"gorm.io/gorm"

	"redoocehub/domains/user/entities"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) entities.UserRepository {
	return &userRepository{DB: db}
}

func (u *userRepository) Create(user *entities.User) error {
	return u.DB.Create(user).Error
}

func (u *userRepository) GetByEmail(email string) (entities.User, error) {
	var user entities.User
	err := u.DB.Where("email = ?", email).First(&user).Error
	return user, err
}

func (u *userRepository) GetByID(id string) (entities.User, error) {
	var user entities.User
	err := u.DB.Where("id = ?", id).First(&user).Error
	return user, err
}
