package repository

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"user.manager-crud-go/src/configuration/logger"
	"user.manager-crud-go/src/configuration/rest_err"
	"user.manager-crud-go/src/model"
	"user.manager-crud-go/src/model/repository/entity"
	"user.manager-crud-go/src/model/repository/entity/converter"
)

func (ur *userRepository) FindUserByEmail(
	email string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserByEmail repository",
		zap.String("journey", "findUserByEmail"))

	collection_name := os.Getenv(MONGO_USER_DB)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}

	err := collection.FindOne(context.Background(), filter).Decode(userEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with email: %s", email)
			logger.Error(errorMessage,
				err,
				zap.String("journey", "findUserByEmail"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}

		errorMessage := "Error trying to find user by email"
		logger.Error(errorMessage,
			err,
			zap.String("journey", "findUserByEmail"))
		return nil, rest_err.NewInternalServerError(errorMessage)

	}

	logger.Info("FindUserByEmail repository executed successfully",
		zap.String("journey", "findUserByEmail"),
		zap.String("userId", userEntity.ID.Hex()),
		zap.String("email", userEntity.Email))

	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserById(
	id string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init findUserById repository",
		zap.String("journey", "findUserById"))

	collection_name := os.Getenv(MONGO_USER_DB)
	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectId}}

	err := collection.FindOne(context.Background(), filter).Decode(userEntity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with id: %s", id)
			logger.Error(errorMessage,
				err,
				zap.String("journey", "findUserById"))
			return nil, rest_err.NewNotFoundError(errorMessage)
		}

		errorMessage := "Error trying to find user by id"
		logger.Error(errorMessage,
			err,
			zap.String("journey", "findUserById"))
		return nil, rest_err.NewInternalServerError(errorMessage)

	}

	logger.Info("FindUserById repository executed successfully",
		zap.String("journey", "findUserByEmail"),
		zap.String("userId", userEntity.ID.Hex()))

	return converter.ConvertEntityToDomain(*userEntity), nil
}
