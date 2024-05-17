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
		return infra.NewErrorResponse(c, http.StatusInternalServerError, constant.ErrInternalServer, constant.ErrGetAllOrganization, err.Error())
	}

	response := entities.ToGetAllResponseOrganizations(organizations)

	return infra.NewSuccessResponse(c, http.StatusOK, constant.SuccessOk, constant.SuccessGetAllOrganization, response)
}

func (oc *OrganizationController) GetByID(c echo.Context) error {
	organization, err := oc.OrganizationUsecase.GetByID(uuid.MustParse(c.Param("id")))

	if err != nil {
		return infra.NewErrorResponse(c, http.StatusNotFound, constant.ErrNotFound, constant.ErrGetAllOrganization, err.Error())
	}

	response := entities.ToResponseOrganizationDetail(&organization, &organization.User)

	return infra.NewSuccessResponse(c, http.StatusOK, constant.SuccessOk, constant.SuccessGetOrganization, response)
}

func (oc *OrganizationController) Create(c echo.Context) error {
	var request dto.OrganizationRequest

	formHeader, errFile := c.FormFile("profile_image")

	if errFile != nil {
		return infra.NewErrorResponse(c, http.StatusBadRequest, constant.ErrBadRequest, constant.ErrFailedGetFile, errFile.Error())
	}

	formFile, errFile := formHeader.Open()

	if errFile != nil {
		return infra.NewErrorResponse(c, http.StatusBadRequest, constant.ErrBadRequest, constant.ErrFailedOpenFile, errFile.Error())
	}

	if err := c.Bind(&request); err != nil {
		return infra.NewErrorResponse(c, http.StatusBadRequest, constant.ErrBadRequest, constant.ErrBinding, err.Error())
	}

	request.ProfileImage = formFile

	if err := validation.ValidateRequest(request); err != nil {
		return infra.NewErrorResponse(c, http.StatusBadRequest, constant.ErrValidation, constant.ErrBinding, err)
	}

	request.UserID = uuid.MustParse(c.Get("x-user-id").(string))

	uploadUrl, err := usecases.NewMediaUpload().FileUpload(dto.File{File: formFile}, entities.CloudinaryEnvSetting{
		CloudName:    oc.Env.CLOUDINARY_CLOUD_NAME,
		ApiKey:       oc.Env.CLOUDINARY_API_KEY,
		ApiSecret:    oc.Env.CLOUDINARY_API_SECRET,
		UploadFolder: oc.Env.CLOUDINARY_UPLOAD_FOLDER,
	})

	if err != nil {
		return infra.NewErrorResponse(c, http.StatusInternalServerError, constant.ErrInternalServer, constant.ErrFailedUploadFile, err.Error())
	}

	request.ProfileImage = uploadUrl

	organization, err := oc.OrganizationUsecase.Create(&request)

	if err != nil {
		return infra.NewErrorResponse(c, http.StatusInternalServerError, constant.ErrInternalServer, constant.ErrCreateOrganization, err.Error())
	}

	userData, err := oc.OrganizationUsecase.GetUser(organization.UserID)

	if err != nil {
		return infra.NewErrorResponse(c, http.StatusInternalServerError, constant.ErrInternalServer, constant.ErrGetUser, err.Error())
	}

	response := entities.ToResponseOrganizationDetail(organization, &userData)

	return infra.NewSuccessResponse(c, http.StatusCreated, constant.SuccessCreated, constant.SuccessCreateOrganization, response)
}

func (oc *OrganizationController) Update(c echo.Context) error {
	orgExist, err := oc.OrganizationUsecase.GetByID(uuid.MustParse(c.Param("id")))

	if err != nil {
		return infra.NewErrorResponse(c, http.StatusNotFound, constant.ErrNotFound, constant.ErrGetOrganization, err.Error())
	}

	var request dto.OrganizationRequest

	if err = c.Bind(&request); err != nil {
		return infra.NewErrorResponse(c, http.StatusBadRequest, constant.ErrBadRequest, constant.ErrBinding, err.Error())
	}

	if err := validation.ValidateRequest(request); err != nil {
		return infra.NewErrorResponse(c, http.StatusBadRequest, constant.ErrBadRequest, constant.ErrValidation, err)
	}

	orgExist.Name = request.Name
	orgExist.Description = request.Description
	orgExist.Type = request.Type
	orgExist.Email = request.Email
	orgExist.Website = request.Website
	orgExist.Phone = request.Phone

	if err = oc.OrganizationUsecase.Update(&orgExist); err != nil {
		return infra.NewErrorResponse(c, http.StatusInternalServerError, constant.ErrInternalServer, constant.ErrUpdateOrganization, err.Error())
	}

	response := entities.ToResponseOrganization(&orgExist)

	return infra.NewSuccessResponse(c, http.StatusOK, constant.SuccessOk, constant.SuccessUpdateOrganization, response)
}

func (oc *OrganizationController) Delete(c echo.Context) error {
	if err := oc.OrganizationUsecase.Delete(uuid.MustParse(c.Param("id"))); err != nil {
		return infra.NewErrorResponse(c, http.StatusNotFound, constant.ErrNotFound, constant.ErrDeleteOrganization, err.Error())
	}

	return infra.NewSuccessResponse(c, http.StatusOK, constant.SuccessOk, constant.SuccessDeleteOrganization, nil)
}

func (oc *OrganizationController) GetAllByUserId(c echo.Context) error {
	organizations, err := oc.OrganizationUsecase.GetAllByUserId(uuid.MustParse(c.Get("x-user-id").(string)))

	if err != nil {
		return infra.NewErrorResponse(c, http.StatusInternalServerError, constant.ErrInternalServer, constant.ErrGetOrganizationByUserId, err.Error())
	}

	response := entities.ToGetAllResponseOrganizations(organizations)

	return infra.NewSuccessResponse(c, http.StatusOK, constant.SuccessOk, constant.SuccessGetOrganizationByUserId, response)
}
