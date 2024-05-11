package dto

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type File struct {
	File multipart.File `json:"file,omitempty" validate:"required"`
}

type Url struct {
	Url string `json:"url,omitempty" validate:"required"`
}

type MediaDto struct {
	StatusCode int       `json:"statusCode"`
	Message    string    `json:"message"`
	Data       *echo.Map `json:"data"`
}
