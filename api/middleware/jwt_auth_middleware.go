package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"redoocehub/domains/infra"
	"redoocehub/internal/constant"
	"redoocehub/internal/tokenutil"
)

func JwtAuthMiddleware(secret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			t := strings.Split(authHeader, " ")
			if len(t) == 2 {
				authToken := t[1]
				authorized, err := tokenutil.IsAuthorized(authToken, secret)
				if authorized {
					userID, err := tokenutil.ExtractIDFromToken(authToken, secret)
					if err != nil {
						return infra.NewErrorResponse(c, http.StatusInternalServerError, constant.ErrInternalServer, constant.ErrFailedExtractToken, err.Error())
					}
					c.Set("x-user-id", userID)
					return next(c)
				}
				return infra.NewErrorResponse(c, http.StatusUnauthorized, constant.ErrUnauthorized, constant.ErrUnauthorized, err.Error())
			}
			return infra.NewErrorResponse(c, http.StatusBadRequest, constant.ErrBadRequest, constant.ErrUnauthorizedAuth, nil)
		}
	}
}
