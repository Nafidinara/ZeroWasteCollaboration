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
		AddressUsecase:      usecases.NewAddressUsecase(ar, timeout),
		OrganizationUsecase: usecases.NewOrganizationUsecase(repositories.NewOrganizationRepository(db), timeout),
		Env:                 env,
	}

	e.GET("/addresses", ac.GetAllAddress)

	protectedRouter := e.Group("")
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.ACCESS_TOKEN_SECRET))

	protectedRouter.POST("/users/address", ac.CreateUserAddress)
	protectedRouter.POST("/organizations/address", ac.CreateOrganizationAddress)
	protectedRouter.DELETE("/addresses/:id", ac.Delete)
}
