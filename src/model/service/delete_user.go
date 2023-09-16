package service

import (
	"go.uber.org/zap"
	"user.manager-crud-go/src/configuration/logger"
	"user.manager-crud-go/src/configuration/rest_err"
)

func (ud *userDomainService) DeleteUserService(userId string) *rest_err.RestErr {
	logger.Info("Init deleteUser domain",
		zap.String("journey", "deleteUser"),
	)

	err := ud.userRepository.DeleteUser(userId)
	if err != nil {
		logger.Error("Error trying to call repository",
			err,
			zap.String("journey", "deleteUser"))
		return err
	}

	logger.Info(
		"DeleteUser service executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "updateUser"))
	return nil
}
