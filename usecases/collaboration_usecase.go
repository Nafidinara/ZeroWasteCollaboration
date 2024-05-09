package usecases

import (
	"time"

	"github.com/google/uuid"

	"redoocehub/domains/dto"
	"redoocehub/domains/entities"
)

type CollaborationUsecase struct {
	collaborationRepository entities.CollaborationRepository
	contextTimeout          time.Duration
}

func NewCollaborationUsecase(collaborationRepository entities.CollaborationRepository, timeout time.Duration) entities.CollaborationUsecase {
	return &CollaborationUsecase{
		collaborationRepository: collaborationRepository,
		contextTimeout:          timeout,
	}
}

func (c *CollaborationUsecase) GetByID(id uuid.UUID) (entities.Collaboration, error) {
	return c.collaborationRepository.GetByID(id)
}

func (c *CollaborationUsecase) Create(request *dto.CollaborationRequest) (*entities.Collaboration, error) {
	collaboration := &entities.Collaboration{
		ID:             uuid.New(),
		UserId:         request.UserId,
		OrganizationId: request.OrganizationId,
		ProposalId:     request.ProposalId,
		Status:         request.Status,
	}

	newCollaboration, err := c.collaborationRepository.Create(collaboration)

	if err != nil {
		return nil, err
	}

	return newCollaboration, nil
}

func (c *CollaborationUsecase) Update(collaboration *entities.Collaboration) error {
	collaboration.UpdatedAt = time.Now()
	return c.collaborationRepository.Update(collaboration)
}

func (c *CollaborationUsecase) Delete(id uuid.UUID) error {
	return c.collaborationRepository.Delete(&entities.Collaboration{ID: id})
}
