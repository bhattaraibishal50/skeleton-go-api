package routes

import (
	"behealth-api/api/controllers"
	"behealth-api/lib"
)

// AuthRoutes struct
type AuthRoutes struct {
	logger         lib.Logger
	handler        lib.RequestHandler
	authController controllers.JWTAuthController
}

// Setup user routes
func (s AuthRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	auth := s.handler.Gin.Group("/auth")
	{
		auth.POST("/login", s.authController.SignIn)
		auth.POST("/register", s.authController.Register)
	}
}

// NewAuthRoutes creates new user controller
func NewAuthRoutes(
	handler lib.RequestHandler,
	authController controllers.JWTAuthController,
	logger lib.Logger,
) AuthRoutes {
	return AuthRoutes{
		handler:        handler,
		logger:         logger,
		authController: authController,
	}
}
