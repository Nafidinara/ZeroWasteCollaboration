package repositories

import (
	"gorm.io/gorm"

	"redoocehub/domains/entities"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) entities.UserRepository {
	return &userRepository{DB: db}
}

func (u *userRepository) Create(user *entities.User) (*entities.User, error) {
	if err := u.DB.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepository) GetByEmail(email string) (entities.User, error) {
	var user entities.User
	err := u.DB.Where("email = ?", email).Preload("Organizations").First(&user).Error
	return user, err
}

func (u *userRepository) GetByID(id string) (entities.User, error) {
	var user entities.User
	err := u.DB.Where("users.id = ?", id).Preload("Organizations").
		Joins("INNER JOIN addresses ON users.id = addresses.user_id").
		First(&user).Error
	// fmt.Println(user)
	return user, err
}
