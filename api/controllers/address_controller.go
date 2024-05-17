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
	"redoocehub/internal/constant"
	"redoocehub/internal/validation"
)

type AddressController struct {
	AddressUsecase entities.AddressUsecase
	OrganizationUsecase entities.OrganizationUsecase
	Env            *bootstrap.Env
}

func NewAddressController(addressUsecase entities.AddressUsecase) *AddressController {
	return &AddressController{AddressUsecase: addressUsecase}
}

func (ac *AddressController) CreateUserAddress(c echo.Context) error {
	var request dto.UserAddressRequest

	if err := c.Bind(&request); err != nil {
		return infra.NewErrorResponse(c, http.StatusBadRequest, constant.ErrBadRequest, constant.ErrBinding, err.Error())
	}

	if err := validation.ValidateRequest(request); err != nil {
		return infra.NewErrorResponse(c, http.StatusBadRequest, constant.ErrBadRequest, constant.ErrValidation, err)
	}

	request.UserId = uuid.MustParse(c.Get("x-user-id").(string))

	newAddress, err := ac.AddressUsecase.CreateUserAddress(&request)

	if err != nil {
		return infra.NewErrorResponse(c, http.StatusInternalServerError, constant.ErrInternalServer, constant.ErrCreateUserAddress, err.Error())
	}

	return infra.NewSuccessResponse(c, http.StatusOK, constant.SuccessOk, constant.SuccessCreateAddress, newAddress)
}

func (ac *AddressController) CreateOrganizationAddress(c echo.Context) error {
	var request dto.OrganizationAddressRequest

	if err := c.Bind(&request); err != nil {
		return infra.NewErrorResponse(c, http.StatusBadRequest, constant.ErrBadRequest, constant.ErrBinding, err.Error())
	}

	if err := validation.ValidateRequest(request); err != nil {
		return infra.NewErrorResponse(c, http.StatusBadRequest, constant.ErrBadRequest, constant.ErrValidation, err)
	}

	organization, err := ac.OrganizationUsecase.GetByID(request.OrganizationId)

	if err != nil || organization.ID == uuid.Nil {
		return infra.NewErrorResponse(c, http.StatusNotFound, constant.ErrNotFound, constant.ErrGetOrganization, err.Error())
	}

	newAddress, err := ac.AddressUsecase.CreateOrganizationAddress(&request)

	if err != nil {
		return infra.NewErrorResponse(c, http.StatusInternalServerError, constant.ErrInternalServer, constant.ErrCreateOrganizationAddress, err.Error())
	}

	return infra.NewSuccessResponse(c, http.StatusOK, constant.SuccessOk, constant.SuccessCreateAddress, newAddress)
}

func (ac *AddressController) GetAllAddress(c echo.Context) error {
	//get organization_id from param
	organizationId := c.QueryParam("organization_id")
	userId := c.QueryParam("user_id")

	var address []entities.Address
	var err error

	if organizationId != "" {
		address, err = ac.AddressUsecase.GetAllOrganizationAddress(uuid.MustParse(organizationId))
	} else if userId != "" {
		address, err = ac.AddressUsecase.GetAllUserAddress(uuid.MustParse(userId))
	} else {
		return infra.NewErrorResponse(c, http.StatusBadRequest, constant.ErrBadRequest, constant.ErrParameterNotFound, fmt.Sprintf("organization_id: %s, user_id: %s", organizationId, userId))
	}

	if err != nil {
		return infra.NewErrorResponse(c, http.StatusInternalServerError, constant.ErrInternalServer, constant.ErrGetAllAddress, err.Error())
	}

	return infra.NewSuccessResponse(c, http.StatusOK, constant.SuccessOk, constant.SuccessGetAllAddress, address)
}

// delete
func (ac *AddressController) Delete(c echo.Context) error {
	if err := ac.AddressUsecase.Delete(uuid.MustParse(c.Param("id"))); err != nil {
		return infra.NewErrorResponse(c, http.StatusNotFound, constant.ErrNotFound, constant.ErrDeleteAddress, err.Error())
	}

	return infra.NewSuccessResponse(c, http.StatusOK, constant.SuccessOk, constant.SuccessDeleteAddress, nil)
}
