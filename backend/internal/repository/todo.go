package repository

import (
	"context"
	"os"

	"github.com/AdarshJha-1/Taskify/backend/config"
	"github.com/AdarshJha-1/Taskify/backend/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var todoCollection *mongo.Collection = config.GetCollection(config.MongoClient, os.Getenv("TODO_COLLECTION"))

func GetTodosOfAUser(id primitive.ObjectID) ([]model.Todo, error) {
	cursor, err := todoCollection.Find(context.Background(), bson.M{"user_id": id})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	var Todos []model.Todo
	err = cursor.All(context.Background(), &Todos)
	if err != nil {
		return nil, err
	}
	return Todos, nil
}

func CreateTodoForAUser(todo model.Todo) (interface{}, error) {

	result, err := todoCollection.InsertOne(context.Background(), todo)
	if err != nil {
		return nil, err
	}
	return result, nil
}
