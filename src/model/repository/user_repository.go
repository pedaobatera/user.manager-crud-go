package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"user.manager-crud-go/src/configuration/rest_err"
	"user.manager-crud-go/src/model"
)

func NewUserRepository(
	database *mongo.Database,
) UserRepository {
	return &userRepository{
		database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.RestErr)
}
