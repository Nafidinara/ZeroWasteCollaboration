package usecases

import (
	"time"

	"github.com/google/uuid"

	"redoocehub/domains/entities"
)

type ProposalUsecase struct {
	proposalRepository entities.ProposalRepository
	contextTimeout     time.Duration
}

func NewProposalUsecase(proposalRepository entities.ProposalRepository, timeout time.Duration) entities.ProposalUsecase {
	return &ProposalUsecase{
		proposalRepository: proposalRepository,
		contextTimeout:     timeout,
	}
}

func (u *ProposalUsecase) GetAll() ([]entities.Proposal, error) {
	return u.proposalRepository.GetAll()
}

func (u *ProposalUsecase) GetByID(id uuid.UUID) (entities.Proposal, error) {
	return u.proposalRepository.GetByID(id)
}

//create 
func (u *ProposalUsecase) Create(proposal *entities.Proposal) (*entities.Proposal, error) {
	return u.proposalRepository.Create(proposal)
}

//update
func (u *ProposalUsecase) Update(proposal *entities.Proposal) error {
	proposal.UpdatedAt = time.Now()
	return u.proposalRepository.Update(proposal)
}

//delete
func (u *ProposalUsecase) Delete(id uuid.UUID) error {
	return u.proposalRepository.Delete(&entities.Proposal{ID: id})
}