package service

import (
	"go.uber.org/zap"
	"user.manager-crud-go/src/configuration/logger"
	"user.manager-crud-go/src/configuration/rest_err"
	"user.manager-crud-go/src/model"
)

func (ud *userDomainService) UpdateUserService(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init UpdateUser domain",
		zap.String("journey", "updateUser"),
	)

	err := ud.userRepository.UpdateUser(userId, userDomain)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "updateUser"))
		return err
	}

	logger.Info(
		"UpdateUser service executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"))
	return nil

}
