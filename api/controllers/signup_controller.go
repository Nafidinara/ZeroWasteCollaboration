package controllers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	"redoocehub/bootstrap"
	"redoocehub/domains"
)

type SignupController struct {
	SignupUsecase domains.SignupUsecase
	Env           *bootstrap.Env
}

func (sc *SignupController) Signup(c echo.Context) error {
	var request domains.SignupRequest

	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domains.ErrorResponse{Message: err.Error()})
	}

	_, err = sc.SignupUsecase.GetUserByEmail(c.Request().Context(), request.Email)

	if err == nil {
		return c.JSON(http.StatusConflict, domains.ErrorResponse{Message: "User already exists"})
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, domains.ErrorResponse{Message: err.Error()})
	}

	request.Password = string(encryptedPassword)

	user := &domains.User{
		ID:       uuid.New(),
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	err = sc.SignupUsecase.Create(c.Request().Context(), user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, domains.ErrorResponse{Message: err.Error()})
	}

	accessToken, err := sc.SignupUsecase.CreateAccessToken(user, sc.Env.ACCESS_TOKEN_SECRET, sc.Env.ACCESS_TOKEN_EXPIRY_HOUR)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, domains.ErrorResponse{Message: err.Error()})
	}

	refreshToken, err := sc.SignupUsecase.CreateRefreshToken(user, sc.Env.REFRESH_TOKEN_SECRET, sc.Env.REFRESH_TOKEN_EXPIRY_HOUR)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, domains.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, domains.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
