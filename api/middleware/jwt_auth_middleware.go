package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"

	"redoocehub/domains/infra"
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
						return c.JSON(http.StatusUnauthorized, infra.ErrorResponse{
							StatusCode: "Unauthorized",
							Message:    err.Error(),
							Data:       nil,
						})
					}
					c.Set("x-user-id", userID)
					return next(c)
				}
				return c.JSON(http.StatusUnauthorized, infra.ErrorResponse{
					StatusCode: "Unauthorized",
					Message:    err.Error(),
					Data:       nil,
				})
			}
			return c.JSON(http.StatusUnauthorized, infra.ErrorResponse{
				StatusCode: "Unauthorized",
				Message:    "You are not authorized to access this resource",
				Data:       nil,
			})
		}
	}
}
