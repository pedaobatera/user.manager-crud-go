package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"user.manager-crud-go/src/configuration/logger"
	"user.manager-crud-go/src/configuration/rest_err"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info("Init deleteUser controller",
		zap.String("journey", "deleteUser"),
	)

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid user id, must be a valid hex")
		c.JSON(errRest.Code, errRest)
	}

	err := uc.service.DeleteUserService(userId)
	if err != nil {
		logger.Error(
			"Error trying to deleteUser user",
			err,
			zap.String("journey", "deleteUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"User deleteUser successfully",
		zap.String("user_id", userId),
		zap.String("journey", "deleteUser"))

	c.Status(http.StatusNoContent)
}
