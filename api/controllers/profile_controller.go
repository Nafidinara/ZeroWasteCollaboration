package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"redoocehub/domains"
)

type ProfileController struct {
	ProfileUsecase domains.ProfileUsecase
}

func (pc *ProfileController) Fetch(c echo.Context) error {
	userID := c.Get("x-user-id").(string)

	profile, err := pc.ProfileUsecase.GetProfileByID(c.Request().Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domains.ErrorResponse{Message: err.Error()})
		return err
	}

	return c.JSON(http.StatusOK, profile)
}
