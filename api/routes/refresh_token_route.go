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

func NewRefreshTokenRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, e *echo.Group) {
	ur := repositories.NewUserRepository(db)
	rc := &controllers.RefreshTokenController{
		RefreshTokenUsecase: usecases.NewRefreshTokenUsecase(ur, timeout),
		Env:                 env,
	}
	e.POST("/refresh-token", rc.RefreshToken)
}
