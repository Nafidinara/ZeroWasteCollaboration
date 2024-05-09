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

func NewCollaborationRouter(env *bootstrap.Env, timeout time.Duration, db *gorm.DB, e *echo.Group) {
	cr := repositories.NewCollaborationRepository(db)

	pr := repositories.NewProposalRepository(db)

	cc := &controllers.CollaborationController{
		CollaborationUsecase: usecases.NewCollaborationUsecase(cr, timeout),
		ProposalUsecase:      usecases.NewProposalUsecase(pr, timeout),
		Env:                  env,
	}

	e.GET("/collaborations/:id", cc.GetByID)

	protectedRouter := e.Group("")

	protectedRouter.Use(middleware.JwtAuthMiddleware(env.ACCESS_TOKEN_SECRET))

	protectedRouter.POST("/collaborations", cc.Create)
}
