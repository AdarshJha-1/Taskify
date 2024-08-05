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

// Initialize the MongoDB collection for todos
var todoCollection *mongo.Collection = config.GetCollection(config.MongoClient, os.Getenv("TODO_COLLECTION"))

// GetTodosOfAUser retrieves all todos associated with a specific user
func GetTodosOfAUser(id primitive.ObjectID) ([]model.Todo, error) {

	// Find all todos where the user_id matches the provided id
	cursor, err := todoCollection.Find(context.Background(), bson.M{"user_id": id})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var todos []model.Todo
	// Decode all documents in the cursor into a slice of Todo model
	err = cursor.All(context.Background(), &todos)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

// CreateTodoForAUser inserts a new todo into the collection
func CreateTodoForAUser(todo model.Todo) (interface{}, error) {
	// Insert the new todo document into the collection
	result, err := todoCollection.InsertOne(context.Background(), todo)
	if err != nil {
		return nil, err
	}
	// Return the ID of the inserted document
	return result.InsertedID, nil
}

// DeleteTodoOfAUser deletes a specific todo associated with a user
func DeleteTodoOfAUser(userId, todoId primitive.ObjectID) (int64, error) {

	// Create a filter to match the todo ID and user ID
	filter := bson.M{
		"_id":     todoId,
		"user_id": userId,
	}

	// Delete the todo document matching the filter
	result, err := todoCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		return result.DeletedCount, err
	}
	// Return the count of deleted documents
	return result.DeletedCount, nil
}

// UpdateIsCompleted toggles the completion status of a specific todo
func UpdateIsCompleted(userId, todoId primitive.ObjectID) (int64, error) {

	// Create a filter to match the todo ID and user ID
	filter := bson.M{
		"_id":     todoId,
		"user_id": userId,
	}
	var todo model.Todo
	// Retrieve the todo document to check its current completion status
	err := todoCollection.FindOne(context.Background(), filter).Decode(&todo)
	if err != nil {
		return 0, err
	}
	// Define an update to toggle the is_completed field
	update := bson.M{
		"$set": bson.M{
			"is_completed": !todo.IsCompleted,
		},
	}
	// Apply the update to the todo document
	result, err := todoCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return result.ModifiedCount, err
	}
	// Return the count of modified documents
	return result.ModifiedCount, nil
}
