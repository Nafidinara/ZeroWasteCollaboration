package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	"redoocehub/bootstrap"
	"redoocehub/domains/dto"
	"redoocehub/domains/entities"
	"redoocehub/domains/infra"
)

type UserController struct {
	UserUsecase entities.UserUsecase
	Env         *bootstrap.Env
}

func (uc *UserController) Register(c echo.Context) error {

	var request dto.RegisterRequest

	err := c.Bind(&request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, infra.ErrorResponse{
			StatusCode: "Bad Request",
			Message:    err.Error(),
			Data:       nil,
		})
	}

	_, err = uc.UserUsecase.GetUserByEmail(c.Request().Context(), request.Email)

	if err == nil {
		return c.JSON(http.StatusConflict, infra.ErrorResponse{
			StatusCode: "Conflict",
			Message:    "User already exists",
			Data:       nil,
		})
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
			StatusCode: "Internal Server Error",
			Message:    err.Error(),
			Data:       nil,
		})
	}

	request.Password = string(encryptedPassword)

	user, err := uc.UserUsecase.Create(c.Request().Context(), &request)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
			StatusCode: "Internal Server Error",
			Message:    err.Error(),
		})
	}

	accessToken, err := uc.UserUsecase.CreateAccessToken(user, uc.Env.ACCESS_TOKEN_SECRET, uc.Env.ACCESS_TOKEN_EXPIRY_HOUR)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
			StatusCode: "Internal Server Error",
			Message:    err.Error(),
			Data:       nil,
		})
	}

	refreshToken, err := uc.UserUsecase.CreateRefreshToken(user, uc.Env.REFRESH_TOKEN_SECRET, uc.Env.REFRESH_TOKEN_EXPIRY_HOUR)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
			StatusCode: "Internal Server Error",
			Message:    err.Error(),
			Data:       nil,
		})
	}

	user.RefreshToken = refreshToken

	response := entities.ToRegisterResponseUser(user, accessToken)

	return c.JSON(http.StatusCreated, infra.SuccessResponse{
		StatusCode: "Created",
		Message:    "User created successfully",
		Data:       response,
	})
}

func (uc *UserController) Login(c echo.Context) error {
	var request dto.LoginRequest

	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, infra.ErrorResponse{
			StatusCode: "Bad Request",
			Message:    err.Error(),
			Data:       nil,
		})
	}

	user, err := uc.UserUsecase.GetUserByEmail(c.Request().Context(), request.Email)
	
	if err != nil {
		return c.JSON(http.StatusNotFound, infra.ErrorResponse{
			StatusCode: "Not Found",
			Message:    "User not found with the given email",
			Data:       nil,
		})
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		return c.JSON(http.StatusUnauthorized, infra.ErrorResponse{
			StatusCode: "Unauthorized",
			Message:    "Invalid credentials",
			Data:       nil,
		})
	}

	accessToken, err := uc.UserUsecase.CreateAccessToken(&user, uc.Env.ACCESS_TOKEN_SECRET, uc.Env.ACCESS_TOKEN_EXPIRY_HOUR)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
			StatusCode: "Internal Server Error",
			Message:    err.Error(),
			Data:       nil,
		})
	}

	refreshToken, err := uc.UserUsecase.CreateRefreshToken(&user, uc.Env.REFRESH_TOKEN_SECRET, uc.Env.REFRESH_TOKEN_EXPIRY_HOUR)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
			StatusCode: "Internal Server Error",
			Message:    err.Error(),
			Data:       nil,
		})
	}

	loginUser := entities.EntityToDtoUser(&user)

	loginResponse := entities.ToLoginResponseUser(loginUser, accessToken, refreshToken)

	return c.JSON(http.StatusOK, infra.SuccessResponse{
		StatusCode: "OK",
		Message:    "Login successful",
		Data:       loginResponse,
	})
}

func (uc *UserController) Profile(c echo.Context) error {
	userID := c.Get("x-user-id").(string)

	profile, err := uc.UserUsecase.GetProfileByID(c.Request().Context(), userID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
			StatusCode: "Internal Server Error",
			Message:    err.Error(),
			Data:       nil,
		})
	}

	response := entities.ToProfileResponseUser(profile)

	return c.JSON(http.StatusOK, infra.SuccessResponse{
		StatusCode: "OK",
		Message:    "Profile retrieved successfully",
		Data:       response,
	})
}

func (uc *UserController) RefreshToken(c echo.Context) error {
	var request dto.RefreshTokenRequest

	err := c.Bind(&request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, infra.ErrorResponse{
			StatusCode: "Bad Request",
			Message:    err.Error(),
			Data:       nil,
		})
	}

	user, err := uc.UserUsecase.GetUserByID(c.Request().Context(), request.RefreshToken)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, infra.ErrorResponse{
			StatusCode: "Unauthorized",
			Message:    err.Error(),
			Data:       nil,
		})
	}

	accessToken, err := uc.UserUsecase.CreateAccessToken(&user, uc.Env.ACCESS_TOKEN_SECRET, uc.Env.ACCESS_TOKEN_EXPIRY_HOUR)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
			StatusCode: "Internal Server Error",
			Message:    err.Error(),
			Data:       nil,
		})
	}

	refreshToken, err := uc.UserUsecase.CreateRefreshToken(&user, uc.Env.REFRESH_TOKEN_SECRET, uc.Env.REFRESH_TOKEN_EXPIRY_HOUR)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, infra.ErrorResponse{
			StatusCode: "Internal Server Error",
			Message:    err.Error(),
			Data:       nil,
		})
	}

	response := entities.ToRefreshTokenResponseUser(accessToken, refreshToken)

	return c.JSON(http.StatusOK, infra.SuccessResponse{
		StatusCode: "OK",
		Message:    "Token refreshed successfully",
		Data:       response,
	})
}
