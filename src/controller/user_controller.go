package controller

import (
	"github.com/gin-gonic/gin"
	"user.manager-crud-go/src/model/service"
)

func NewUserControllerInterface(
	service service.UserDomainService,
) UserControllerInterface {
	return &userControllerInterface{
		service: service,
	}
}

type UserControllerInterface interface {
	FindUserById(c *gin.Context)
	FindUsersByEmail(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	LoginUser(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}
