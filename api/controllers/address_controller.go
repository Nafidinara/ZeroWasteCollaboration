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
	Env            *bootstrap.Env
}

func NewAddressController(addressUsecase entities.AddressUsecase) *AddressController {
	return &AddressController{AddressUsecase: addressUsecase}
}

func (ac *AddressController) CreateUserAddress(c echo.Context) error {
	var request dto.UserAddressRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, infra.ErrorResponse{
			StatusCode: constant.ErrBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	if err := validation.ValidateRequest(request); err != nil {
		return c.JSON(http.StatusBadRequest, infra.ErrorResponse{
			StatusCode: constant.ErrBadRequest,
			Message:    constant.ErrValidation,
			Data:       err,
		})
	}

	userId := c.Get("x-user-id").(string)

	request.UserId = uuid.MustParse(userId)

	newAddress, err := ac.AddressUsecase.CreateUserAddress(&request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
			StatusCode: constant.ErrInternalServer,
			Message:    constant.ErrCreateUserAddress,
			Data:       err.Error(),
		})
	}

	return c.JSON(http.StatusOK, infra.SuccessResponse{
		StatusCode: constant.SuccessCreated,
		Message:    constant.SuccessCreateAddress,
		Data:       newAddress,
	})
}

func (ac *AddressController) CreateOrganizationAddress(c echo.Context) error {
	var request dto.OrganizationAddressRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, infra.ErrorResponse{
			StatusCode: constant.ErrBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	if err := validation.ValidateRequest(request); err != nil {
		return c.JSON(http.StatusBadRequest, infra.ErrorResponse{
			StatusCode: constant.ErrBadRequest,
			Message:    constant.ErrValidation,
			Data:       err,
		})
	}

	newAddress, err := ac.AddressUsecase.CreateOrganizationAddress(&request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
			StatusCode: constant.ErrInternalServer,
			Message:    constant.ErrCreateOrganizationAddress,
			Data:       err.Error(),
		})
	}

	return c.JSON(http.StatusOK, infra.SuccessResponse{
		StatusCode: constant.SuccessCreated,
		Message:    constant.SuccessCreateAddress,
		Data:       newAddress,
	})
}

func (ac *AddressController) GetAllAddress(c echo.Context) error {
	//get organization_id from param
	organizationId := c.QueryParam("organization_id")
	userId := c.QueryParam("user_id")

	var address []entities.Address
	var err error

	fmt.Println(organizationId, userId)

	if organizationId != "" {
		address, err = ac.AddressUsecase.GetAllOrganizationAddress(uuid.MustParse(organizationId))
	} else if userId != "" {
		address, err = ac.AddressUsecase.GetAllUserAddress(uuid.MustParse(userId))
	} else {
		return c.JSON(http.StatusBadRequest, infra.ErrorResponse{
			StatusCode: constant.ErrBadRequest,
			Message:    constant.ErrParameterNotFound,
			Data:       nil,
		})
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
			StatusCode: constant.ErrInternalServer,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, infra.SuccessResponse{
		StatusCode: constant.SuccessOk,
		Message:    constant.ErrInternalServer,
		Data:       address,
	})
}

// delete
func (ac *AddressController) Delete(c echo.Context) error {

	idParam := c.Param("id")

	id := uuid.MustParse(idParam)

	err := ac.AddressUsecase.Delete(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
			StatusCode: constant.ErrInternalServer,
			Message:    constant.ErrDeleteAddress,
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, infra.SuccessResponse{
		StatusCode: constant.SuccessOk,
		Message:    constant.SuccessDeleteAddress,
		Data:       nil,
	})
}
