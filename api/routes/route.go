package routes

import (
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"redoocehub/api/middleware"
	"redoocehub/bootstrap"
)

func SetupRoutes(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, e *echo.Echo) {
	publicRouter := e.Group("")	
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)

	protectedRouter := e.Group("")
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.ACCESS_TOKEN_SECRET))

	NewProfileRouter(env, timeout, db, protectedRouter)
}
