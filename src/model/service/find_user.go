package service

import (
	"go.uber.org/zap"
	"user.manager-crud-go/src/configuration/logger"
	"user.manager-crud-go/src/configuration/rest_err"
	"user.manager-crud-go/src/model"
)

func (ud *userDomainService) FindUserService(string) (*model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUser domain", zap.String("journey", "findUser"))

	return nil, nil
}
