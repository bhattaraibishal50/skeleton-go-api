package controllers

import (
	"behealth-api/api/responses"
	"behealth-api/api/services"
	"behealth-api/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserController data type
type UserController struct {
	service services.UserService
	logger  infrastructure.Logger
}

// NewUserController creates new user controller
func NewUserController(userService services.UserService, logger infrastructure.Logger) UserController {
	return UserController{
		service: userService,
		logger:  logger,
	}
}

// GetUser gets the user
func (u UserController) GetUser(c *gin.Context) {
	responses.SuccessJSON(c, http.StatusOK, "get users")
}
