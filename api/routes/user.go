package routes

import (
	"behealth-api/api/controllers"
	"behealth-api/infrastructure"
)

// UserRoutes struct
type UserRoutes struct {
	logger         infrastructure.Logger
	handler        infrastructure.RequestHandler
	userController controllers.UserController
}

// Setup user routes
func (s UserRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api")
	{
		api.GET("/user", s.userController.GetUser)
	}
}

// NewUserRoutes creates new user controller
func NewUserRoutes(
	logger infrastructure.Logger,
	handler infrastructure.RequestHandler,
	userController controllers.UserController,
) UserRoutes {
	return UserRoutes{
		handler:        handler,
		logger:         logger,
		userController: userController,
	}
}
