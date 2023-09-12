package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"user.manager-crud-go/src/configuration/database/mongodb"
	"user.manager-crud-go/src/configuration/logger"
	"user.manager-crud-go/src/controller"
	"user.manager-crud-go/src/controller/routes"
	"user.manager-crud-go/src/model/repository"
	"user.manager-crud-go/src/model/service"
)

func main() {
	logger.Info("Starting the application...")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf(
			"Error when trying to connect to MongoDB: %s \n	",
			err.Error())
		return
	}

	// Init dependencies
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
