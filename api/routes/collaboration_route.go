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
	or := repositories.NewOrganizationRepository(db)

	cc := &controllers.CollaborationController{
		CollaborationUsecase: usecases.NewCollaborationUsecase(cr, timeout),
		ProposalUsecase:      usecases.NewProposalUsecase(pr, timeout),
		OrganizationUsecase:  usecases.NewOrganizationUsecase(or, timeout),
		Env:                  env,
	}

	e.GET("/collaborations/:id", cc.GetByID)

	protectedRouter := e.Group("")

	protectedRouter.Use(middleware.JwtAuthMiddleware(env.ACCESS_TOKEN_SECRET))

	protectedRouter.GET("/collaborations", cc.GetAllByUserId)
	protectedRouter.POST("/collaborations", cc.Create)
	protectedRouter.PUT("/collaborations/:id", cc.Update)
	protectedRouter.DELETE("/collaborations/:id", cc.Delete)
}
