package handler

import (
	"net/http"
	"project-p-back/internal/entity"
	"project-p-back/internal/service"
	"project-p-back/pkg/jwtoken"
	"project-p-back/pkg/response"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.IUserService
}

func NewUserHandler(userService service.IUserService) *userHandler {
	var userHandler = userHandler{}
	userHandler.userService = userService
	return &userHandler
}

func (handler *userHandler) CreateUser(c *gin.Context) {
	var user entity.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = handler.userService.CreateUser(&user)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	response.ResponseCreated(c)
}

func (handler *userHandler) LoginUser(c *gin.Context) {
	var attempt entity.User
	err := c.ShouldBindJSON(&attempt)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := handler.userService.LoginUser(&attempt)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusNotFound)
		return
	}

	token, err := jwtoken.CreateToken(user.Id)
	if err != nil {
		response.ResponseError(c, err.Error(), http.StatusUnauthorized)
	}

	data := map[string]interface{}{
		"acess_token": token.AccessToken,
		"expired":     token.ExpiredToken,
		"user_id":     user.Id,
	}

	response.ResponseOKWithData(c, data)
}
