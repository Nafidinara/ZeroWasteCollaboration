package repositories

import (
	"fmt"

	"gorm.io/gorm"

	"redoocehub/domains/dto"
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
		First(&user).Error

	if err != nil {
		return user, fmt.Errorf("user not found")
	}

	err = u.DB.Table("addresses").
		Where("user_id = ?", id).
		Find(&user.Addresses).Error

	if err != nil {
		return user, fmt.Errorf("address not found")
	}

	return user, err
}

func (u *userRepository) Update(user *entities.User) error {
	return u.DB.Save(user).Error
}

func (u *userRepository) GetDashboardData(id string) (*dto.DashboardData, error) {
	var dashboardData dto.DashboardData

	var err error
	var user entities.User

	err = u.DB.Where("id = ?", id).Preload("Organizations").Find(&user).Error

	if err != nil {
		dashboardData.OrganizationCount = 0
	}

	dashboardData.OrganizationCount = int64(len(user.Organizations))

	err = u.DB.Table("addresses").
		Where("user_id = ?", id).
		Count(&dashboardData.AddressCount).Error

	if err != nil {
		dashboardData.AddressCount = 0
	}

	err = u.DB.Table("collaborations").
		Joins("INNER JOIN organizations ON collaborations.organization_id = organizations.id").
		Joins("INNER JOIN users ON organizations.user_id = users.id").
		Where("users.id = ?", id).
		Select(
			"COUNT(CASE WHEN collaborations.status = 'accepted' THEN 1 END) AS accepted",
			"COUNT(CASE WHEN collaborations.status = 'rejected' THEN 1 END) AS rejected",
			"COUNT(CASE WHEN collaborations.status = 'waiting' THEN 1 END) AS waiting",
			"COUNT(CASE WHEN collaborations.status = 'running' THEN 1 END) AS running",
		).
		Scan(&dashboardData.CollaborationReceive).
		Error

	if err != nil {
		fmt.Println("err in collaboration receive: ", err)
		dashboardData.CollaborationReceive = dto.CollaborationStatusCount{
			Accepted: 0,
			Rejected: 0,
			Waiting:  0,
			Running:  0,
		}
	}

	err = u.DB.Table("collaborations").
		Where("user_id = ?", id).
		Select(
			"COUNT(CASE WHEN collaborations.status = 'accepted' THEN 1 END) AS accepted",
			"COUNT(CASE WHEN collaborations.status = 'rejected' THEN 1 END) AS rejected",
			"COUNT(CASE WHEN collaborations.status = 'waiting' THEN 1 END) AS waiting",
			"COUNT(CASE WHEN collaborations.status = 'running' THEN 1 END) AS running",
		).
		Scan(&dashboardData.CollaborationSend).
		Error

	if err != nil {
		fmt.Println("err in collaboration send: ", err)
		dashboardData.CollaborationSend = dto.CollaborationStatusCount{
			Accepted: 0,
			Rejected: 0,
			Waiting:  0,
			Running:  0,
		}
	}

	fmt.Println("dashboardData: ", dashboardData)

	return &dashboardData, err
}
