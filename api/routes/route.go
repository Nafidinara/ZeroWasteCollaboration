package routes

import (
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"redoocehub/bootstrap"
)

func SetupRoutes(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, e *echo.Echo) {
	prefixRouter := e.Group("api/v1")

	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, "Welcome to redoocehub API service, v.1!")
	})
	
	NewUserRouter(env, timeout, db, prefixRouter)
	NewOrganizationRouter(env, timeout, db, prefixRouter)
	NewAddressRouter(env, timeout, db, prefixRouter)
	NewCollaborationRouter(env, timeout, db, prefixRouter)
}
