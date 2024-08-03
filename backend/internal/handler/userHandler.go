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

		// Creating local user object
		var user model.User

		// Getting user data and decoding it into local user object
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			res := response.Response{Status: http.StatusBadRequest, Message: "Invalid input", Data: map[string]interface{}{"error": err.Error()}}
			json.NewEncoder(w).Encode(res)
			return
		}

		// Closing request body
		defer r.Body.Close()

		// Checking if user already exists or not
		userExists := repository.CheckExistingUser(user.Email, user.Username)
		if userExists {
			w.WriteHeader(http.StatusConflict)
			res := response.Response{Status: http.StatusConflict, Message: "User already exists", Data: map[string]interface{}{"error": "User already exists"}}
			json.NewEncoder(w).Encode(res)
			return
		}

		// Hashing password
		hashedPassword, err := utils.HashPassword(user.Password)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			res := response.Response{Status: http.StatusInternalServerError, Message: "Internal Server Error", Data: map[string]interface{}{"error": err.Error()}}
			json.NewEncoder(w).Encode(res)
			return
		}

		// Creating new user bson objectId and assigning it in user's id field
		user.Id = primitive.NewObjectID()

		// Updating password with hash password
		user.Password = hashedPassword

		// Creating new user with CreateUser function
		result, err := repository.CreateUser(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			res := response.Response{Status: http.StatusInternalServerError, Message: "Failed to create user", Data: map[string]interface{}{"error": err.Error()}}
			json.NewEncoder(w).Encode(res)
			return
		}

		// Sending success message with newly create user id
		w.WriteHeader(http.StatusCreated)
		res := response.Response{Status: http.StatusCreated, Message: "User created successfully", Data: map[string]interface{}{"data": result}}
		json.NewEncoder(w).Encode(res)
	}
}
