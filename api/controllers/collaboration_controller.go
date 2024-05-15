package controllers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"redoocehub/bootstrap"
	"redoocehub/domains/dto"
	"redoocehub/domains/entities"
	"redoocehub/domains/infra"
	"redoocehub/internal/constant"
	"redoocehub/internal/email"
	"redoocehub/internal/validation"
	"redoocehub/usecases"
)

type CollaborationController struct {
	CollaborationUsecase entities.CollaborationUsecase
	ProposalUsecase      entities.ProposalUsecase
	CloudinaryUsecase    entities.CloudinaryUsecase
	Env                  *bootstrap.Env
}

func NewCollaborationController(
	collaborationUsecase entities.CollaborationUsecase,
	proposalUsecase entities.ProposalUsecase,
	cloudinaryUsecase entities.CloudinaryUsecase,
	env *bootstrap.Env,
) *CollaborationController {
	return &CollaborationController{
		CollaborationUsecase: collaborationUsecase,
		ProposalUsecase:      proposalUsecase,
		CloudinaryUsecase:    cloudinaryUsecase,
		Env:                  env,
	}
}

// get by id
func (cc *CollaborationController) GetByID(c echo.Context) error {
	idParam := c.Param("id")

	id := uuid.MustParse(idParam)

	collaboration, err := cc.CollaborationUsecase.GetByID(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
			StatusCode: constant.ErrInternalServer,
			Message:    constant.ErrNotFoundCollaboration,
			Data:       err.Error(),
		})
	}

	response := entities.ToResponseCollaboration(&collaboration)

	return c.JSON(http.StatusOK, infra.SuccessResponse{
		StatusCode: constant.SuccessOk,
		Message:    constant.SuccessGetCollaboration,
		Data:       response,
	})
}

// create
func (cc *CollaborationController) Create(c echo.Context) error {

	var request dto.CollaborationRequest

	formHeader, errFile := c.FormFile("attachment")

	if errFile != nil {
		return c.JSON(http.StatusBadRequest, infra.ErrorResponse{
			StatusCode: constant.ErrBadRequest,
			Message:    constant.ErrFailedGetFile,
			Data:       errFile.Error(),
		})
	}

	formFile, errFile := formHeader.Open()

	if errFile != nil {
		return c.JSON(http.StatusBadRequest, infra.ErrorResponse{
			StatusCode: constant.ErrBadRequest,
			Message:    constant.ErrFailedOpenFile,
			Data:       errFile.Error(),
		})
	}

	err := c.Bind(&request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, infra.ErrorResponse{
			StatusCode: constant.ErrBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	userId := c.Get("x-user-id").(string)

	request.UserId = uuid.MustParse(userId)

	request.Attachment = formFile

	if err := validation.ValidateRequest(request); err != nil {
		return c.JSON(http.StatusBadRequest, infra.ErrorResponse{
			StatusCode: constant.ErrBadRequest,
			Message:    constant.ErrValidation,
			Data:       err,
		})
	}

	uploadUrl, err := usecases.NewMediaUpload().FileUpload(dto.File{File: formFile}, entities.CloudinaryEnvSetting{
		CloudName:    cc.Env.CLOUDINARY_CLOUD_NAME,
		ApiKey:       cc.Env.CLOUDINARY_API_KEY,
		ApiSecret:    cc.Env.CLOUDINARY_API_SECRET,
		UploadFolder: cc.Env.CLOUDINARY_UPLOAD_FOLDER,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
			StatusCode: constant.ErrInternalServer,
			Message:    constant.ErrUploadFile,
			Data:       err.Error(),
		})
	}

	proposalReq := dto.ProposalRequest{
		Subject:    request.Subject,
		Content:    request.Content,
		Attachment: uploadUrl,
	}

	proposal, err := cc.ProposalUsecase.Create(&proposalReq)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
			StatusCode: constant.ErrInternalServer,
			Message:    constant.ErrFailedProposal,
			Data:       err.Error(),
		})
	}

	request.ProposalId = proposal.ID

	collaboration, err := cc.CollaborationUsecase.Create(&request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
			StatusCode: constant.ErrInternalServer,
			Message:    constant.ErrCreateCollaboration,
			Data:       err.Error(),
		})
	}

	newCollaboration, err := cc.CollaborationUsecase.GetByID(collaboration.ID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
			StatusCode: constant.ErrInternalServer,
			Message:    constant.ErrNotFoundCollaboration,
			Data:       err.Error(),
		})
	}

	response := entities.ToResponseCollaboration(&newCollaboration)

	emailReq := dto.EmailRequest{
		OrganizationEmail:  response.Organization.Email,
		UserFullName:       response.User.FullName,
		ProposalSubject:    response.Proposal.Subject,
		ProposalContent:    response.Proposal.Content,
		ProposalAttachment: response.Proposal.Attachment,
	}

	err = email.NewEmailService().SendEmail(emailReq)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
			StatusCode: constant.ErrInternalServer,
			Message:    constant.ErrFailedSendEmail,
			Data:       err.Error(),
		})
	}

	return c.JSON(http.StatusOK, infra.SuccessResponse{
		StatusCode: constant.SuccessOk,
		Message:    constant.SuccessCreateCollaboration,
		Data:       response,
	})
}

// update
func (cc *CollaborationController) Update(c echo.Context) error {
	var request dto.CollaborationUpdateStatusRequest

	err := c.Bind(&request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, infra.ErrorResponse{
			StatusCode: constant.ErrBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	if err := validation.ValidateRequest(request); err != nil {
		return c.JSON(http.StatusBadRequest, infra.ErrorResponse{
			StatusCode: constant.ErrBadRequest,
			Message:    constant.ErrParameterNotFound,
			Data:       err,
		})
	}

	idParam := c.Param("id")
	id := uuid.MustParse(idParam)

	userId := c.Get("x-user-id").(string)

	collaboration, err := cc.CollaborationUsecase.Update(id, uuid.MustParse(userId), &request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
			StatusCode: constant.ErrInternalServer,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	response := entities.ToResponseCollaboration(collaboration)

	return c.JSON(http.StatusOK, infra.SuccessResponse{
		StatusCode: constant.SuccessOk,
		Message:    constant.SuccessUpdateCollaboration,
		Data:       response,
	})
}

// get all by user
func (cc *CollaborationController) GetAllByUserId(c echo.Context) error {
	userId := c.Get("x-user-id").(string)

	collaborations, err := cc.CollaborationUsecase.GetAllByUserId(uuid.MustParse(userId))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
			StatusCode: constant.ErrInternalServer,
			Message:    constant.ErrNotFoundCollaboration,
			Data:       err.Error(),
		})
	}

	response := entities.ToResponseCollaborations(collaborations)

	return c.JSON(http.StatusOK, infra.SuccessResponse{
		StatusCode: constant.SuccessOk,
		Message:    constant.SuccessGetAllCollaboration,
		Data:       response,
	})
}

// delete
func (cc *CollaborationController) Delete(c echo.Context) error {
	idParam := c.Param("id")

	id := uuid.MustParse(idParam)

	err := cc.CollaborationUsecase.Delete(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
			StatusCode: constant.ErrInternalServer,
			Message:    constant.ErrFailedDeleteCollaboration,
			Data:       err.Error(),
		})
	}

	return c.JSON(http.StatusOK, infra.SuccessResponse{
		StatusCode: constant.SuccessOk,
		Message:    constant.SuccessDeleteCollaboration,
		Data:       nil,
	})
}
