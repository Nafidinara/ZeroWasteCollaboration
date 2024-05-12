package controllers

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"redoocehub/bootstrap"
	"redoocehub/domains/dto"
	"redoocehub/domains/entities"
	"redoocehub/domains/infra"
	"redoocehub/internal/validation"
	"redoocehub/usecases"
)

type CollaborationController struct {
	CollaborationUsecase entities.CollaborationUsecase
	ProposalUsecase      entities.ProposalUsecase
	CloudinaryUsecase    entities.CloudinaryUsecase
	Env                  *bootstrap.Env
}

func NewCollaborationController(collaborationUsecase entities.CollaborationUsecase, proposalUsecase entities.ProposalUsecase, cloudinaryUsecase entities.CloudinaryUsecase, env *bootstrap.Env) *CollaborationController {
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

	formHeader, errFile := c.FormFile("attachment")

	if errFile != nil {
		return c.JSON(http.StatusBadRequest, infra.ErrorResponse{
			StatusCode: "Bad Request",
			Message:    errFile.Error(),
			Data:       nil,
		})
	}

	formFile, errFile := formHeader.Open()

	if errFile != nil {
		return c.JSON(http.StatusBadRequest, infra.ErrorResponse{
			StatusCode: "Bad Request",
			Message:    errFile.Error(),
			Data:       nil,
		})
	}

	err := c.Bind(&request)

	request.Attachment = formFile

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

	uploadUrl, err := usecases.NewMediaUpload().FileUpload(dto.File{File: formFile}, entities.CloudinaryEnvSetting{
		CloudName: cc.Env.CLOUDINARY_CLOUD_NAME,
		ApiKey:    cc.Env.CLOUDINARY_API_KEY,
		ApiSecret: cc.Env.CLOUDINARY_API_SECRET,
		UploadFolder: cc.Env.CLOUDINARY_UPLOAD_FOLDER,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
			StatusCode: "Internal Server Error",
			Message:    err.Error(),
			Data:       nil,
		})
	}

	fmt.Println("uploadUrl: ", uploadUrl)

	proposalReq := dto.ProposalRequest{
		Subject:    request.Subject,
		Content:    request.Content,
		Attachment: uploadUrl,
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

	newCollaboration, err := cc.CollaborationUsecase.GetByID(collaboration.ID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
			StatusCode: "Internal Server Error",
			Message:    err.Error(),
			Data:       nil,
		})
	}

	response := entities.ToResponseCollaboration(&newCollaboration)

	return c.JSON(http.StatusOK, infra.SuccessResponse{
		StatusCode: "OK",
		Message:    "Success created collaboration",
		Data:       response,
	})
}
