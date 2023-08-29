package main

import (
	"go.mongodb.org/mongo-driver/mongo"
	"user.manager-crud-go/src/controller"
	"user.manager-crud-go/src/model/repository"
	"user.manager-crud-go/src/model/service"
)

func initDepedencies(database *mongo.Database,
) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)

}
