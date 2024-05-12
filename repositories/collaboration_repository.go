package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"redoocehub/domains/entities"
)

type CollaborationRepository struct {
	DB *gorm.DB
}

func NewCollaborationRepository(db *gorm.DB) entities.CollaborationRepository {
	return &CollaborationRepository{DB: db}
}

func (o *CollaborationRepository) GetByID(id uuid.UUID) (entities.Collaboration, error) {
	var collaboration entities.Collaboration
	err := o.DB.Where("id = ?", id).
		Preload("User").
		Preload("Organization").
		Preload("Proposal").
		First(&collaboration).Error
	return collaboration, err
}

func (o *CollaborationRepository) Create(collaboration *entities.Collaboration) (*entities.Collaboration, error) {
	if err := o.DB.Create(collaboration).Error; err != nil {
		return nil, err
	}
	return collaboration, nil
}

func (o *CollaborationRepository) Update(collaboration *entities.Collaboration) error {
	return o.DB.Save(collaboration).Error
}

func (o *CollaborationRepository) Delete(collaboration *entities.Collaboration) error {
	return o.DB.Delete(collaboration).Error
}

func (o *CollaborationRepository) GetAllByUserId(userId uuid.UUID) ([]entities.Collaboration, error) {
	var collaborations []entities.Collaboration
	err := o.DB.Where("user_id = ?", userId).
		Preload("User").
		Preload("Organization").
		Preload("Proposal").
		Find(&collaborations).Error
	return collaborations, err
}
