package controllers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"redoocehub/bootstrap"
	"redoocehub/domains/dto"
	"redoocehub/domains/entities"
	"redoocehub/domains/infra"
	"redoocehub/internal/validation"
)

type CollaborationController struct {
	CollaborationUsecase entities.CollaborationUsecase
	ProposalUsecase      entities.ProposalUsecase
	Env                  *bootstrap.Env
}

func NewCollaborationController(collaborationUsecase entities.CollaborationUsecase, proposalUsecase entities.ProposalUsecase, env *bootstrap.Env) *CollaborationController {
	return &CollaborationController{CollaborationUsecase: collaborationUsecase, ProposalUsecase: proposalUsecase, Env: env}
}

// get by id
func (cc *CollaborationController) GetByID(c echo.Context) error {
	idParam := c.Param("id")

	id := uuid.MustParse(idParam)

	collaboration, err := cc.CollaborationUsecase.GetByID(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
			StatusCode: "Internal Server Error",
			Message:    err.Error(),
			Data:       nil,
		})
	}

	response := entities.ToResponseCollaboration(&collaboration)

	return c.JSON(http.StatusOK, infra.SuccessResponse{
		StatusCode: "OK",
		Message:    "Success retrieved collaboration",
		Data:       response,
	})
}

// create
func (cc *CollaborationController) Create(c echo.Context) error {

	var request dto.CollaborationRequest

	err := c.Bind(&request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, infra.ErrorResponse{
			StatusCode: "Bad Request",
			Message:    err.Error(),
			Data:       nil,
		})
	}

	if err := validation.ValidateRequest(request); err != nil {
		return c.JSON(http.StatusBadRequest, infra.ErrorResponse{
			StatusCode: "Bad Request",
			Message:    "make sure you follow the input requirements",
			Data:       err,
		})
	}

	proposalReq := dto.ProposalRequest{
		Subject:    request.Subject,
		Content:    request.Content,
		Attachment: request.Attachment,
	}

	proposal, err := cc.ProposalUsecase.Create(&proposalReq)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
			StatusCode: "Internal Server Error",
			Message:    err.Error(),
			Data:       nil,
		})
	}

	request.ProposalId = proposal.ID

	userId := c.Get("x-user-id").(string)

	request.UserId = uuid.MustParse(userId)

	collaboration, err := cc.CollaborationUsecase.Create(&request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
			StatusCode: "Internal Server Error",
			Message:    err.Error(),
			Data:       nil,
		})
	}

	response := entities.ToResponseCollaboration(collaboration)

	return c.JSON(http.StatusOK, infra.SuccessResponse{
		StatusCode: "OK",
		Message:    "Success created collaboration",
		Data:       response,
	})
}
