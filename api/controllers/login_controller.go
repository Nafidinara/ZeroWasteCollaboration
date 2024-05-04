package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	"redoocehub/bootstrap"
	"redoocehub/domains"
)

type LoginController struct {
	LoginUsecase domains.LoginUsecase
	Env          *bootstrap.Env
}

func (lc *LoginController) Login(c echo.Context) error {
	var request domains.LoginRequest

	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domains.ErrorResponse{Message: err.Error()})
		return err
	}

	user, err := lc.LoginUsecase.GetUserByEmail(c.Request().Context(), request.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, domains.ErrorResponse{Message: "User not found with the given email"})
		return err
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		c.JSON(http.StatusUnauthorized, domains.ErrorResponse{Message: "Invalid credentials"})
		return err
	}

	accessToken, err := lc.LoginUsecase.CreateAccessToken(&user, lc.Env.ACCESS_TOKEN_SECRET, lc.Env.ACCESS_TOKEN_EXPIRY_HOUR)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domains.ErrorResponse{Message: err.Error()})
		return err
	}

	refreshToken, err := lc.LoginUsecase.CreateRefreshToken(&user, lc.Env.REFRESH_TOKEN_SECRET, lc.Env.REFRESH_TOKEN_EXPIRY_HOUR)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domains.ErrorResponse{Message: err.Error()})
		return err
	}

	loginResponse := domains.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return c.JSON(http.StatusOK, loginResponse)
}
