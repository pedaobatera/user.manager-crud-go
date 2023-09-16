package controller

import (
	"fmt"
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"user.manager-crud-go/src/configuration/logger"
	"user.manager-crud-go/src/configuration/rest_err"
	"user.manager-crud-go/src/model"
	"user.manager-crud-go/src/view"
)

func (uc *userControllerInterface) FindUserById(c *gin.Context) {
	logger.Info("Init findUserById controller",
		zap.String("journey", "findUserById"),
	)

	user, err := model.VerifyToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info(fmt.Sprintf("User authenticated successfully: %#v", user))

	userId := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error trying to parse user id",
			err,
			zap.String("journey", "findUserById"),
		)
		erroMessage := rest_err.NewBadRequestError(
			"User id is not valid",
		)

		c.JSON(erroMessage.Code, erroMessage)
		return
	}

	userDomain, err := uc.service.FindUserByIdServices(userId)
	if err != nil {
		logger.Error("Error trying to find user by id services",
			err,
			zap.String("journey", "findUserById"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserById controller executed successfully",
		zap.String("journey", "findUserById"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		userDomain,
	))

}

func (uc *userControllerInterface) FindUsersByEmail(c *gin.Context) {
	logger.Info("Init findUserByEmail controller",
		zap.String("journey", "findUserByEmail"),
	)
	userEmail := c.Param("userEmail")

	if _, err := mail.ParseAddress(userEmail); err != nil {
		logger.Error("Error trying to parse user email",
			err,
			zap.String("journey", "findUserByEmail"),
		)
		erroMessage := rest_err.NewBadRequestError(
			"User email is not valid",
		)

		c.JSON(erroMessage.Code, erroMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmailServices(userEmail)
	if err != nil {
		logger.Error("Error trying to find user by email services",
			err,
			zap.String("journey", "findUserByEmail"),
		)
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserByEmail controller executed successfully",
		zap.String("journey", "findUserById"),
	)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		userDomain,
	))
}
