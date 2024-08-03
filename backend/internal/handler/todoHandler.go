package handler

import (
	"encoding/json"
	"net/http"

	"github.com/AdarshJha-1/Taskify/backend/internal/model"
	"github.com/AdarshJha-1/Taskify/backend/internal/repository"
	"github.com/AdarshJha-1/Taskify/backend/internal/response"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTodos() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		var todo model.Todo

		if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			res := response.Response{Status: http.StatusBadRequest, Message: "Invalid input", Data: map[string]interface{}{"error": err.Error()}}
			json.NewEncoder(w).Encode(res)
			return
		}
		defer r.Body.Close()
		result, err := repository.CreateTodoForAUser(todo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			res := response.Response{Status: http.StatusInternalServerError, Message: "Failed to create todo", Data: map[string]interface{}{"error": err.Error()}}
			json.NewEncoder(w).Encode(res)
			return
		}
		w.WriteHeader(http.StatusCreated)
		res := response.Response{Status: http.StatusCreated, Message: "Todo created successfully", Data: map[string]interface{}{"data": result}}
		json.NewEncoder(w).Encode(res)
	}
}

func GetTodo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		var todos []model.Todo

		defer r.Body.Close()
		params := mux.Vars(r)
		id := params["id"]
		userId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			res := response.Response{Status: http.StatusBadRequest, Message: "Invalid ID Formate", Data: map[string]interface{}{"error": err.Error()}}
			json.NewEncoder(w).Encode(res)
			return
		}

		todos, err = repository.GetTodosOfAUser(userId)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			res := response.Response{Status: http.StatusNotFound, Message: "Todos Not Found", Data: map[string]interface{}{"error": err.Error()}}
			json.NewEncoder(w).Encode(res)
			return
		}
		w.WriteHeader(http.StatusFound)
		res := response.Response{Status: http.StatusFound, Message: "Todos Founded", Data: map[string]interface{}{"todos": todos}}
		json.NewEncoder(w).Encode(res)
	}
}
