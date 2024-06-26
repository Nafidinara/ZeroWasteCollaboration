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

func NewUserRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, e *echo.Group) {
	ur := repositories.NewUserRepository(db)

	uc := &controllers.UserController{
		UserUsecase:    usecases.NewUserUsecase(ur, timeout),
		ChatbotUsecase: usecases.NewChatbotUsecase(timeout),
		Env:            env,
	}

	userRouter := e.Group("/users")

	userRouter.POST("/login", uc.Login)
	userRouter.POST("/register", uc.Register)
	userRouter.POST("/refresh", uc.RefreshToken)

	protectedRouter := userRouter.Group("")
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.ACCESS_TOKEN_SECRET))
	protectedRouter.GET("/profile", uc.Profile)
	protectedRouter.PUT("/profile", uc.Update)
	protectedRouter.GET("/dashboard", uc.Dashboard)
	protectedRouter.POST("/chatbot", uc.SendMessageChatbot)
}
