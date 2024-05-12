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

func NewOrganizationRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, e *echo.Group) {
	ur := repositories.NewOrganizationRepository(db)

	uc := &controllers.OrganizationController{
		OrganizationUsecase: usecases.NewOrganizationUsecase(ur, timeout),
		Env:                 env,
	}

	e.GET("/organizations", uc.GetAll)
	e.GET("/organizations/:id", uc.GetByID)

	protectedRouter := e.Group("")
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.ACCESS_TOKEN_SECRET))

	protectedRouter.GET("/users/organizations", uc.GetAllByUserId)
	protectedRouter.POST("/organizations", uc.Create)
	protectedRouter.PUT("/organizations/:id", uc.Update)
	protectedRouter.DELETE("/organizations/:id", uc.Delete)
}
