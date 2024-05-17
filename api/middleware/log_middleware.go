package middleware

import "github.com/labstack/echo/v4"
import "github.com/labstack/echo/v4/middleware"

type LoggerConfig struct {
	Format string
}

func (c *LoggerConfig) Init() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: c.Format,
	})
}