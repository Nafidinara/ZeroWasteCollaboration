package routes

import (
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"redoocehub/api/controllers"
	"redoocehub/api/middleware"
	"redoocehub/bootstrap"
	"redoocehub/repositories"
	"redoocehub/usecases"
)

func NewAddressRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, e *echo.Group) {
	ar := repositories.NewAddressRepository(db)

	ac := &controllers.AddressController{
		AddressUsecase: usecases.NewAddressUsecase(ar, timeout),
		Env:            env,
	}

	e.GET("/addresses", ac.GetAllAddress)

	protectedRouter := e.Group("")
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.ACCESS_TOKEN_SECRET))

	protectedRouter.POST("/user-addresses", ac.CreateUserAddress)
	protectedRouter.POST("/organization-addresses", ac.CreateOrganizationAddress)
	protectedRouter.DELETE("/addresses/:id", ac.Delete)
}
