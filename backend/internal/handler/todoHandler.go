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

	id := r.Context().Value(config.UserIDKey).(string)
	// Converting it from string to bson objectId formate
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res := response.Response{Status: http.StatusBadRequest, Message: "Invalid ID Formate", Data: map[string]interface{}{"error": err.Error()}}
		json.NewEncoder(w).Encode(res)
		return
	}

	// Getting user input for todo
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res := response.Response{Status: http.StatusBadRequest, Message: "Invalid input", Data: map[string]interface{}{"error": err.Error()}}
		json.NewEncoder(w).Encode(res)
		return
	}

	// Closing request body
	defer r.Body.Close()

	// Updating user_id field in todo
	todo.Id = primitive.NewObjectID()
	todo.UserID = userId
	todo.IsCompleted = false

	// Creating todo by calling CreateTodoForAUser function
	result, err := repository.CreateTodoForAUser(todo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res := response.Response{Status: http.StatusInternalServerError, Message: "Failed to create todo", Data: map[string]interface{}{"error": err.Error()}}
		json.NewEncoder(w).Encode(res)
		return
	}

	// Sending success message with created todo id
	w.WriteHeader(http.StatusCreated)
	res := response.Response{Status: http.StatusCreated, Message: "Todo created successfully", Data: map[string]interface{}{"data": result}}
	json.NewEncoder(w).Encode(res)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var todos []model.Todo

	// Closing request body
	defer r.Body.Close()

	id := r.Context().Value(config.UserIDKey).(string)
	// Converting it from string to bson objectId formate
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res := response.Response{Status: http.StatusBadRequest, Message: "Invalid ID Formate", Data: map[string]interface{}{"error": err.Error()}}
		json.NewEncoder(w).Encode(res)
		return
	}

	// Fetching user todos from database by providing user_id
	todos, err = repository.GetTodosOfAUser(userId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		res := response.Response{Status: http.StatusNotFound, Message: "Todos Not Found", Data: map[string]interface{}{"error": err.Error()}}
		json.NewEncoder(w).Encode(res)
		return
	}
	if len(todos) == 0 {
		w.WriteHeader(http.StatusOK)
		res := response.Response{Status: http.StatusOK, Message: "No todos present, create one", Data: map[string]interface{}{"error": err}}
		json.NewEncoder(w).Encode(res)
		return
	}
	// Sending success message with all todos of that user
	w.WriteHeader(http.StatusFound)
	res := response.Response{Status: http.StatusFound, Message: "Todos Founded", Data: map[string]interface{}{"todos": todos}}
	json.NewEncoder(w).Encode(res)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// Closing request body
	defer r.Body.Close()

	id := r.Context().Value(config.UserIDKey).(string)
	// Converting it from string to bson objectId formate
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res := response.Response{Status: http.StatusBadRequest, Message: "Invalid ID Formate", Data: map[string]interface{}{"error": err.Error()}}
		json.NewEncoder(w).Encode(res)
		return
	}

	todoId, err := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res := response.Response{Status: http.StatusBadRequest, Message: "Invalid ID Formate", Data: map[string]interface{}{"error": err.Error()}}
		json.NewEncoder(w).Encode(res)
		return
	}

	// Fetching user todos from database by providing user_id
	result, err := repository.DeleteTodoOfAUser(userId, todoId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		res := response.Response{Status: http.StatusNotFound, Message: "Todos Not Found", Data: map[string]interface{}{"error": err.Error()}}
		json.NewEncoder(w).Encode(res)
		return
	}
	if result == 0 {
		w.WriteHeader(http.StatusOK)
		res := response.Response{Status: http.StatusOK, Message: "No todo found to delete with this id", Data: map[string]interface{}{"error": err}}
		json.NewEncoder(w).Encode(res)
		return
	}
	// Sending success message with all todos of that user
	w.WriteHeader(http.StatusFound)
	res := response.Response{Status: http.StatusFound, Message: "Todos Deleted Successfully", Data: map[string]interface{}{"todos deleted": result}}
	json.NewEncoder(w).Encode(res)
}

func ToggleIsCompletedTodo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// Closing request body
	defer r.Body.Close()

	id := r.Context().Value(config.UserIDKey).(string)
	// Converting it from string to bson objectId formate
	userId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res := response.Response{Status: http.StatusBadRequest, Message: "Invalid ID Formate", Data: map[string]interface{}{"error": err}}
		json.NewEncoder(w).Encode(res)
		return
	}

	todoId, err := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res := response.Response{Status: http.StatusBadRequest, Message: "Invalid ID Formate", Data: map[string]interface{}{"error": err}}
		json.NewEncoder(w).Encode(res)
		return
	}

	// Fetching user todos from database by providing user_id
	fmt.Println("before")
	result, err := repository.UpdateIsCompleted(userId, todoId)
	fmt.Println("after")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		res := response.Response{Status: http.StatusNotFound, Message: "Todos Not Found", Data: map[string]interface{}{"error": err}}
		json.NewEncoder(w).Encode(res)
		return
	}
	if result == 0 {
		w.WriteHeader(http.StatusOK)
		res := response.Response{Status: http.StatusOK, Message: "No todo found to toggle with this id", Data: map[string]interface{}{"error": err}}
		json.NewEncoder(w).Encode(res)
		return
	}
	// Sending success message with all todos of that user
	w.WriteHeader(http.StatusFound)
	res := response.Response{Status: http.StatusFound, Message: "Todos Updated Successfully", Data: map[string]interface{}{"todos updated": result}}
	json.NewEncoder(w).Encode(res)
}
