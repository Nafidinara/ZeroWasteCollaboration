package routes

import (
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"redoocehub/api/controllers"
	"redoocehub/bootstrap"
	"redoocehub/repositories"
	"redoocehub/usecases"
)

func NewProfileRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, e *echo.Group) {
	ur := repositories.NewUserRepository(db)
	pc := &controllers.ProfileController{
		ProfileUsecase: usecases.NewProfileUsecase(ur, timeout),
	}

	e.GET("/profile", pc.Fetch)
}
