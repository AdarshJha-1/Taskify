package repository

import (
	"context"
	"os"

	"github.com/AdarshJha-1/Taskify/backend/config"
	"github.com/AdarshJha-1/Taskify/backend/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = config.GetCollection(config.MongoClient, os.Getenv("USER_COLLECTION"))

func CreateUser(user model.User) (interface{}, error) {
	result, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return result, nil
}
