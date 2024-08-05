package repository

import (
	"context"
	"fmt"
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
	var todos []model.Todo
	err = cursor.All(context.Background(), &todos)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func CreateTodoForAUser(todo model.Todo) (interface{}, error) {

	result, err := todoCollection.InsertOne(context.Background(), todo)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

func DeleteTodoOfAUser(userId, todoId primitive.ObjectID) (int64, error) {
	filter := bson.M{
		"_id":     todoId,
		"user_id": userId,
	}
	result, err := todoCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		return result.DeletedCount, err
	}
	return result.DeletedCount, nil
}

func UpdateIsCompleted(userId, todoId primitive.ObjectID) (int64, error) {
	filter := bson.M{
		"_id":     todoId,
		"user_id": userId,
	}
	var todo model.Todo
	err := todoCollection.FindOne(context.Background(), filter).Decode(&todo)
	if err != nil {
		return 0, err
	}
	update := bson.M{
		"$set": bson.M{
			"is_completed": !todo.IsCompleted,
		},
	}
	fmt.Println("here")
	result, err := todoCollection.UpdateOne(context.Background(), filter, update)
	fmt.Println("here", err)
	if err != nil {
		return result.ModifiedCount, err
	}
	return result.ModifiedCount, nil
}
