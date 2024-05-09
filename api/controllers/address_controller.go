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

	userId := c.Get("x-user-id").(string)

	request.UserId = uuid.MustParse(userId)

	newAddress, err := ac.AddressUsecase.CreateUserAddress(&request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
			StatusCode: "Internal Server Error",
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, infra.SuccessResponse{
		StatusCode: "OK",
		Message:    "Success created new address",
		Data:       newAddress,
	})
}

func (ac *AddressController) CreateOrganizationAddress(c echo.Context) error {
	var request dto.OrganizationAddressRequest

	if err := c.Bind(&request); err != nil {
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

	newAddress, err := ac.AddressUsecase.CreateOrganizationAddress(&request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
			StatusCode: "Internal Server Error",
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, infra.SuccessResponse{
		StatusCode: "OK",
		Message:    "Success created new address",
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
			StatusCode: "Bad Request",
			Message:    "make sure you input the parameter organization_id or user_id",
			Data:       nil,
		})
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
			StatusCode: "Internal Server Error",
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, infra.SuccessResponse{
		StatusCode: "OK",
		Message:    "Success retrieved all address",
		Data:       address,
	})
}

// func (oc *OrganizationController) GetAll(c echo.Context) error {
// 	organizations, err := oc.OrganizationUsecase.GetAll()

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
// 			StatusCode: "Internal Server Error",
// 			Message:    err.Error(),
// 			Data:       nil,
// 		})
// 	}

// 	response := entities.ToGetAllResponseOrganizations(organizations)

// 	return c.JSON(http.StatusOK, infra.SuccessResponse{
// 		StatusCode: "OK",
// 		Message:    "Success retrieved all organizations",
// 		Data:       response,
// 	})
// }

// func (oc *OrganizationController) GetByID(c echo.Context) error {

// 	idParam := c.Param("id")

// 	id := uuid.MustParse(idParam)

// 	organization, err := oc.OrganizationUsecase.GetByID(id)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
// 			StatusCode: "Internal Server Error",
// 			Message:    err.Error(),
// 			Data:       nil,
// 		})
// 	}

// 	response := entities.ToResponseOrganizationDetail(&organization, &organization.User)

// 	return c.JSON(http.StatusOK, infra.SuccessResponse{
// 		StatusCode: "OK",
// 		Message:    "Success retrieved organization",
// 		Data:       response,
// 	})
// }

// func (oc *OrganizationController) Create(c echo.Context) error {

// 	var request dto.OrganizationRequest

// 	err := c.Bind(&request)

// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, infra.ErrorResponse{
// 			StatusCode: "Bad Request",
// 			Message:    err.Error(),
// 			Data:       nil,
// 		})
// 	}

// 	if err := validation.ValidateRequest(request); err != nil {
// 		return c.JSON(http.StatusBadRequest, infra.ErrorResponse{
// 			StatusCode: "Bad Request",
// 			Message:    "make sure you follow the input requirements",
// 			Data:       err,
// 		})
// 	}

// 	userId := c.Get("x-user-id").(string)

// 	request.UserID = uuid.MustParse(userId)

// 	organization, err := oc.OrganizationUsecase.Create(&request)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
// 			StatusCode: "Internal Server Error",
// 			Message:    err.Error(),
// 			Data:       nil,
// 		})
// 	}

// 	userData, err := oc.OrganizationUsecase.GetUser(organization.UserID)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
// 			StatusCode: "Internal Server Error",
// 			Message:    err.Error(),
// 			Data:       nil,
// 		})
// 	}

// 	response := entities.ToResponseOrganizationDetail(organization, &userData)

// 	return c.JSON(http.StatusOK, infra.SuccessResponse{
// 		StatusCode: "OK",
// 		Message:    "Success created organization",
// 		Data:       response,
// 	})
// }

// func (oc *OrganizationController) Update(c echo.Context) error {
// 	idParam := c.Param("id")

// 	id := uuid.MustParse(idParam)

// 	orgExist, err := oc.OrganizationUsecase.GetByID(id)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
// 			StatusCode: "Internal Server Error",
// 			Message:    err.Error(),
// 			Data:       nil,
// 		})
// 	}

// 	if orgExist.ID == uuid.Nil {
// 		return c.JSON(http.StatusNotFound, infra.ErrorResponse{
// 			StatusCode: "Not Found",
// 			Message:    "Organization not found",
// 			Data:       nil,
// 		})
// 	}

// 	var request dto.OrganizationRequest

// 	err = c.Bind(&request)

// 	if err := validation.ValidateRequest(request); err != nil {
// 		return c.JSON(http.StatusBadRequest, infra.ErrorResponse{
// 			StatusCode: "Bad Request",
// 			Message:    "make sure you follow the input requirements",
// 			Data:       err,
// 		})
// 	}

// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, infra.ErrorResponse{
// 			StatusCode: "Bad Request",
// 			Message:    err.Error(),
// 			Data:       nil,
// 		})
// 	}

// 	orgExist.Name = request.Name
// 	orgExist.Description = request.Description
// 	orgExist.Type = request.Type
// 	orgExist.Email = request.Email
// 	orgExist.Website = request.Website
// 	orgExist.Phone = request.Phone

// 	err = oc.OrganizationUsecase.Update(&orgExist)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
// 			StatusCode: "Internal Server Error",
// 			Message:    err.Error(),
// 			Data:       nil,
// 		})
// 	}

// 	response := entities.ToResponseOrganization(&orgExist)

// 	return c.JSON(http.StatusOK, infra.SuccessResponse{
// 		StatusCode: "OK",
// 		Message:    "Success updated organization",
// 		Data:       response,
// 	})
// }

// func (oc *OrganizationController) Delete(c echo.Context) error {
// 	idParam := c.Param("id")

// 	id := uuid.MustParse(idParam)

// 	err := oc.OrganizationUsecase.Delete(id)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
// 			StatusCode: "Internal Server Error",
// 			Message:    err.Error(),
// 			Data:       nil,
// 		})
// 	}

// 	return c.JSON(http.StatusOK, infra.SuccessResponse{
// 		StatusCode: "OK",
// 		Message:    "Success deleted organization",
// 		Data:       nil,
// 	})
// }
