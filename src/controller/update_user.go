package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"user.manager-crud-go/src/configuration/logger"
	"user.manager-crud-go/src/configuration/rest_err"
	"user.manager-crud-go/src/configuration/validation"
	"user.manager-crud-go/src/controller/model/request"
	"user.manager-crud-go/src/model"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("Init updateUser controller",
		zap.String("journey", "updateUser"),
	)

	var userRequest request.UserUpdateRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "updateUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid user id, must be a valid hex")
		c.JSON(errRest.Code, errRest)
	}

	domain := model.NewUserUpdateDomain(
		userRequest.Name,
		userRequest.Age,
	)

	err := uc.service.UpdateUserService(userId, domain)
	if err != nil {
		logger.Error(
			"Error trying to updateUser user",
			err,
			zap.String("journey", "updateUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"User updated successfully",
		zap.String("user_id", userId),
		zap.String("journey", "updateUser"))

	c.Status(http.StatusNoContent)
}
