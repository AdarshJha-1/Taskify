package repository

import (
	"context"
	"os"

	"github.com/AdarshJha-1/Taskify/backend/config"
	"github.com/AdarshJha-1/Taskify/backend/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userCollection *mongo.Collection = config.GetCollection(config.MongoClient, os.Getenv("USER_COLLECTION"))

func CreateUser(user model.User) (interface{}, error) {

	result, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func CheckExistingUser(email, username string) bool {

	filter := bson.M{"$or": []bson.M{
		{"email": email},
		{"username": username},
	}}

	projection := bson.M{"password": 0}
	err := userCollection.FindOne(context.Background(), filter, options.FindOne().SetProjection(projection)).Err()

	if err == mongo.ErrNoDocuments {
		return false
	} else if err != nil {
		return false
	}
	return true
}

func GetUser(identifier, password string) (*model.User, error) {
	filter := bson.M{"$or": []bson.M{
		{"email": identifier},
		{"username": identifier},
	}}

	var user model.User
	err := userCollection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
