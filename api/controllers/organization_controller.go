package controllers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"redoocehub/domains/infra"
	"redoocehub/domains/organization/dto"
	"redoocehub/domains/organization/entities"
)

type OrganizationController struct {
	OrganizationUsecase entities.OrganizationUsecase
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
		Message:    "Success",
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
		Message:    "Success",
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

	organization := entities.Organization{
		ID:       uuid.New(),
		Name:        request.Name,
		Description: request.Description,
		Type:        request.Type,
		Email:       request.Email,
		Website:     request.Website,
		Phone:       request.Phone,
	}

	err = oc.OrganizationUsecase.Create(&organization)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
			StatusCode: "Internal Server Error",
			Message:    err.Error(),
			Data:       nil,
		})
	}

	response := dto.OrganizationResponse{
		ID:          organization.ID,
		Name:        organization.Name,
		Description: organization.Description,
		Type:        organization.Type,
		Email:       organization.Email,
		Website:     organization.Website,
		Phone:       organization.Phone,
	}

	return c.JSON(http.StatusOK, infra.SuccessResponse{
		StatusCode: "OK",
		Message:    "Success",
		Data:       response,
	})
}

func (oc *OrganizationController) Update(c echo.Context) error {

	var request dto.OrganizationRequest

	err := c.Bind(&request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, infra.ErrorResponse{
			StatusCode: "Bad Request",
			Message:    err.Error(),
			Data:       nil,
		})
	}

	organization := entities.Organization{
		Name:        request.Name,
		Description: request.Description,
		Type:        request.Type,
		Email:       request.Email,
		Website:     request.Website,
		Phone:       request.Phone,
	}

	err = oc.OrganizationUsecase.Update(&organization)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
			StatusCode: "Internal Server Error",
			Message:    err.Error(),
			Data:       nil,
		})
	}

	response := dto.OrganizationResponse{
		ID:          organization.ID,
		Name:        organization.Name,
		Description: organization.Description,
		Type:        organization.Type,
		Email:       organization.Email,
		Website:     organization.Website,
		Phone:       organization.Phone,
	}

	return c.JSON(http.StatusOK, infra.SuccessResponse{
		StatusCode: "OK",
		Message:    "Success",
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
		Message:    "Success",
		Data:       nil,
	})
}