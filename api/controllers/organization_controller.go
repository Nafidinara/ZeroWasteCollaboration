package controllers

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"redoocehub/bootstrap"
	"redoocehub/domains/dto"
	"redoocehub/domains/entities"
	"redoocehub/domains/infra"
	"redoocehub/internal/validation"
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

	return c.JSON(http.StatusOK, infra.SuccessResponse{
		StatusCode: "OK",
		Message:    "Success retrieved all organizations",
		Data:       organizations,
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

	return c.JSON(http.StatusOK, infra.SuccessResponse{
		StatusCode: "OK",
		Message:    "Success retrieved organization",
		Data:       organization,
	})
}

func (oc *OrganizationController) Create(c echo.Context) error {

	var request dto.OrganizationRequest

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

	fundingDate, err := time.Parse("2006-01-02", request.FoundingDate)

	if err != nil {
		return c.JSON(http.StatusBadRequest, infra.ErrorResponse{
			StatusCode: "Bad Request",
			Message:    err.Error(),
			Data:       nil,
		})
	}

	userId := c.Get("x-user-id").(string)

	organization := entities.Organization{
		ID:           uuid.New(),
		UserID:       uuid.MustParse(userId),
		Name:         request.Name,
		Description:  request.Description,
		Type:         request.Type,
		Email:        request.Email,
		ProfileImage: request.ProfileImage,
		FoundingDate: fundingDate,
		Website:      request.Website,
		Phone:        request.Phone,
	}

	err = oc.OrganizationUsecase.Create(&organization)

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

	user := entities.User{
		ID:       userData.ID,
		Email:    userData.Email,
		Username: userData.Username,
		FullName: userData.FullName,
		Gender:   userData.Gender,
	}

	response := dto.OrganizationResponse{
		ID:           organization.ID,
		Name:         organization.Name,
		Description:  organization.Description,
		Type:         organization.Type,
		Email:        organization.Email,
		ProfileImage: organization.ProfileImage,
		FoundingDate: fundingDate,
		Website:      organization.Website,
		Phone:        organization.Phone,
		CreatedAt:    organization.CreatedAt,
		UpdatedAt:    organization.UpdatedAt,
		DeletedAt:    organization.DeletedAt,
		User:         user,
	}

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

	response := dto.OrganizationResponse{
		ID:          orgExist.ID,
		Name:        orgExist.Name,
		Description: orgExist.Description,
		Type:        orgExist.Type,
		Email:       orgExist.Email,
		Website:     orgExist.Website,
		Phone:       orgExist.Phone,
	}

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
