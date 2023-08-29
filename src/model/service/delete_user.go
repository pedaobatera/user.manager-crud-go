package service

import (
	"go.uber.org/zap"
	"user.manager-crud-go/src/configuration/logger"
	"user.manager-crud-go/src/configuration/rest_err"
)

func (ud *userDomainService) DeleteUserService(string) *rest_err.RestErr {
	logger.Info("Init DeleteUser domain", zap.String("journey", "deleteUser"))

	return nil
}
