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
	OrganizationUsecase  entities.OrganizationUsecase
	Env                  *bootstrap.Env
}

func NewCollaborationController(
	collaborationUsecase entities.CollaborationUsecase,
	proposalUsecase entities.ProposalUsecase,
	cloudinaryUsecase entities.CloudinaryUsecase,
	organizationUsecase entities.OrganizationUsecase,
	env *bootstrap.Env,
) *CollaborationController {
	return &CollaborationController{
		CollaborationUsecase: collaborationUsecase,
		ProposalUsecase:      proposalUsecase,
		CloudinaryUsecase:    cloudinaryUsecase,
		OrganizationUsecase:  organizationUsecase,
		Env:                  env,
	}
}

// get by id
func (cc *CollaborationController) GetByID(c echo.Context) error {
	collaboration, err := cc.CollaborationUsecase.GetByID(uuid.MustParse(c.Param("id")))

	if err != nil {
		return infra.NewErrorResponse(c, http.StatusNotFound, constant.ErrNotFound, constant.ErrGetCollaboration, err.Error())
	}

	response := entities.ToResponseCollaboration(&collaboration)

	return infra.NewSuccessResponse(c, http.StatusOK, constant.SuccessOk, constant.SuccessGetCollaboration, response)
}

// create
func (cc *CollaborationController) Create(c echo.Context) error {
	formHeader, errFile := c.FormFile("attachment")

	if errFile != nil {
		return infra.NewErrorResponse(c, http.StatusBadRequest, constant.ErrBadRequest, constant.ErrFailedGetFile, errFile.Error())
	}

	formFile, errFile := formHeader.Open()

	if errFile != nil {
		return infra.NewErrorResponse(c, http.StatusBadRequest, constant.ErrBadRequest, constant.ErrFailedOpenFile, errFile.Error())
	}

	var request dto.CollaborationRequest

	if err := c.Bind(&request); err != nil {
		return infra.NewErrorResponse(c, http.StatusBadRequest, constant.ErrBadRequest, constant.ErrBinding, err.Error())
	}

	request.UserId = uuid.MustParse(c.Get("x-user-id").(string))

	request.Attachment = formFile

	if err := validation.ValidateRequest(request); err != nil {
		return infra.NewErrorResponse(c, http.StatusBadRequest, constant.ErrBadRequest, constant.ErrValidation, err)
	}

	organization , err := cc.OrganizationUsecase.GetByID(request.OrganizationId)

	if err != nil || organization.ID == uuid.Nil {
		return infra.NewErrorResponse(c, http.StatusNotFound, constant.ErrNotFound, constant.ErrGetOrganization, err.Error())
	}

	uploadUrl, err := usecases.NewMediaUpload().FileUpload(dto.File{File: formFile}, entities.CloudinaryEnvSetting{
		CloudName:    cc.Env.CLOUDINARY_CLOUD_NAME,
		ApiKey:       cc.Env.CLOUDINARY_API_KEY,
		ApiSecret:    cc.Env.CLOUDINARY_API_SECRET,
		UploadFolder: cc.Env.CLOUDINARY_UPLOAD_FOLDER,
	})

	if err != nil {
		return infra.NewErrorResponse(c, http.StatusInternalServerError, constant.ErrInternalServer, constant.ErrFailedUploadFile, err.Error())
	}

	proposalReq := dto.ProposalRequest{
		Subject:    request.Subject,
		Content:    request.Content,
		Attachment: uploadUrl,
	}

	proposal, err := cc.ProposalUsecase.Create(&proposalReq)

	if err != nil {
		return infra.NewErrorResponse(c, http.StatusInternalServerError, constant.ErrInternalServer, constant.ErrCreateProposal, err.Error())
	}

	request.ProposalId = proposal.ID

	collaboration, err := cc.CollaborationUsecase.Create(&request)

	if err != nil {
		return infra.NewErrorResponse(c, http.StatusInternalServerError, constant.ErrInternalServer, constant.ErrCreateCollaboration, err.Error())
	}

	newCollaboration, err := cc.CollaborationUsecase.GetByID(collaboration.ID)

	if err != nil {
		return infra.NewErrorResponse(c, http.StatusInternalServerError, constant.ErrInternalServer, constant.ErrGetCollaboration, err.Error())
	}

	response := entities.ToResponseCollaboration(&newCollaboration)

	emailReq := dto.EmailRequest{
		OrganizationEmail:  response.Organization.Email,
		UserFullName:       response.User.FullName,
		ProposalSubject:    response.Proposal.Subject,
		ProposalContent:    response.Proposal.Content,
		ProposalAttachment: response.Proposal.Attachment,
	}

	emailConfig := entities.EmailConfig{
		SMTPUsername: cc.Env.SMTP_USERNAME,
		SMTPPassword: cc.Env.SMTP_PASSWORD,
		SMTPServer:   cc.Env.SMTP_SERVER,
		SMTPPort:     cc.Env.SMTP_PORT,
	}

	err = email.SendEmail(emailConfig, emailReq)

	if err != nil {
		return infra.NewErrorResponse(c, http.StatusInternalServerError, constant.ErrInternalServer, constant.ErrSendEmail, err.Error())
	}

	return infra.NewSuccessResponse(c, http.StatusCreated, constant.SuccessCreated, constant.SuccessCreateCollaboration, response)
}

// update
func (cc *CollaborationController) Update(c echo.Context) error {
	var request dto.CollaborationUpdateStatusRequest

	if err := c.Bind(&request); err != nil {
		return infra.NewErrorResponse(c, http.StatusBadRequest, constant.ErrBadRequest, constant.ErrBinding, err.Error())
	}

	if err := validation.ValidateRequest(request); err != nil {
		return infra.NewErrorResponse(c, http.StatusBadRequest, constant.ErrBadRequest, constant.ErrValidation, err)
	}

	collaboration, err := cc.CollaborationUsecase.Update(uuid.MustParse(c.Param("id")), uuid.MustParse(c.Get("x-user-id").(string)), &request)

	if err != nil {
		return infra.NewErrorResponse(c, http.StatusInternalServerError, constant.ErrInternalServer, constant.ErrUpdateCollaboration, err.Error())
	}

	response := entities.ToResponseCollaboration(collaboration)

	return infra.NewSuccessResponse(c, http.StatusOK, constant.SuccessOk, constant.SuccessUpdateCollaboration, response)
}

// get all by user
func (cc *CollaborationController) GetAllByUserId(c echo.Context) error {
	collaborations, err := cc.CollaborationUsecase.GetAllByUserId(uuid.MustParse(c.Get("x-user-id").(string)))

	if err != nil {
		return infra.NewErrorResponse(c, http.StatusInternalServerError, constant.ErrInternalServer, constant.ErrGetAllCollaboration, err.Error())
	}

	response := entities.ToResponseCollaborations(collaborations)

	return infra.NewSuccessResponse(c, http.StatusOK, constant.SuccessOk, constant.SuccessGetAllCollaboration, response)
}

// delete
func (cc *CollaborationController) Delete(c echo.Context) error {

	existOrganization, err := cc.CollaborationUsecase.GetByID(uuid.MustParse(c.Param("id")))

	if err != nil {
		return infra.NewErrorResponse(c, http.StatusNotFound, constant.ErrNotFound, constant.ErrGetCollaboration, err.Error())
	}

	if err := cc.CollaborationUsecase.Delete(uuid.MustParse(c.Param("id"))); err != nil {
		return infra.NewErrorResponse(c, http.StatusNotFound, constant.ErrNotFound, constant.ErrDeleteCollaboration, err.Error())
	}

	if err := cc.ProposalUsecase.Delete(existOrganization.ProposalId); err != nil {
		return infra.NewErrorResponse(c, http.StatusNotFound, constant.ErrNotFound, constant.ErrDeleteProposal, err.Error())
	}

	return infra.NewSuccessResponse(c, http.StatusOK, constant.SuccessOk, constant.SuccessDeleteCollaboration, nil)
}
