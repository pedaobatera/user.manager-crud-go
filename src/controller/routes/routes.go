package routes

import (
	"github.com/gin-gonic/gin"
	"user.manager-crud-go/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/getUserById/:userId", controller.FindUserById)
	r.GET("/getAllUsers", controller.FindAllUsers)
	r.GET("/getUsersByEmail/:userEmail", controller.FindUsersByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)

}
