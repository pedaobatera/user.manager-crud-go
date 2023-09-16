package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"user.manager-crud-go/src/configuration/logger"
	"user.manager-crud-go/src/configuration/rest_err"
	"user.manager-crud-go/src/configuration/validation"
	"user.manager-crud-go/src/controller/model/request"
	"user.manager-crud-go/src/model"
	"user.manager-crud-go/src/view"
)

func (uc *userControllerInterface) LoginUser(c *gin.Context) {
	logger.Info("Init loginUser controller",
		zap.String("journey", "loginUser"),
	)
	var userRequest request.UserLogin

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info",
			err,
			zap.String("journey", "loginUser"),
		)
		errRest := rest_err.NewBadRequestError("Some fields are incorrect")
		errRest = validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserLoginDomain(
		userRequest.Email,
		userRequest.Password,
	)

	domainResult, token, err := uc.service.LoginUserService(domain)
	if err != nil {
		logger.Error(
			"Error trying to login user",
			err,
			zap.String("journey", "loginUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info(
		"User login successfully",
		zap.String("user_id", domainResult.GetID()),
		zap.String("journey", "loginUser"))

	c.Header("Authorization", token)
	c.JSON(http.StatusCreated, view.ConvertDomainToResponse(domainResult))
}
