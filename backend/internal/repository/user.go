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

// Initialize the MongoDB collection for users
var userCollection *mongo.Collection = config.GetCollection(config.MongoClient, os.Getenv("USER_COLLECTION"))

// CreateUser inserts a new user into the user collection
func CreateUser(user model.User) (interface{}, error) {

	result, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CheckExistingUser checks if a user with the given email or username already exists
func CheckExistingUser(email, username string) bool {

	// Create a filter to check for the presence of email or username
	filter := bson.M{"$or": []bson.M{
		{"email": email},
		{"username": username},
	}}

	// Exclude the password field from the result
	projection := bson.M{"password": 0}
	err := userCollection.FindOne(context.Background(), filter, options.FindOne().SetProjection(projection)).Err()

	if err == mongo.ErrNoDocuments {
		// No user found with the given email or username
		return false
	} else if err != nil {
		// An error occurred during the query
		return false
	}
	// User exists with the given email or username
	return true
}

// GetUser retrieves a user by email or username, including their password
func GetUser(identifier, password string) (*model.User, error) {

	// Create a filter to search by either email or username
	filter := bson.M{"$or": []bson.M{
		{"email": identifier},
		{"username": identifier},
	}}

	var user model.User

	// Find the user in the collection and decode the result into a User model
	err := userCollection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// No user found with the given identifier
			return nil, nil
		}
		// An error occurred during the query
		return nil, err
	}

	// Return the retrieved user
	return &user, nil
}
