package usecases

import (
	"errors"
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

func (c *CollaborationUsecase) Update(id uuid.UUID, userId uuid.UUID, request *dto.CollaborationUpdateStatusRequest) (*entities.Collaboration, error) {
	existingCollaboration, err := c.collaborationRepository.GetByID(id)

	if err != nil {
		return nil, err
	}

	if existingCollaboration.Organization.UserID != userId {
		return nil, errors.New("unauthorized access, you are not the owner of this organization")
	}

	existingCollaboration.Status = request.Status

	return &existingCollaboration, c.collaborationRepository.Update(&existingCollaboration)
}

func (c *CollaborationUsecase) Delete(id uuid.UUID) error {
	existingCollaboration, err := c.collaborationRepository.GetByID(id)

	if err != nil {
		return err
	}

	return c.collaborationRepository.Delete(&existingCollaboration)
}

func (c *CollaborationUsecase) GetAllByUserId(userId uuid.UUID) ([]entities.Collaboration, error) {
	return c.collaborationRepository.GetAllByUserId(userId)
}
