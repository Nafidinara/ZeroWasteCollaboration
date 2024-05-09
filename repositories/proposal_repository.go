package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"redoocehub/domains/entities"
)

type ProposalRepository struct {
	DB *gorm.DB
}

func NewProposalRepository(db *gorm.DB) entities.ProposalRepository {
	return &ProposalRepository{DB: db}
}

// get all
func (o *ProposalRepository) GetAll() ([]entities.Proposal, error) {
	var proposals []entities.Proposal

	err := o.DB.Find(&proposals).Error

	return proposals, err
}

// get by id
func (o *ProposalRepository) GetByID(id uuid.UUID) (entities.Proposal, error) {
	var proposal entities.Proposal

	err := o.DB.Where("id = ?", id).First(&proposal).Error

	return proposal, err
}

// create
func (o *ProposalRepository) Create(proposal *entities.Proposal) (*entities.Proposal, error) {
	if err := o.DB.Create(proposal).Error; err != nil {
		return nil, err
	}
	return proposal, nil
}

// update
func (o *ProposalRepository) Update(proposal *entities.Proposal) error {
	return o.DB.Save(proposal).Error
}

//delete
func (o *ProposalRepository) Delete(proposal *entities.Proposal) error {
	return o.DB.Delete(proposal).Error
}
