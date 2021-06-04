package controllers

import (
	"behealth-api/api/responses"
	"behealth-api/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserController data type
type UserController struct {
	logger infrastructure.Logger
}

// NewUserController creates new user controller
func NewUserController(logger infrastructure.Logger) UserController {
	return UserController{
		logger: logger,
	}
}

// GetUser gets the user
func (u UserController) GetUser(c *gin.Context) {
	responses.SuccessJSON(c, http.StatusOK, "get users")
}
