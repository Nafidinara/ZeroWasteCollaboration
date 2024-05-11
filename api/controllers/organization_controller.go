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

type OrganizationController struct {
	OrganizationUsecase entities.OrganizationUsecase
	Env                 *bootstrap.Env
}

func NewOrganizationController(organizationUsecase entities.OrganizationUsecase) *OrganizationController {
	return &OrganizationController{OrganizationUsecase: organizationUsecase}
}

func (oc *OrganizationController) GetAll(c echo.Context) error {
	organizations, err := oc.OrganizationUsecase.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
			StatusCode: "Internal Server Error",
			Message:    err.Error(),
			Data:       nil,
		})
	}

	response := entities.ToGetAllResponseOrganizations(organizations)

	return c.JSON(http.StatusOK, infra.SuccessResponse{
		StatusCode: "OK",
		Message:    "Success retrieved all organizations",
		Data:       response,
	})
}

func (oc *OrganizationController) GetByID(c echo.Context) error {

	idParam := c.Param("id")

	id := uuid.MustParse(idParam)

	organization, err := oc.OrganizationUsecase.GetByID(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
			StatusCode: "Internal Server Error",
			Message:    err.Error(),
			Data:       nil,
		})
	}

	response := entities.ToResponseOrganizationDetail(&organization, &organization.User)

	return c.JSON(http.StatusOK, infra.SuccessResponse{
		StatusCode: "OK",
		Message:    "Success retrieved organization",
		Data:       response,
	})
}

func (oc *OrganizationController) Create(c echo.Context) error {

	var request dto.OrganizationRequest

	formHeader, errFile := c.FormFile("profile_image")

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

	if err != nil {
		return c.JSON(http.StatusBadRequest, infra.ErrorResponse{
			StatusCode: "Bad Request",
			Message:    err.Error(),
			Data:       nil,
		})
	}

	request.ProfileImage = formFile

	if err := validation.ValidateRequest(request); err != nil {
		return c.JSON(http.StatusBadRequest, infra.ErrorResponse{
			StatusCode: "Bad Request",
			Message:    "make sure you follow the input requirements",
			Data:       err,
		})
	}

	userId := c.Get("x-user-id").(string)

	request.UserID = uuid.MustParse(userId)

	uploadUrl, err := usecases.NewMediaUpload().FileUpload(dto.File{File: formFile}, entities.CloudinaryEnvSetting{
		CloudName: oc.Env.CLOUDINARY_CLOUD_NAME,
		ApiKey:    oc.Env.CLOUDINARY_API_KEY,
		ApiSecret: oc.Env.CLOUDINARY_API_SECRET,
		UploadFolder: oc.Env.CLOUDINARY_UPLOAD_FOLDER,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
			StatusCode: "Internal Server Error",
			Message:    err.Error(),
			Data:       nil,
		})
	}

	fmt.Println("uploadUrl: ", uploadUrl)

	request.ProfileImage = uploadUrl

	organization, err := oc.OrganizationUsecase.Create(&request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
			StatusCode: "Internal Server Error",
			Message:    err.Error(),
			Data:       nil,
		})
	}

	userData, err := oc.OrganizationUsecase.GetUser(organization.UserID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
			StatusCode: "Internal Server Error",
			Message:    err.Error(),
			Data:       nil,
		})
	}

	response := entities.ToResponseOrganizationDetail(organization, &userData)

	return c.JSON(http.StatusOK, infra.SuccessResponse{
		StatusCode: "OK",
		Message:    "Success created organization",
		Data:       response,
	})
}

func (oc *OrganizationController) Update(c echo.Context) error {
	idParam := c.Param("id")

	id := uuid.MustParse(idParam)

	orgExist, err := oc.OrganizationUsecase.GetByID(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
			StatusCode: "Internal Server Error",
			Message:    err.Error(),
			Data:       nil,
		})
	}

	if orgExist.ID == uuid.Nil {
		return c.JSON(http.StatusNotFound, infra.ErrorResponse{
			StatusCode: "Not Found",
			Message:    "Organization not found",
			Data:       nil,
		})
	}

	var request dto.OrganizationRequest

	err = c.Bind(&request)

	if err := validation.ValidateRequest(request); err != nil {
		return c.JSON(http.StatusBadRequest, infra.ErrorResponse{
			StatusCode: "Bad Request",
			Message:    "make sure you follow the input requirements",
			Data:       err,
		})
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, infra.ErrorResponse{
			StatusCode: "Bad Request",
			Message:    err.Error(),
			Data:       nil,
		})
	}

	orgExist.Name = request.Name
	orgExist.Description = request.Description
	orgExist.Type = request.Type
	orgExist.Email = request.Email
	orgExist.Website = request.Website
	orgExist.Phone = request.Phone

	err = oc.OrganizationUsecase.Update(&orgExist)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
			StatusCode: "Internal Server Error",
			Message:    err.Error(),
			Data:       nil,
		})
	}

	response := entities.ToResponseOrganization(&orgExist)

	return c.JSON(http.StatusOK, infra.SuccessResponse{
		StatusCode: "OK",
		Message:    "Success updated organization",
		Data:       response,
	})
}

func (oc *OrganizationController) Delete(c echo.Context) error {
	idParam := c.Param("id")

	id := uuid.MustParse(idParam)

	err := oc.OrganizationUsecase.Delete(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
			StatusCode: "Internal Server Error",
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, infra.SuccessResponse{
		StatusCode: "OK",
		Message:    "Success deleted organization",
		Data:       nil,
	})
}
