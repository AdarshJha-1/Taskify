package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AdarshJha-1/Taskify/backend/config"
	"github.com/AdarshJha-1/Taskify/backend/internal/model"
	"github.com/AdarshJha-1/Taskify/backend/internal/repository"
	"github.com/AdarshJha-1/Taskify/backend/internal/response"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Creating new todo for the user with their user_id
func CreateTodos(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var todo model.Todo

	// Retrieve user ID from the request context
	id := r.Context().Value(config.UserIDKey).(string)

	// Convert user ID from string to BSON ObjectID format
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res := response.Response{Status: http.StatusBadRequest, Success: false, Message: "Invalid ID Formate", Data: map[string]interface{}{"error": err.Error()}}
		json.NewEncoder(w).Encode(res)
		return
	}

	// Decode the request body into the Todo model
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res := response.Response{Status: http.StatusBadRequest, Success: false, Message: "Invalid input", Data: map[string]interface{}{"error": err.Error()}}
		json.NewEncoder(w).Encode(res)
		return
	}

	// Close the request body
	defer r.Body.Close()

	// Set up the new Todo object
	todo.Id = primitive.NewObjectID()
	todo.UserID = userId
	todo.IsCompleted = false

	// Create the Todo in the database
	result, err := repository.CreateTodoForAUser(todo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res := response.Response{Status: http.StatusInternalServerError, Success: false, Message: "Failed to create todo", Data: map[string]interface{}{"error": err.Error()}}
		json.NewEncoder(w).Encode(res)
		return
	}

	// Respond with success and the created todo's ID
	w.WriteHeader(http.StatusCreated)
	res := response.Response{Status: http.StatusCreated, Success: true, Message: "Todo created successfully", Data: map[string]interface{}{"data": result}}
	json.NewEncoder(w).Encode(res)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var todos []model.Todo

	// Close the request body
	defer r.Body.Close()

	// Retrieve user ID from the request context
	id := r.Context().Value(config.UserIDKey).(string)

	// Convert user ID from string to BSON ObjectID format
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res := response.Response{Status: http.StatusBadRequest, Success: false, Message: "Invalid ID Formate", Data: map[string]interface{}{"error": err.Error()}}
		json.NewEncoder(w).Encode(res)
		return
	}

	// Fetch the todos for the user from the database
	todos, err = repository.GetTodosOfAUser(userId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		res := response.Response{Status: http.StatusNotFound, Success: false, Message: "Todos Not Found", Data: map[string]interface{}{"error": err.Error()}}
		json.NewEncoder(w).Encode(res)
		return
	}
	if len(todos) == 0 {
		w.WriteHeader(http.StatusOK)
		res := response.Response{Status: http.StatusOK, Success: true, Message: "No todos present, create one", Data: map[string]interface{}{"error": err}}
		json.NewEncoder(w).Encode(res)
		return
	}

	// Respond with success and the user's todos
	w.WriteHeader(http.StatusOK)
	res := response.Response{Status: http.StatusOK, Success: true, Message: "Todos Founded", Data: map[string]interface{}{"todos": todos}}
	json.NewEncoder(w).Encode(res)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// Close the request body
	defer r.Body.Close()

	// Retrieve user ID from the request context
	id := r.Context().Value(config.UserIDKey).(string)

	// Convert user ID from string to BSON ObjectID format
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res := response.Response{Status: http.StatusBadRequest, Success: false, Message: "Invalid ID Formate", Data: map[string]interface{}{"error": err.Error()}}
		json.NewEncoder(w).Encode(res)
		return
	}
	// Retrieve the todo ID from the URL parameters and convert it to BSON ObjectID format
	todoId, err := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res := response.Response{Status: http.StatusBadRequest, Success: false, Message: "Invalid ID Formate", Data: map[string]interface{}{"error": err.Error()}}
		json.NewEncoder(w).Encode(res)
		return
	}

	// Delete the todo from the database
	result, err := repository.DeleteTodoOfAUser(userId, todoId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		res := response.Response{Status: http.StatusNotFound, Success: false, Message: "Todos Not Found", Data: map[string]interface{}{"error": err.Error()}}
		json.NewEncoder(w).Encode(res)
		return
	}
	if result == 0 {
		w.WriteHeader(http.StatusOK)
		res := response.Response{Status: http.StatusOK, Success: true, Message: "No todo found to delete with this id", Data: map[string]interface{}{"error": err}}
		json.NewEncoder(w).Encode(res)
		return
	}

	// Respond with success and the number of deleted todos
	w.WriteHeader(http.StatusOK)
	res := response.Response{Status: http.StatusOK, Success: true, Message: "Todos Deleted Successfully", Data: map[string]interface{}{"todos deleted": result}}
	json.NewEncoder(w).Encode(res)
}

func ToggleIsCompletedTodo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// Close the request body
	defer r.Body.Close()

	// Retrieve user ID from the request context
	id := r.Context().Value(config.UserIDKey).(string)

	// Convert user ID from string to BSON ObjectID format
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res := response.Response{Status: http.StatusBadRequest, Success: false, Message: "Invalid ID Formate", Data: map[string]interface{}{"error": err}}
		json.NewEncoder(w).Encode(res)
		return
	}

	// Retrieve the todo ID from the URL parameters and convert it to BSON ObjectID format
	todoId, err := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res := response.Response{Status: http.StatusBadRequest, Success: false, Message: "Invalid ID Formate", Data: map[string]interface{}{"error": err}}
		json.NewEncoder(w).Encode(res)
		return
	}

	// Update the todo's completion status in the database
	fmt.Println("before")
	result, err := repository.UpdateIsCompleted(userId, todoId)
	fmt.Println("after")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		res := response.Response{Status: http.StatusNotFound, Success: false, Message: "Todos Not Found", Data: map[string]interface{}{"error": err}}
		json.NewEncoder(w).Encode(res)
		return
	}
	if result == 0 {
		w.WriteHeader(http.StatusOK)
		res := response.Response{Status: http.StatusOK, Success: true, Message: "No todo found to toggle with this id", Data: map[string]interface{}{"error": err}}
		json.NewEncoder(w).Encode(res)
		return
	}

	// Respond with success and the number of updated todos
	w.WriteHeader(http.StatusOK)
	res := response.Response{Status: http.StatusOK, Success: true, Message: "Todos Updated Successfully", Data: map[string]interface{}{"todos updated": result}}
	json.NewEncoder(w).Encode(res)
}
