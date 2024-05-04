package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"redoocehub/bootstrap"
	"redoocehub/domains"
)

type RefreshTokenController struct {
	RefreshTokenUsecase domains.RefreshTokenUsecase
	Env                 *bootstrap.Env
}

func (rc *RefreshTokenController) RefreshToken(c echo.Context) error {
	var request domains.RefreshTokenRequest
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, domains.ErrorResponse{Message: err.Error()})
	}

	user, err := rc.RefreshTokenUsecase.GetUserByID(c.Request().Context(), request.RefreshToken)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, domains.ErrorResponse{Message: err.Error()})
	}

	accessToken, err := rc.RefreshTokenUsecase.CreateAccessToken(&user, rc.Env.ACCESS_TOKEN_SECRET, rc.Env.ACCESS_TOKEN_EXPIRY_HOUR)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, domains.ErrorResponse{Message: err.Error()})
	}

	refreshToken, err := rc.RefreshTokenUsecase.CreateRefreshToken(&user, rc.Env.REFRESH_TOKEN_SECRET, rc.Env.REFRESH_TOKEN_EXPIRY_HOUR)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, domains.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, domains.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
