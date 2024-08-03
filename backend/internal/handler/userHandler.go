package handler

import (
	"encoding/json"
	"net/http"

	"github.com/AdarshJha-1/Taskify/backend/internal/model"
	"github.com/AdarshJha-1/Taskify/backend/internal/repository"
	"github.com/AdarshJha-1/Taskify/backend/internal/response"
	"github.com/AdarshJha-1/Taskify/backend/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var user model.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			res := response.Response{Status: http.StatusBadRequest, Message: "Invalid input", Data: map[string]interface{}{"error": err.Error()}}
			json.NewEncoder(w).Encode(res)
			return
		}
		defer r.Body.Close()
		hashedPassword, err := utils.HashPassword(user.Password)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			res := response.Response{Status: http.StatusInternalServerError, Message: "Internal Server Error", Data: map[string]interface{}{"error": err.Error()}}
			json.NewEncoder(w).Encode(res)
			return
		}

		user.Id = primitive.NewObjectID()
		user.Password = hashedPassword
		result, err := repository.CreateUser(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			res := response.Response{Status: http.StatusInternalServerError, Message: "Failed to create user", Data: map[string]interface{}{"error": err.Error()}}
			json.NewEncoder(w).Encode(res)
			return
		}
		w.WriteHeader(http.StatusCreated)
		res := response.Response{Status: http.StatusCreated, Message: "User created successfully", Data: map[string]interface{}{"data": result}}
		json.NewEncoder(w).Encode(res)
	}
}
