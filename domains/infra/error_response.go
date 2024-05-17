package infra

import "github.com/labstack/echo/v4"

type ErrorResponse struct {
	StatusCode string `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}

func NewErrorResponse(c echo.Context, httpCode int, StatusCode string, Message string, Data any) error {
	return c.JSON(httpCode, SuccessResponse{
		StatusCode: StatusCode,
		Message:    Message,
		Data:       Data,
	})
}
