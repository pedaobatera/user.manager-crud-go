package service

import (
	"user.manager-crud-go/src/configuration/rest_err"
	"user.manager-crud-go/src/model"
	"user.manager-crud-go/src/model/repository"
)

func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUserService(model.UserDomainInterface) (
		model.UserDomainInterface, *rest_err.RestErr)
}
