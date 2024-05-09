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
}

// func NewOrganizationRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, e *echo.Group) {
// 	ur := repositories.NewOrganizationRepository(db)

// 	uc := &controllers.OrganizationController{
// 		OrganizationUsecase: usecases.NewOrganizationUsecase(ur, timeout),
// 		Env:                 env,
// 	}

// 	e.GET("/organizations", uc.GetAll)
// 	e.GET("/organizations/:id", uc.GetByID)

// 	protectedRouter := e.Group("")
// 	protectedRouter.Use(middleware.JwtAuthMiddleware(env.ACCESS_TOKEN_SECRET))

// 	protectedRouter.POST("/organizations", uc.Create)
// 	protectedRouter.PUT("/organizations/:id", uc.Update)
// 	protectedRouter.DELETE("/organizations/:id", uc.Delete)
// }
